package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	version string
)

type options struct {
	user, password, url, value, output string
	version                            bool
	artefacts                          []string
}

func metadataDownload(url, ga string, option options) string {
	ga = strings.Replace(ga, ".", "/", -1)
	ga = strings.Replace(ga, ":", "/", -1)

	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+ga+"/maven-metadata.xml", nil)
	if option.user != "" && option.password != "" {
		req.Header.Add("Authorization", "Basic "+basicAuth(option.user, option.password))
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error downloading the meta data information for " + ga + " from " + url)
		os.Exit(1)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	metadataXML := buf.String()
	return metadataXML
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s [options] (groupId:artefactId)+\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	var option options
	flag.BoolVar(&(option.version), "version", false, "Version of the tool")
	flag.StringVar(&(option.user), "username", "", "the username for basic authentication")
	flag.StringVar(&(option.password), "password", "", "the password for basic authentication")
	flag.StringVar(&(option.url), "url", "", "the url to the remote repository !required!")
	flag.StringVar(&(option.value), "value", "release", "artefact version ( versions | release | latest )")
	flag.StringVar(&(option.output), "output", "version", "the output format ( version | full )")
	flag.Parse()
	option.artefacts = flag.Args()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}
	if option.version {
		fmt.Println("Version: ", version)
		os.Exit(0)
	}

	if option.url == "" {
		fmt.Println("Missing parameter url for the remote repository!")
		flag.Usage()
		os.Exit(1)
	}

	if len(option.artefacts) <= 0 {
		fmt.Println("Missing the list of artefacts!")
		flag.Usage()
		os.Exit(1)
	}

	for _, artefact := range option.artefacts {
		metadataXML := metadataDownload(option.url, artefact, option)
		switch option.value {
		case "latest":
			latestRegex, _ := regexp.Compile("<latest>(.+?)</latest>")
			latest := latestRegex.FindStringSubmatch(metadataXML)[1]
			if option.output == "full" {
				latest = artefact + ":" + latest
			}
			fmt.Println(latest)
		case "release":
			releaseRegex, _ := regexp.Compile("<release>(.+?)</release>")
			release := releaseRegex.FindStringSubmatch(metadataXML)[1]
			if option.output == "full" {
				release = artefact + ":" + release
			}
			fmt.Println(release)
		case "versions":
			versionRegex, _ := regexp.Compile("<version>(.+?)</version>")
			versions := versionRegex.FindAllStringSubmatch(metadataXML, -1)

			for _, item := range versions {
				version := item[1]
				if option.output == "full" {
					version = artefact + ":" + version
				}
				fmt.Println(version)
			}
		}
	}
}

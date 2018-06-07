package internal

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// Options are attributes to download meta data
type Options struct {
	Username, Password, URL, Value, Output string
}

// ShowVersion the main functionality of the mvnrc client
func ShowVersion(option Options, artefact string) {
	metadataXML := downloadMetadata(option.URL, artefact, option.Username, option.Password)
	switch option.Value {
	case "latest":
		latestRegex, _ := regexp.Compile("<latest>(.+?)</latest>")
		latest := latestRegex.FindStringSubmatch(metadataXML)[1]
		if option.Output == "full" {
			latest = artefact + ":" + latest
		}
		fmt.Println(latest)
	case "release":
		releaseRegex, _ := regexp.Compile("<release>(.+?)</release>")
		release := releaseRegex.FindStringSubmatch(metadataXML)[1]
		if option.Output == "full" {
			release = artefact + ":" + release
		}
		fmt.Println(release)
	case "versions":
		versionRegex, _ := regexp.Compile("<version>(.+?)</version>")
		versions := versionRegex.FindAllStringSubmatch(metadataXML, -1)

		for _, item := range versions {
			version := item[1]
			if option.Output == "full" {
				version = artefact + ":" + version
			}
			fmt.Println(version)
		}
	}
}

func downloadMetadata(url, artefact, username, password string) string {
	artefact = strings.Replace(artefact, ".", "/", -1)
	artefact = strings.Replace(artefact, ":", "/", -1)

	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+artefact+"/maven-metadata.xml", nil)
	if username != "" && password != "" {
		req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error downloading the meta data information for " + artefact + " from " + url)
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

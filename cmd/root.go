package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/lorislab/mvnrc/internal"
)

var (
	// Version of the project
	version string
	// Build is a hash code from the repository
	build string
)

var rootCmd = &cobra.Command{
	Use:     "mvnrc",
	Short:   "mvnrc is remote maven client to read version",
	Long:    "A very simple utility to read artefact version from the remote repository",
	Version: version + " [" + build + "]",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, artefact := range args {
			internal.ShowVersion(option, artefact)
		}
	},
}

var option = internal.Options{}

func init() {
	//rootCmd.SetUsageTemplate("mvnrc [flags] <groupdId:artefactId>")
	rootCmd.Flags().StringVarP(&(option.Username), "username", "u", "", "The user name for the basic authentication")
	rootCmd.Flags().StringVarP(&(option.Password), "password", "p", "", "The password for the basic authentication")
	rootCmd.Flags().StringVarP(&(option.URL), "url", "r", "", "The url to the remote repository (required)")
	rootCmd.Flags().StringVarP(&(option.Value), "value", "e", "release", "The artefact version ( versions | release | latest )")
	rootCmd.Flags().StringVarP(&(option.Output), "output", "o", "version", "The output format ( version | full )")
	rootCmd.MarkFlagRequired("url")
}

// Execute method execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"CloudDevKubernetes/util"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: versionTitle,
	Long:  `CloudDevKubernetes.com`,
	Run: func(_ *cobra.Command, _ []string) {
		util.VersionFucn()
	},
}
var Version string

func init() {
	rootCmd.AddCommand(versionCmd)
}

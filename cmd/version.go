package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	Version = "0.0.1"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of SiteMap",
	Long:  `All software has versions. This is SiteMap's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SiteMap Static Site Generator " + Version + " -- HEAD")
	},
}

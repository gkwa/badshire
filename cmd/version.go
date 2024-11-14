package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"

	"github.com/gkwa/badshire/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of badshire",
	Long:  `All software has versions. This is badshire's`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := version.GetBuildInfo()
		fmt.Println(buildInfo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

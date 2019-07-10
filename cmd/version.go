package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of DGraphCms",
	Long:  `All software has versions. This is DGraphCms's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DGraphCms v0.1 -- HEAD")
	},
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initDBCmd)
}

var initDBCmd = &cobra.Command{
	Use:   "initdb",
	Short: "Add the database schema into the dgrapgh database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented jet")
	},
}

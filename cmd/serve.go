package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the GraphQL api server",
	Long:  `ENV variables 
	PORT default 80
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented jet")
	},
}

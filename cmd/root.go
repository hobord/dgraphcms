package cmd
import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
  Use:   "dgraphcms",
  Short: "dgraphcms is a graph database based cms with GraphQL api",
  Long: ``,
  Run: func(cmd *cobra.Command, args []string) {
    // Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

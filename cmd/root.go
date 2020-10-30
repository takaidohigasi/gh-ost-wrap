package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/takaidohigasi/gh-ost-wrap/config"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Short: "gh-ost wrapper",
	Long:  "gh-ost command wrapper",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
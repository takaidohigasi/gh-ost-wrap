package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var ghostDefaultConfig = "gh-ost.yaml"


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "execute",
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
	//cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP("config", "c", ghostDefaultConfig, "config path for default gh-ost parameteter")
}
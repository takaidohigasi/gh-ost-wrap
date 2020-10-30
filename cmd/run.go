package cmd

import (
	//	"context"
	//	"errors"
	"fmt"
	"github.com/takaidohigasi/gh-ost-wrap/config"

	"github.com/spf13/cobra"
    "github.com/iancoleman/strcase"

	//	"io/ioutil"
	//	"os"
	//	"os/signal"
	//	"sync"
	//	"syscall"
	//	"time"

	//	log "github.com/sirupsen/logrus"

)

var (
	ghostConfig     GhostConf
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "gh-ost wrapper",
	Long:  "gh-ost command wrapper",
	RunE:  ghostRun,
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&ghostConfigFile, "config", "c", ghostDefaultConfig, "config path for default gh-ost parameteter")
}

func ghostRun(cmd *cobra.Command, args []string) error {

	ghostConfig, err := config.LoadFromFile(ghostConfigFile)
	if err != nil {
		return err
	}
	fmt.Println(strcase.ToKebab("MaxLoad"))
	fmt.Print("%v", ghostConfig)
	return nil
}

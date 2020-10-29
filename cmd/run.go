package cmd

import (
//	"context"
//	"errors"
	"fmt"
//	"io/ioutil"
//	"os"
//	"os/signal"
//	"sync"
//	"syscall"
//	"time"

//	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)



var runCmd = &cobra.Command{
	Use:   "run",
	Short: "gh-ost wrapper",
	Long:  "gh-ost command wrapper",
	RunE:  ghostRun,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func ghostRun (cmd *cobra.Command, args []string) error {
	fmt.Print("%v", ghostConfig)
	return nil
}
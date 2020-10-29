package cmd

import (
//	"context"
//	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
//	"os"
//	"os/signal"
//	"sync"
//	"syscall"
//	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type GhostConf struct {
	MaxLoad         string     `yaml:"max-load"`
	CriticalLoad    string     `yaml:"critical-load"`
}

var runCmd = &cobra.Command{
	Use:   "run",
	RunE:  ghostRun,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func ghostRun (cmd *cobra.Command, args []string) error {
	buf, err := ioutil.ReadFile("./gh-ost-defaults.yaml.sample")
	if err != nil {
		fmt.Println(err)
		return
	}

	var conf GhostConf
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("d: %+v", conf)
}
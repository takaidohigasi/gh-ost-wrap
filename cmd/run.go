package cmd

import (
	//	"context"
	//	"errors"
	"fmt"
	config "github.com/takaidohigasi/gh-ost-wrap/lib"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
    "github.com/iancoleman/strcase"
	yaml "gopkg.in/yaml.v2"

	//	"io/ioutil"
	//	"os"
	//	"os/signal"
	//	"sync"
	//	"syscall"
	//	"time"

	//	log "github.com/sirupsen/logrus"

)

type GhostConf struct {
	ChunkSize                    int      `yaml:"chunk-size"`
	CriticalLoad                 string   `yaml:"critical-load"`
	CriticalLoadHibernateSeconds int      `yaml:"critical-load-hibernate-seconds"`
	CutOverLockTimeoutSeconds    int      `yaml:"cut-over-lock-timeout-seconds"`
	DefaultRetries               int      `yaml:"default-retries"`
	DefaultOptions               []string `yaml:"default-options"`
	MaxLagMillis                 int      `yaml:"max-lag-millis"`
	MaxLoad                      string   `yaml:"max-load"`
	NiceRatio                    float32  `yaml:"nice-ratio"`
}

var (
	ghostConfigFile string
	ghostConfig     GhostConf
)

var ghostDefaultConfig = "gh-ost.yaml"

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

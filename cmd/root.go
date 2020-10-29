package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "run",
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
	rootCmd.PersistentFlags().StringVarP(&ghostConfigFile, "config", "c", ghostDefaultConfig, "config path for default gh-ost parameteter")

	// check file existence
	_, err := os.Stat(ghostConfigFile)
	if err != nil {
		er(err)
	}

	// read yaml
	buf, err := ioutil.ReadFile(ghostConfigFile)
	if err != nil {
		er(err)
	}

	err = yaml.Unmarshal(buf, &ghostConfig)
	if err != nil {
		er(err)
	}
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

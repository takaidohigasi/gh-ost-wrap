package config

import (
	"io/ioutil"
	"os"
)

type GhostConf struct {
	ChunkSize                    int      `yaml:"chunk-size"`
	Conf                         string   `yaml:"conf"`
	CriticalLoad                 string   `yaml:"critical-load"`
	CriticalLoadHibernateSeconds int      `yaml:"critical-load-hibernate-seconds"`
	CutOverLockTimeoutSeconds    int      `yaml:"cut-over-lock-timeout-seconds"`
	DefaultRetries               int      `yaml:"default-retries"`
	DefaultOptions               []string `yaml:"default-options"`
	MaxLagMillis                 int      `yaml:"max-lag-millis"`
	MaxLoad                      string   `yaml:"max-load"`
	NiceRatio                    float32  `yaml:"nice-ratio"`
}

type ThrottleControllReplicas struct {
	Replicas []string
}

func LoadFromFile(path string) (GhostConf, error) {
	// check file existence
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// read yaml
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var ghostConfig GhostConf
	err = yaml.Unmarshal(buf, &ghostConfig)
	if err != nil {
		return nil, err
	}
	return ghostConfig, err
}

// func ToString
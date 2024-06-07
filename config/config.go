package config

import (
	_ "embed"
	"github.com/spf13/viper"
	"os"
)

//go:embed config.json
var config []byte

var (
	Conf *Config
)

func init() {
	InitConfig()
}

func InitConfig() {
	filepath := "config.json"
	_, err := os.Stat(filepath)
	// If no configuration in the directory.
	// Generate the configuration.
	// todo INCOMPLETE: Generate without config.json
	if err != nil {
		if os.IsNotExist(err) {
			err := os.WriteFile(filepath, config, 0644)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}

func LrString(left, right string) string {
	if len(left) == 0 {
		return right
	}
	return left
}

func LrInt(left, right int) int {
	if left == 0 {
		return right
	}
	return left
}

package config

import "C"
import (
	"Go-Grasscutter/src/utils"
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
	"log"
	"os"
)

//go:embed config.json
var config []byte

var (
	Conf *Config
	L    *viper.Viper
)

func InitConfig() {
	filepath := "config.json"
	_, err := os.Stat(filepath)
	// If no configuration in the directory.
	// Generate the configuration.
	// todo Generate without config.json
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
	initLanguage()
}

func initLanguage() {
	L = viper.New()
	L.AddConfigPath("./")
	prefix := "resources/languages/"
	// todo fix
	r, err := utils.GetResource().ReadFile(prefix + Conf.Language.Language + ".json")
	if err != nil {
		log.Println(err)
		r, err = utils.GetResource().ReadFile(prefix + "zh-CN.json")
		if err != nil {
			// en-US.json should exist
			panic(err)
		}
	}
	if err = L.ReadConfig(bytes.NewBuffer(r)); err != nil {
		// en-US.json should exist
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

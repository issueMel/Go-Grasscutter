package config

import (
	_ "embed"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

//go:embed config.json
var config []byte

func InitConfig() {
	fp := filepath.Join("config.json")
	_, err := os.Stat(fp)
	// If no configuration in the directory.
	// Generate the configuration.
	if err != nil {
		if os.IsNotExist(err) {
			err := os.WriteFile("config.json", config, 0644)
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
		log.Println(err)
		return
	}
}

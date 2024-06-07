package lang

import (
	"Go-Grasscutter/log"
	"Go-Grasscutter/utils"
	"bytes"
	"github.com/spf13/viper"
)

// res/languages
var L *viper.Viper

func LoadLanguage() {
	L = viper.New()
	L.AddConfigPath("./")
	prefix := "res/languages/"
	// todo INCOMPLETE: use config code
	//r, err := utils.GetResource().ReadFile(prefix + config.Conf.Language.Language + ".json")
	code, err := getLanguageCode()
	r, err := utils.GetResource().ReadFile(prefix + code + ".json")
	if err != nil {
		log.SugaredLogger.Info(err)
		r, err = utils.GetResource().ReadFile(prefix + "en-US.json")
		if err != nil {
			// en-US.json should exist
			panic(err)
		}
	}
	if err = L.ReadConfig(bytes.NewBuffer(r)); err != nil {
		// language config should exist
		panic(err)
	}
}

func Translate(key string) string {
	if L.Get(key).(string) == "" {
		log.SugaredLogger.Error("Failed to format string: ", key)
	}
	return L.Get(key).(string)
}

func getLanguageCode() (string, error) {
	var language string
	var err error
	language, err = getLanguage()
	if err != nil {
		return "", err
	}
	return language, nil
}

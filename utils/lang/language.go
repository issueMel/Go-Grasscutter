package lang

import (
	"Go-Grasscutter/log"
	"Go-Grasscutter/utils"
	"bytes"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"os"
	"runtime"
	"syscall"
	"unsafe"
)

// resources/languages
var L *viper.Viper

func LoadLanguage() {
	L = viper.New()
	L.AddConfigPath("./")
	prefix := "resources/languages/"
	// todo use config code
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

	switch runtime.GOOS {
	case "windows":
		language, err = getWindowsLanguage()
	default: // for Linux, macOS, and other Unix-like systems
		language = getLinuxLanguage()
		if language == "" {
			err = errors.New("Can`t fetch the system language")
		}
	}
	if err != nil {
		return "", err
	}
	return language, nil
}

func getWindowsLanguage() (string, error) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procGetUserDefaultLocaleName := kernel32.NewProc("GetUserDefaultLocaleName")

	var buf [85]uint16
	ret, _, err := procGetUserDefaultLocaleName.Call(uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 {
		return "", err
	}
	return syscall.UTF16ToString(buf[:]), nil
}

func getLinuxLanguage() string {
	language := os.Getenv("LC_ALL")
	if language == "" {
		language = os.Getenv("LANG")
	}
	return language
}

//go:build !windows

package lang

import (
	"errors"
	"os"
)

func getLanguage() (string, error) {
	language := os.Getenv("LC_ALL")
	if language == "" {
		language = os.Getenv("LANG")
	}
	if language == "" {
		return "", errors.New("Can`t fetch the system language")
	}
	return language, nil
}

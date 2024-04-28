package utils

import (
	"os"
)

func readResource(resourcePath string) []byte {
	data, err := os.ReadFile(resourcePath)
	if err != nil {
		panic(err)
	}
	return data
}

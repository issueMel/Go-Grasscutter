package utils

import (
	"os"
)

func ReadResource(resourcePath string) []byte {
	data, err := os.ReadFile(resourcePath)
	if err != nil {
		panic(err)
	}
	return data
}

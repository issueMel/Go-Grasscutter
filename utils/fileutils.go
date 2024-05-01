package utils

import (
	"embed"
)

var r embed.FS

func InitResource(res embed.FS) {
	r = res
}

func ReadResource(resourcePath string) []byte {
	data, err := r.ReadFile("resources" + resourcePath)
	if err != nil {
		panic(err)
	}
	return data
}

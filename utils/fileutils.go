package utils

import (
	"embed"
)

var r *embed.FS

func InitResource(res embed.FS) {
	r = &res
}
func GetResource() *embed.FS {
	return r
}
func ReadResource(resourcePath string) []byte {
	// filepath.Join() give wrong path in Windows with embed
	data, err := r.ReadFile("res" + resourcePath)
	if err != nil {
		panic(err)
	}
	return data
}

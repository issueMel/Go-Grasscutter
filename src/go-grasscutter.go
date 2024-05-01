package main

import (
	"Go-Grasscutter/src/db"
	"Go-Grasscutter/src/server/http"
	"Go-Grasscutter/utils"
	"embed"
	_ "embed"
)

//go:embed resources/*
var resource embed.FS

func main() {
	utils.InitResource(resource)
	utils.LoadKeys() // Load keys from buffers.

	// Parse start-up arguments.

	// Get the server run mode.

	// Create command map.

	// Initialize db.
	db.InitDatabase()
	// Initialize the default systems.
	//_ = new(auth.PasswordAuthenticator)
	//_ = new(command.DefaultPermissionHandler)
	// Create server instances.
	r := http.InitRouter()
	http.ProxyNoRouteRequest("http://127.0.0.1:443")
	r.Spin()
}

package main

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/db"
	"Go-Grasscutter/server/http/router"

	"Go-Grasscutter/utils"
	"Go-Grasscutter/utils/crypto"
	"Go-Grasscutter/utils/lang"
	"embed"
)

//go:embed resources/*
var resource embed.FS

func main() {
	// Load server resource.
	utils.InitResource(resource)
	// Load server configuration.
	config.InitConfig()
	// Load translation files.
	lang.LoadLanguage()

	crypto.LoadKeys() // Load keys from buffers.

	// Parse start-up arguments.

	// Get the server run mode.

	// Create command map.

	// Initialize db.
	db.InitDatabase()

	// Create server instances.
	r := router.InitRouter()

	// Start servers.
	//go kcp.Init()
	r.Spin()
}

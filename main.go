package main

import (
	"embed"
	"etcdm/internal"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "etcdm",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: append([]interface{}{}, internal.BindApi()...),
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

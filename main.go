package main

import (
	_ "embed"
	"github.com/wailsapp/wails"
	"link-server/context"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "天天算力客户端",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(&context.AccountApp)
	app.Bind(&context.LoginApp)
	app.Bind(&context.P2pApp)
	app.Bind(&context.ResourceApp)
	app.Bind(&context.SettingApp)
	err := app.Run()
	if err != nil {
		panic("Wails app run error")
	}
}

package main

import (
	"embed"
	"fmt"

	"bytemystery.com/picbutton"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

//go:embed assets/*
var content embed.FS

func main() {
	a := app.New()
	w := a.NewWindow("-Test-")

	imgPlayUp, _ := content.ReadFile("assets/play_u.png")
	imgPlayDown, _ := content.ReadFile("assets/play_d.png")
	imgPlayUpX, _ := content.ReadFile("assets/play_ux.png")
	imgPlayDownX, _ := content.ReadFile("assets/play_dv.png")

	imgStopUp, _ := content.ReadFile("assets/stop_u.png")
	imgStopDown, _ := content.ReadFile("assets/stop_d.png")

	var button1 *picbutton.PicButton
	var button2 *picbutton.PicButton
	button1 = picbutton.NewPicButton(imgPlayUp, imgPlayDown, imgPlayUpX, imgPlayDownX, true, 0, func() {
		fmt.Println("Add click", button1.GetLastkeyModifier())
	})
	button2 = picbutton.NewPicButton(imgStopUp, imgStopDown, nil, nil, false, 0, func() {
		fmt.Println("Add click", button2.GetLastkeyModifier())
	})

	x := container.NewHBox(button1, button2)
	w.SetContent(x)
	w.ShowAndRun()
}

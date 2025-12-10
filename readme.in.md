# PicButton for [fyne](https://fyne.io/)
// Implements a button (normal push button or toggle button) with
// customer provided pictures for up, down and inactive state
// if inactive pictures are missing they will be generated from the provided up and down pictures.
// Pictures can be changed on the fly.
// You can also specify which mouse button can be used to press / toggle the button.
// Also the keyboard keyState and used Mouse button can be retrieved for implementing click + Ctrl
// or right click + Shift.

Author: Reiner Pröls  
Licence: MIT  

## Getting PicButton
`go get https://github.com/Reiner1210/picbutton`

## Import PicButton
`bytemystery.com/picbutton`

## Usage of PicButton
`button := picbutton.NewPicButton(imgUp, imgDown, imgUpX, imgDownX, false, 0, func() {// Do what has to be done})`

Example:
```
package main

import (
	"embed"
	"fmt"

	"bytemystery.com/picbutton"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
)

//go:embed assets/*
var content embed.FS

func main() {
	a := app.New()
	w := a.NewWindow("PicButton")
	w.SetFixedSize(true)

	imgPlayUp, _ := content.ReadFile("assets/play_u.png")
	imgPlayDown, _ := content.ReadFile("assets/play_d.png")
	imgPlay2Down, _ := content.ReadFile("assets/play2_d.png")
	imgPlayUpX, _ := content.ReadFile("assets/play_ux.png")
	imgPlayDownX, _ := content.ReadFile("assets/play_dx.png")

	imgStopUp, _ := content.ReadFile("assets/stop_u.png")
	imgStopDown, _ := content.ReadFile("assets/stop_d.png")

	var play *picbutton.PicButton
	var stop *picbutton.PicButton
	play = picbutton.NewPicButton(imgPlayUp, imgPlayDown, imgPlayUpX, imgPlayDownX, true, desktop.MouseButtonPrimary|desktop.MouseButtonSecondary|desktop.MouseButtonTertiary,
		func() {
			fmt.Println("Play click", play.GetLastKeyModifier(), play.GetLastMouseButton())
			stop.SetEnabled(play.IsDown())
		})
	stop = picbutton.NewPicButton(imgStopUp, imgStopDown, nil, nil, false, desktop.MouseButtonPrimary|desktop.MouseButtonSecondary|desktop.MouseButtonTertiary, func() {
		fmt.Println("Stop click", stop.GetLastKeyModifier(), stop.GetLastMouseButton())

		if stop.GetLastKeyModifier() == fyne.KeyModifierControl {
			play.SetDImg(imgPlay2Down)
		} else if stop.GetLastKeyModifier() == fyne.KeyModifierShift {
			play.SetDImg(imgPlayDown)
		} else {
			play.SetDown(false)
			stop.SetEnabled(false)
		}
	})
	stop.SetEnabled(false)

	x := container.NewHBox(play, stop)
	w.SetContent(x)
	w.ShowAndRun()
}
```

![Screenshot of the test program.](examples/screenshots/1.png)

<!-- BEGIN AUTO -->

<!-- END AUTO -->
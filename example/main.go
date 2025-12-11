package main

import (
	"embed"
	"fmt"
	"image/color"

	"github.com/bytemystery-com/picbutton"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

//go:embed assets/*
var content embed.FS

func main() {
	a := app.New()
	w := a.NewWindow("PicButton")
	w.SetFixedSize(true)
	w.SetPadded(false)

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

	bg := canvas.NewRectangle(color.NRGBA{R: 192, G: 192, B: 192, A: 255})
	sep := widget.NewSeparator()
	hbox := container.NewHBox(sep, play, stop, sep)
	vbox := container.NewVBox(sep, hbox, sep)
	w.SetContent(container.NewStack(bg, vbox))

	w.ShowAndRun()
}

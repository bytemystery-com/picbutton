// Copyright (c) 2025-2016 Reiner Pröls
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// SPDX-License-Identifier: MIT

package main

import (
	"embed"
	"fmt"
	"image/color"
	"os"

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

	w.CenterOnScreen()

	imgPlayUp, _ := content.ReadFile("assets/play_u.png")
	imgPlayDown, _ := content.ReadFile("assets/play_d.png")
	imgPlay2Down, _ := content.ReadFile("assets/play2_d.png")
	imgPlayUpX, _ := content.ReadFile("assets/play_ux.png")
	imgPlayDownX, _ := content.ReadFile("assets/play_dx.png")

	imgStopUp, _ := content.ReadFile("assets/stop_u.png")
	imgStopDown, _ := content.ReadFile("assets/stop_d.png")

	imgExitUp, _ := content.ReadFile("assets/exit_u.png")
	imgExitDown, _ := content.ReadFile("assets/exit_d.png")

	imgHoover, _ := content.ReadFile("assets/hoover.png")
	imgHooverBottom, _ := content.ReadFile("assets/hoover_b.png")
	imgHooverTop, _ := content.ReadFile("assets/hoover_t.png")

	var play *picbutton.PicButton
	var stop *picbutton.PicButton
	var exit1 *picbutton.PicButton
	var exit2 *picbutton.PicButton
	play = picbutton.NewPicButton(imgPlayUp, imgPlayDown, imgPlayUpX, imgPlayDownX, true,
		func() {
			fmt.Println("Play click primary", play.GetLastKeyModifier(), play.GetLastMouseButton())
			stop.SetEnabled(play.IsDown())
		},
		func() {
			fmt.Println("Play click secondary", play.GetLastKeyModifier(), play.GetLastMouseButton())
			stop.SetEnabled(play.IsDown())
		})
	// desktop.MouseButtonTertiary
	stop = picbutton.NewPicButtonEx(imgStopUp, imgStopDown, nil, nil, false, true,
		desktop.MouseButtonPrimary|desktop.MouseButtonSecondary|desktop.MouseButtonTertiary,
		func() {
			str := "primary"
			if stop.GetLastMouseButton() == desktop.MouseButtonTertiary {
				str = "tertiary"
			}
			fmt.Println("Stop click", str, stop.GetLastKeyModifier())

			if stop.GetLastKeyModifier() == fyne.KeyModifierControl {
				play.SetDImg(imgPlay2Down)
			} else if stop.GetLastKeyModifier() == fyne.KeyModifierShift {
				play.SetDImg(imgPlayDown)
			} else {
				play.SetDown(false)
				stop.SetEnabled(false)
			}
		},
		func() {
			fmt.Println("Stop click secondary", stop.GetLastKeyModifier(), stop.GetLastMouseButton())

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

	// without padding
	exit1 = picbutton.NewPicButtonEx(imgExitUp, imgExitDown, nil, nil, false, false, 0,
		func() {
			exit2.SetEnabled(true)
			exit1.SetEnabled(false)
		}, nil)

	// without padding
	exit2 = picbutton.NewPicButtonEx(imgExitUp, imgExitDown, nil, nil, false, false, 0,
		func() {
			os.Exit(0)
		}, nil)
	exit2.SetEnabled(false)

	hoover := picbutton.NewPicButton(imgHoover, imgHoover, nil, nil, false, func() {}, nil)
	hoover.SetHooverImg(imgHooverBottom, imgHooverTop)

	bg := canvas.NewRectangle(color.NRGBA{R: 192, G: 192, B: 192, A: 255})
	sep := widget.NewSeparator()
	hbox := container.NewHBox(sep, play, stop, exit1, exit2, hoover, sep)
	vbox := container.NewVBox(sep, hbox, sep)
	w.SetContent(container.NewStack(bg, vbox))

	w.ShowAndRun()
}

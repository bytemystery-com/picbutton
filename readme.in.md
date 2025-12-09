# PicButton for [fyne](https://fyne.io/)
Implements a button (normal push button or toggle button) with  
customer provided pictures for up, down and inactive state  
if inactive pictures are missing they will be generated from the provided up and down pictures.  
you can also specify which mouse button can be used to press / toggle the button  
also the keyboard keyState can be retrieved for implementing click + Ctrl  


Example:
```
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

	imgUp, _ := content.ReadFile("assets/img_u.png")
	imgDown, _ := content.ReadFile("assets/img_d.png")
	imgUpX, _ := content.ReadFile("assets/img_ux.png")
	imgDownX, _ := content.ReadFile("assets/img_dv.png")

	var button1 *picbutton.PicButton
	var button2 *picbutton.PicButton
	button1 = picbutton.NewPicButton(imgUp, imgDown, imgUpX, imgDownX, false, 0, func() {
		fmt.Println("Add click", button1.GetLastkeyModifier())
	})
	button2 = picbutton.NewPicButton(imgUp, imgDown, nil, nil, true, 0, func() {
		fmt.Println("Add click", button2.GetLastkeyModifier())
	})

	x := container.NewHBox(button1, button2)
	w.SetContent(x)
	w.ShowAndRun()
}
```

<!-- BEGIN AUTO -->

<!-- END AUTO -->
# PicButton for [fyne](https://fyne.io/)
Implements a button (normal push button or toggle button) with  
customer provided pictures for up, down and inactive state  
if inactive pictures are missing they will be generated from the provided up and down pictures.  
Pictures can be changed on the fly.  
You can also specify which mouse button can be used to press / toggle the button.  
Also the keyboard keyState and used Mouse button can be retrieved for implementing click + Ctrl  
or right click + Shift.  
You can also specify if the padding from the theme is used or displaying without a padding.

Author: Reiner Pröls  
Licence: MIT  

## Getting PicButton
```go
go get github.com/bytemystery-com/picbutton
```

## Import PicButton
```go
import github.com/bytemystery-com/picbutton`
```

## Usage of PicButton
```go
button := picbutton.NewPicButton(imgUp, imgDown, imgUpX, imgDownX, false, 
		func() {  
			// Do what has to be done by primary mouseclick
		},
		func() {  
			// Do what has to be done by secondary mouseclick
		})
```


Example:
```go
// Copyright (c) 2025 Reiner Pröls
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

	bg := canvas.NewRectangle(color.NRGBA{R: 192, G: 192, B: 192, A: 255})
	sep := widget.NewSeparator()
	hbox := container.NewHBox(sep, play, stop, exit1, exit2, sep)
	vbox := container.NewVBox(sep, hbox, sep)
	w.SetContent(container.NewStack(bg, vbox))

	w.ShowAndRun()
}
```

## Screenshots from the example
![alt text](/example/screenshots/01.png "Screenshot 01")
![alt text](/example/screenshots/02.png "Screenshot 02")
![alt text](/example/screenshots/03.png "Screenshot 03")

## Docu 

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# picbutton

```go
import "github.com/bytemystery-com/picbutton"
```

## Index

- [type PicButton](<#PicButton>)
  - [func NewPicButton\(uImg \[\]byte, dImg \[\]byte, uxImg \[\]byte, dxImg \[\]byte, isToggle bool, tapped func\(\), tappedSecondary func\(\)\) \*PicButton](<#NewPicButton>)
  - [func NewPicButtonEx\(uImg \[\]byte, dImg \[\]byte, uxImg \[\]byte, dxImg \[\]byte, isToggle, hasPadding bool, buttonMask desktop.MouseButton, tapped func\(\), tappedSecondary func\(\)\) \*PicButton](<#NewPicButtonEx>)
  - [func \(p \*PicButton\) CreateRenderer\(\) fyne.WidgetRenderer](<#PicButton.CreateRenderer>)
  - [func \(p \*PicButton\) Cursor\(\) desktop.Cursor](<#PicButton.Cursor>)
  - [func \(p \*PicButton\) FocusGained\(\)](<#PicButton.FocusGained>)
  - [func \(p \*PicButton\) FocusLost\(\)](<#PicButton.FocusLost>)
  - [func \(p \*PicButton\) GetLastKeyModifier\(\) fyne.KeyModifier](<#PicButton.GetLastKeyModifier>)
  - [func \(p \*PicButton\) GetLastMouseButton\(\) desktop.MouseButton](<#PicButton.GetLastMouseButton>)
  - [func \(b \*PicButton\) IsDown\(\) bool](<#PicButton.IsDown>)
  - [func \(b \*PicButton\) IsEnabled\(\) bool](<#PicButton.IsEnabled>)
  - [func \(p \*PicButton\) MouseDown\(ev \*desktop.MouseEvent\)](<#PicButton.MouseDown>)
  - [func \(p \*PicButton\) MouseIn\(ev \*desktop.MouseEvent\)](<#PicButton.MouseIn>)
  - [func \(p \*PicButton\) MouseMoved\(ev \*desktop.MouseEvent\)](<#PicButton.MouseMoved>)
  - [func \(p \*PicButton\) MouseOut\(\)](<#PicButton.MouseOut>)
  - [func \(p \*PicButton\) MouseUp\(ev \*desktop.MouseEvent\)](<#PicButton.MouseUp>)
  - [func \(p \*PicButton\) SetDImg\(dImg \[\]byte\) error](<#PicButton.SetDImg>)
  - [func \(b \*PicButton\) SetDown\(bDown bool\)](<#PicButton.SetDown>)
  - [func \(p \*PicButton\) SetDxImg\(dxImg \[\]byte\) error](<#PicButton.SetDxImg>)
  - [func \(b \*PicButton\) SetEnabled\(bEnabled bool\)](<#PicButton.SetEnabled>)
  - [func \(p \*PicButton\) SetMinSize\(minSize fyne.Size\)](<#PicButton.SetMinSize>)
  - [func \(p \*PicButton\) SetUImg\(uImg \[\]byte\) error](<#PicButton.SetUImg>)
  - [func \(p \*PicButton\) SetUxImg\(uxImg \[\]byte\) error](<#PicButton.SetUxImg>)
  - [func \(p \*PicButton\) Tapped\(ev \*fyne.PointEvent\)](<#PicButton.Tapped>)
  - [func \(p \*PicButton\) TappedSecondary\(ev \*fyne.PointEvent\)](<#PicButton.TappedSecondary>)
  - [func \(p \*PicButton\) TypedKey\(ev \*fyne.KeyEvent\)](<#PicButton.TypedKey>)
  - [func \(p \*PicButton\) TypedRune\(r rune\)](<#PicButton.TypedRune>)
- [type PicButtonRenderer](<#PicButtonRenderer>)
  - [func \(r \*PicButtonRenderer\) Destroy\(\)](<#PicButtonRenderer.Destroy>)
  - [func \(r \*PicButtonRenderer\) Layout\(size fyne.Size\)](<#PicButtonRenderer.Layout>)
  - [func \(r \*PicButtonRenderer\) MinSize\(\) fyne.Size](<#PicButtonRenderer.MinSize>)
  - [func \(r \*PicButtonRenderer\) Objects\(\) \[\]fyne.CanvasObject](<#PicButtonRenderer.Objects>)
  - [func \(r \*PicButtonRenderer\) Refresh\(\)](<#PicButtonRenderer.Refresh>)


<a name="PicButton"></a>
## type PicButton



```go
type PicButton struct {
    widget.BaseWidget

    OnTapped          func()
    OnTappedSecondary func()
    // contains filtered or unexported fields
}
```

<a name="NewPicButton"></a>
### func NewPicButton

```go
func NewPicButton(uImg []byte, dImg []byte, uxImg []byte, dxImg []byte, isToggle bool, tapped func(), tappedSecondary func()) *PicButton
```

creates a new picture button widget. At least uImg and dImg must be given

<a name="NewPicButtonEx"></a>
### func NewPicButtonEx

```go
func NewPicButtonEx(uImg []byte, dImg []byte, uxImg []byte, dxImg []byte, isToggle, hasPadding bool, buttonMask desktop.MouseButton, tapped func(), tappedSecondary func()) *PicButton
```

creates a new picture button widget. At least uImg and dImg must be given this function has 2 more parameters than NewPicButton you can switch off the padding and you can define a cutom MouseButtonMask \(if you want to use tertiray Mouse button e.g.\) If buttonMask is 0 then MouseButtonPrimary / MouseButtonSecondary is used automatically based on tapped \!= nil and tappedSecondary \!= nil

<a name="PicButton.CreateRenderer"></a>
### func \(\*PicButton\) CreateRenderer

```go
func (p *PicButton) CreateRenderer() fyne.WidgetRenderer
```

Widget interface

<a name="PicButton.Cursor"></a>
### func \(\*PicButton\) Cursor

```go
func (p *PicButton) Cursor() desktop.Cursor
```

Cursorable interface

<a name="PicButton.FocusGained"></a>
### func \(\*PicButton\) FocusGained

```go
func (p *PicButton) FocusGained()
```

Focusable interface

<a name="PicButton.FocusLost"></a>
### func \(\*PicButton\) FocusLost

```go
func (p *PicButton) FocusLost()
```

Focusable interface

<a name="PicButton.GetLastKeyModifier"></a>
### func \(\*PicButton\) GetLastKeyModifier

```go
func (p *PicButton) GetLastKeyModifier() fyne.KeyModifier
```

Get the last keyboard modifier

<a name="PicButton.GetLastMouseButton"></a>
### func \(\*PicButton\) GetLastMouseButton

```go
func (p *PicButton) GetLastMouseButton() desktop.MouseButton
```

Get the last mouse button

<a name="PicButton.IsDown"></a>
### func \(\*PicButton\) IsDown

```go
func (b *PicButton) IsDown() bool
```

Checks if the button \(used in toggle mode\) is down

<a name="PicButton.IsEnabled"></a>
### func \(\*PicButton\) IsEnabled

```go
func (b *PicButton) IsEnabled() bool
```

Checks if the button is enabled

<a name="PicButton.MouseDown"></a>
### func \(\*PicButton\) MouseDown

```go
func (p *PicButton) MouseDown(ev *desktop.MouseEvent)
```

Mouseable interface

<a name="PicButton.MouseIn"></a>
### func \(\*PicButton\) MouseIn

```go
func (p *PicButton) MouseIn(ev *desktop.MouseEvent)
```

Hoverable interface

<a name="PicButton.MouseMoved"></a>
### func \(\*PicButton\) MouseMoved

```go
func (p *PicButton) MouseMoved(ev *desktop.MouseEvent)
```

Hoverable interface

<a name="PicButton.MouseOut"></a>
### func \(\*PicButton\) MouseOut

```go
func (p *PicButton) MouseOut()
```

Hoverable interface

<a name="PicButton.MouseUp"></a>
### func \(\*PicButton\) MouseUp

```go
func (p *PicButton) MouseUp(ev *desktop.MouseEvent)
```

Mouseable interface

<a name="PicButton.SetDImg"></a>
### func \(\*PicButton\) SetDImg

```go
func (p *PicButton) SetDImg(dImg []byte) error
```

Sets a new dImg

<a name="PicButton.SetDown"></a>
### func \(\*PicButton\) SetDown

```go
func (b *PicButton) SetDown(bDown bool)
```

Sets the button in down state \- no Tapped event is triggered

<a name="PicButton.SetDxImg"></a>
### func \(\*PicButton\) SetDxImg

```go
func (p *PicButton) SetDxImg(dxImg []byte) error
```

Sets a new dxImg

<a name="PicButton.SetEnabled"></a>
### func \(\*PicButton\) SetEnabled

```go
func (b *PicButton) SetEnabled(bEnabled bool)
```

Sets the button enabled / disabled

<a name="PicButton.SetMinSize"></a>
### func \(\*PicButton\) SetMinSize

```go
func (p *PicButton) SetMinSize(minSize fyne.Size)
```

Override the automatic from uImg derived minSize

<a name="PicButton.SetUImg"></a>
### func \(\*PicButton\) SetUImg

```go
func (p *PicButton) SetUImg(uImg []byte) error
```

Sets a new uImg

<a name="PicButton.SetUxImg"></a>
### func \(\*PicButton\) SetUxImg

```go
func (p *PicButton) SetUxImg(uxImg []byte) error
```

Sets a new uxImg

<a name="PicButton.Tapped"></a>
### func \(\*PicButton\) Tapped

```go
func (p *PicButton) Tapped(ev *fyne.PointEvent)
```

Tappable interface

<a name="PicButton.TappedSecondary"></a>
### func \(\*PicButton\) TappedSecondary

```go
func (p *PicButton) TappedSecondary(ev *fyne.PointEvent)
```



<a name="PicButton.TypedKey"></a>
### func \(\*PicButton\) TypedKey

```go
func (p *PicButton) TypedKey(ev *fyne.KeyEvent)
```

Focusable interface

<a name="PicButton.TypedRune"></a>
### func \(\*PicButton\) TypedRune

```go
func (p *PicButton) TypedRune(r rune)
```

Focusable interface

<a name="PicButtonRenderer"></a>
## type PicButtonRenderer

PicButtonRenderer implements:

- fyne.WidgetRenderer

```go
type PicButtonRenderer struct {
    // contains filtered or unexported fields
}
```

<a name="PicButtonRenderer.Destroy"></a>
### func \(\*PicButtonRenderer\) Destroy

```go
func (r *PicButtonRenderer) Destroy()
```

WidgetRenderer interface

<a name="PicButtonRenderer.Layout"></a>
### func \(\*PicButtonRenderer\) Layout

```go
func (r *PicButtonRenderer) Layout(size fyne.Size)
```

WidgetRenderer interface

<a name="PicButtonRenderer.MinSize"></a>
### func \(\*PicButtonRenderer\) MinSize

```go
func (r *PicButtonRenderer) MinSize() fyne.Size
```

WidgetRenderer interface

<a name="PicButtonRenderer.Objects"></a>
### func \(\*PicButtonRenderer\) Objects

```go
func (r *PicButtonRenderer) Objects() []fyne.CanvasObject
```

WidgetRenderer interface

<a name="PicButtonRenderer.Refresh"></a>
### func \(\*PicButtonRenderer\) Refresh

```go
func (r *PicButtonRenderer) Refresh()
```

WidgetRenderer interface

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

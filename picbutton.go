// Implements a button (normal push button or toggle button) with
// customer provided pictures for up, down and inactive state
// if inactive pictures are missing they will be generated from the provided up and down pictures.
// Pictures can be changed on the fly.
// You can also specify which mouse button can be used to press / toggle the button.
// Also the keyboard keyState and used Mouse button can be retrieved for implementing click + Ctrl
// or right click + Shift.
//
// Author: Reiner Pröls
// Licence: MIT

package picbutton

import (
	"bytes"
	"errors"
	"image"
	"image/png"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type noCopy struct{}

var (
	_ fyne.Widget        = (*PicButton)(nil)
	_ fyne.Tappable      = (*PicButton)(nil)
	_ desktop.Mouseable  = (*PicButton)(nil)
	_ desktop.Hoverable  = (*PicButton)(nil)
	_ desktop.Cursorable = (*PicButton)(nil)
	_ fyne.Focusable     = (*PicButton)(nil)
	_ fyne.Tappable      = (*PicButton)(nil)

	_ fyne.WidgetRenderer = (*PicButtonRenderer)(nil)
)

// Button (normal push button or toggle button) with pictures for up, down and inactive state
// Implements
//   - fyne.Widget
//   - fyne.Tappable
//   - fyne.Focusable
//   - desktop.Mouseable
//   - desktop.Hoverable
//   - desktop.Cursorable

type PicButton struct {
	noCopy noCopy // so `go vet` can complain if a widget is passed by value (copied)

	widget.BaseWidget
	img_u           fyne.Resource
	img_d           fyne.Resource
	img_ux          fyne.Resource
	img_dx          fyne.Resource
	minSize         fyne.Size
	buttonMask      desktop.MouseButton
	lastKeyModifier fyne.KeyModifier
	lastMouseButton desktop.MouseButton
	img_ux_created  bool
	img_dx_created  bool

	isEnabled         bool
	isDown            bool
	stateIsDown       bool
	mouseDowninButton bool
	isToggle          bool

	onTapped func()
}

// creates a new picture button widget.
// At least uImg and dImg must be given
// If buttonMask is 0 then MouseButtonPrimary is used
func NewPicButton(uImg []byte, dImg []byte, uxImg []byte, dxImg []byte, isToggle bool, buttonMask desktop.MouseButton, tapped func()) *PicButton {
	if dImg == nil || uImg == nil {
		return nil
	}
	res_u := fyne.NewStaticResource("u_pic", uImg)
	res_d := fyne.NewStaticResource("d_pic", dImg)
	var res_ux, res_dx *fyne.StaticResource
	var ux_created bool
	var dx_created bool

	if uxImg == nil {
		uxImg = createGray(uImg)
		ux_created = true
	}
	if uxImg == nil {
		return nil
	} else {
		res_ux = fyne.NewStaticResource("ux_pic", uxImg)
	}

	if dxImg == nil {
		dxImg = createGray(dImg)
		dx_created = true
	}
	if dxImg == nil {
		return nil
	} else {
		res_dx = fyne.NewStaticResource("dx_pic", dxImg)
	}

	img, _, err := image.Decode(bytes.NewReader(uImg))
	if err != nil {
		return nil
	}
	bm := buttonMask
	if bm == 0 {
		bm = desktop.MouseButtonPrimary
	}

	w := &PicButton{
		onTapped:       tapped,
		img_u:          res_u,
		img_d:          res_d,
		img_ux:         res_ux,
		img_dx:         res_dx,
		img_ux_created: ux_created,
		img_dx_created: dx_created,

		minSize:    fyne.NewSize(float32(img.Bounds().Dx()), float32(img.Bounds().Dy())),
		isToggle:   isToggle,
		isEnabled:  true,
		buttonMask: bm,
	}
	w.ExtendBaseWidget(w)
	return w
}

// creates a grayscale image (Rec.601)
func createGray(buf []byte) []byte {
	src, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		return nil
	}
	bounds := src.Bounds()
	gray := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray.Set(x, y, src.At(x, y)) // Gray = 0.299 R + 0.587 G + 0.114 B
		}
	}

	buf2 := bytes.Buffer{}
	err = png.Encode(&buf2, gray)
	if err != nil {
		return nil
	}
	return buf2.Bytes()
}

// Widget interface
func (p *PicButton) CreateRenderer() fyne.WidgetRenderer {
	imgu := canvas.NewImageFromResource(p.img_u)
	imgu.FillMode = canvas.ImageFillStretch
	r := &PicButtonRenderer{
		w:    p,
		img:  imgu,
		objs: []fyne.CanvasObject{imgu},
	}
	return r
}

// Tappable interface
func (p *PicButton) Tapped(ev *fyne.PointEvent) {
	if p.isEnabled && ((p.buttonMask & desktop.MouseButtonPrimary) != 0) {
		if p.onTapped != nil {
			p.onTapped()
		}
	}
}

// Mouseable interface
func (p *PicButton) MouseDown(ev *desktop.MouseEvent) {
	if p.isEnabled && ((ev.Button & p.buttonMask) != 0) {
		if p.isToggle {
			p.isDown = !p.isDown
		} else {
			p.isDown = true
		}
		p.stateIsDown = p.isDown
		p.mouseDowninButton = true
		p.Refresh()
	}
}

// Mouseable interface
func (p *PicButton) MouseUp(ev *desktop.MouseEvent) {
	if p.isEnabled && ((ev.Button & p.buttonMask) != 0) {
		if !p.isToggle {
			p.isDown = false
			p.stateIsDown = p.isDown
			p.lastKeyModifier = ev.Modifier
			p.lastMouseButton = ev.Button
			if p.onTapped != nil && ((ev.Button & desktop.MouseButtonPrimary) == 0) {
				p.onTapped()
			}
			p.Refresh()
		} else {
			if p.onTapped != nil {
				p.lastKeyModifier = ev.Modifier
				p.lastMouseButton = ev.Button
				if (ev.Button & desktop.MouseButtonPrimary) == 0 {
					p.onTapped()
				}
			}
		}
		p.mouseDowninButton = false
	}
}

// Hoverable interface
func (p *PicButton) MouseOut() {
	if p.isEnabled {
		if p.mouseDowninButton {
			if p.isDown {
				p.isDown = false
				p.Refresh()
			}
			p.mouseDowninButton = false
		}
	}
}

// Hoverable interface
func (p *PicButton) MouseIn(ev *desktop.MouseEvent) {
}

// Hoverable interface
func (p *PicButton) MouseMoved(ev *desktop.MouseEvent) {
}

// Cursorable interface
func (p *PicButton) Cursor() desktop.Cursor {
	if p.isEnabled {
		return desktop.PointerCursor
	} else {
		return desktop.DefaultCursor
	}
}

// Focusable interface
func (p *PicButton) FocusGained() {
	p.Refresh()
}

// Focusable interface
func (p *PicButton) FocusLost() {
	p.Refresh()
}

// Focusable interface
func (p *PicButton) TypedKey(ev *fyne.KeyEvent) {
	if (ev.Name == fyne.KeyReturn || ev.Name == fyne.KeySpace) &&
		p.isEnabled && ((p.buttonMask & desktop.MouseButtonPrimary) != 0) {
		if p.isToggle {
			p.SetDown(!p.IsDown())
		}
		if p.onTapped != nil {
			p.onTapped()
		}
	}
}

// Focusable interface
func (p *PicButton) TypedRune(r rune) {
}

// User functions

// Sets the button in down state - no Tapped event is triggered
func (b *PicButton) SetDown(bDown bool) {
	if b.isDown != bDown {
		b.isDown = bDown
		b.stateIsDown = true
		b.Refresh()
	}
}

// Checks if the button (used in toggle mode) is down
func (b *PicButton) IsDown() bool {
	return b.isDown
}

// Sets the button enabled / disabled
func (b *PicButton) SetEnabled(bEnabled bool) {
	if b.isEnabled != bEnabled {
		b.isEnabled = bEnabled
		b.Refresh()
	}
}

// Checks if the button is enabled
func (b *PicButton) IsEnabled() bool {
	return b.isEnabled
}

// Get the last keyboard modifier
func (p *PicButton) GetLastKeyModifier() fyne.KeyModifier {
	return p.lastKeyModifier
}

// Get the last mouse button
func (p *PicButton) GetLastMouseButton() desktop.MouseButton {
	return p.lastMouseButton
}

// Override the automatic from uImg derived minSize
func (p *PicButton) SetMinSize(minSize fyne.Size) {
	p.minSize = minSize
	p.Refresh()
}

// Sets a new uImg
func (p *PicButton) SetUImg(uImg []byte) error {
	if uImg == nil {
		return errors.New("invalid ux image")
	}
	if p.img_ux_created {
		xImg := createGray(uImg)
		if xImg == nil {
			return errors.New("unable to create gray ux image")
		}
		p.img_ux = fyne.NewStaticResource("ux_pic", xImg)
	}
	p.img_u = fyne.NewStaticResource("u_pic", uImg)
	p.Refresh()
	return nil
}

// Sets a new dImg
func (p *PicButton) SetDImg(dImg []byte) error {
	if dImg == nil {
		return errors.New("invalid dx image")
	}
	if p.img_dx_created {
		xImg := createGray(dImg)
		if xImg == nil {
			return errors.New("unable to create gray dx image")
		}
		p.img_dx = fyne.NewStaticResource("dx_pic", xImg)
	}
	p.img_d = fyne.NewStaticResource("d_pic", dImg)
	p.Refresh()
	return nil
}

// Sets a new uxImg
func (p *PicButton) SetUxImg(uxImg []byte) error {
	if uxImg == nil {
		xImg := createGray(p.img_u.Content())
		if xImg == nil {
			return errors.New("unable to create gray ux image")
		}
		p.img_ux = fyne.NewStaticResource("ux_pic", xImg)
		p.img_ux_created = true
	} else {
		p.img_ux = fyne.NewStaticResource("ux_pic", uxImg)
		p.img_ux_created = false
	}
	p.Refresh()
	return nil
}

// Sets a new dxImg
func (p *PicButton) SetDxImg(dxImg []byte) error {
	if dxImg == nil {
		xImg := createGray(p.img_d.Content())
		if xImg == nil {
			return errors.New("unable to create gray dx image")
		}
		p.img_dx = fyne.NewStaticResource("dx_pic", xImg)
		p.img_dx_created = true
	} else {
		p.img_dx = fyne.NewStaticResource("dx_pic", dxImg)
		p.img_dx_created = false
	}
	p.Refresh()
	return nil
}

// PicButtonRenderer implements:
//   - fyne.WidgetRenderer
type PicButtonRenderer struct {
	w    *PicButton
	img  *canvas.Image
	objs []fyne.CanvasObject
}

// WidgetRenderer interface
func (r *PicButtonRenderer) Layout(size fyne.Size) {
	r.img.Resize(fyne.NewSize(size.Width, size.Height))
	r.img.Move(fyne.NewPos(0, 0))
}

// WidgetRenderer interface
func (r *PicButtonRenderer) MinSize() fyne.Size {
	return r.w.minSize
}

// WidgetRenderer interface
func (r *PicButtonRenderer) Refresh() {
	if r.w.isEnabled {
		if r.w.isDown {
			r.img.Resource = r.w.img_d
		} else {
			r.img.Resource = r.w.img_u
		}
	} else {
		if r.w.isDown {
			r.img.Resource = r.w.img_dx
		} else {
			r.img.Resource = r.w.img_ux
		}
	}
	r.objs = []fyne.CanvasObject{r.img}
	r.img.Refresh()
}

// WidgetRenderer interface
func (r *PicButtonRenderer) Destroy() {
}

// WidgetRenderer interface
func (r *PicButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objs
}

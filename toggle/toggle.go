package toggle

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

const (
	defaultImgWidth     = 40
	defaultImgHeight    = 40
	defaultImagePadding = 5
)

type ToggleOpt = func(s *Toggle) error

type Toggle struct {
	widget.DisableableWidget

	Checked bool

	OnChanged func(bool)

	imgWidth, imgHeight float32
	padding             float32
}

func (t *Toggle) Cursor() desktop.Cursor {
	if t.Disabled() {
		return desktop.DefaultCursor
	}
	return desktop.PointerCursor
}

func (t *Toggle) Tapped(event *fyne.PointEvent) {
	if t.Disabled() {
		return
	}
	t.Checked = !t.Checked
	t.Refresh()
	t.OnChanged(t.Checked)
}

func (t *Toggle) Disable() {
	t.Checked = false
	t.DisableableWidget.Disable()
}

func (t *Toggle) CreateRenderer() fyne.WidgetRenderer {
	img := canvas.NewImageFromResource(offResource)
	img.Resize(fyne.NewSize(t.imgWidth, t.imgHeight))
	img.ScaleMode = canvas.ImageScaleSmooth
	img.FillMode = canvas.ImageFillOriginal
	return &toggleRenderer{toggle: t, img: img}
}

func NewToggle(changed func(bool), opts ...ToggleOpt) *Toggle {
	t := &Toggle{
		DisableableWidget: widget.DisableableWidget{},
		OnChanged:         changed,
		imgWidth:          defaultImgWidth,
		imgHeight:         defaultImgHeight,
		padding:           defaultImagePadding,
	}

	for _, opt := range opts {
		opt(t)
	}

	t.ExtendBaseWidget(t)
	t.Refresh()
	return t
}

type toggleRenderer struct {
	toggle *Toggle

	img *canvas.Image
}

func (tr *toggleRenderer) Destroy() {

}

func (tr *toggleRenderer) Layout(size fyne.Size) {
	imgSize := fyne.NewSize(tr.toggle.imgWidth, tr.toggle.imgHeight)
	pos := fyne.NewPos(size.Width/2-imgSize.Width/2, size.Height/2-imgSize.Height/2)

	tr.img.Move(pos)
	tr.img.Resize(imgSize)
}

func (tr *toggleRenderer) MinSize() fyne.Size {
	return fyne.NewSize(tr.toggle.imgWidth+tr.toggle.padding, tr.toggle.imgHeight+tr.toggle.padding)
}

func (tr *toggleRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{tr.img}
}

func (tr *toggleRenderer) Refresh() {
	if tr.toggle.Disabled() {
		tr.img.Resource = disabledResource
	} else if tr.toggle.Checked {
		tr.img.Resource = onResource
	} else {
		tr.img.Resource = offResource
	}
	tr.img.Refresh()
}

func SetImageWidth(width float32) ToggleOpt {
	return func(t *Toggle) error {
		t.imgWidth = width
		return nil
	}
}

func SetImageHeight(height float32) ToggleOpt {
	return func(t *Toggle) error {
		t.imgHeight = height
		return nil
	}
}

func SetImagePadding(padding float32) ToggleOpt {
	return func(t *Toggle) error {
		t.padding = padding
		return nil
	}
}

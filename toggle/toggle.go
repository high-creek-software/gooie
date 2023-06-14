package toggle

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Toggle struct {
	widget.DisableableWidget

	Checked bool

	OnChanged func(bool)
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
	img.Resize(fyne.NewSize(40, 20))
	img.ScaleMode = canvas.ImageScaleSmooth
	img.FillMode = canvas.ImageFillOriginal
	return &toggleRenderer{toggle: t, img: img}
}

func NewToggle(changed func(bool)) *Toggle {
	t := &Toggle{
		DisableableWidget: widget.DisableableWidget{},
		OnChanged:         changed,
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
	imgSize := tr.img.MinSize()
	pos := fyne.NewPos(theme.Padding(), theme.Padding())

	tr.img.Move(pos)
	tr.img.Resize(imgSize)
}

func (tr *toggleRenderer) MinSize() fyne.Size {
	return fyne.NewSize(40, 20)
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

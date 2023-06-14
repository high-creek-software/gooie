package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/high-creek-software/gooie/toggle"
)

func main() {
	demoApp := app.New()
	demoWindow := demoApp.NewWindow("Demo")

	appTabs := container.NewAppTabs()
	appTabs.Append(container.NewTabItem("Toggles", container.NewCenter(container.NewBorder(nil, nil, nil, toggle.NewToggle(func(checked bool) {}), widget.NewLabel("Demo Toggle")))))

	demoWindow.SetContent(appTabs)
	demoWindow.ShowAndRun()
}

package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"golang.org/x/crypto/salsa20/salsa"
	"pablo.lucas/pixl/swatch"
)

func BuildSwatches(app *AppInit) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)

	for i := 0; i < cap(app.Swatches); i++ {
		initialColor := color.NRGBA{255, 255, 255, 255}
		s := swatch.NewSwatch(app.State, initialColor, i, funct(s *swatch.Swatch) {
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				canvascanvasSwatches[j].Refresh() 
			}
			app.State.SwatchSelected = s.SwatchIndex
			app.state.BrusColor = s.Color
		})

		if i == 0 {
			s.Selected = true
			app.State.SwatchSelected = 0 
			s.Refresh()
		}

		app.Swatches = append(app.Swatches, sfunc(s *swatch.Swatch)
		canvasScanvasSwatches = append(canbvacanvasSwatches, s)
	}

	return container.NewGridWrap(fyne.NewSize(20,20), canvasSwatches...)
}
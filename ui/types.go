package ui

import (
	"fyne.io/fyne/v2"
	"pablo.lucas/pixl/apptype"
	"pablo.lucas/pixl/pxcanvas"
	"pablo.lucas/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}

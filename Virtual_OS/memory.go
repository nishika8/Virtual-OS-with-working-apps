package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showMemoryGame() {
	w := myApp.NewWindow("Memory-Card-Game ")
	w.Resize(fyne.NewSize(250, 400))

	attempt := "0"
	score := "0"
	hs := "0"
	Attempt := widget.NewLabel("Attempts : " + attempt)
	Score := widget.NewLabel("Score : " + score)
	High_Score := widget.NewLabel("High Score : " + hs)

	Reset := widget.NewButton("RESET", func() {})
	z00 := widget.NewButton("", func() {})
	z01 := widget.NewButton("", func() {})
	z02 := widget.NewButton("", func() {})
	z03 := widget.NewButton("", func() {})
	z04 := widget.NewButton("", func() {})
	z10 := widget.NewButton("", func() {})
	z11 := widget.NewButton("", func() {})
	z12 := widget.NewButton("", func() {})
	z13 := widget.NewButton("", func() {})
	z14 := widget.NewButton("", func() {})
	z20 := widget.NewButton("", func() {})
	z21 := widget.NewButton("", func() {})
	z22 := widget.NewButton("", func() {})
	z23 := widget.NewButton("", func() {})
	z24 := widget.NewButton("", func() {})
	z30 := widget.NewButton("", func() {})
	z31 := widget.NewButton("", func() {})
	z32 := widget.NewButton("", func() {})
	z33 := widget.NewButton("", func() {})
	z34 := widget.NewButton("", func() {})

	w.SetContent(container.NewVBox(
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(4, Attempt, Score, High_Score, Reset),
			container.NewGridWithColumns(5, z00, z01, z02, z03, z04),
			container.NewGridWithColumns(5, z10, z11, z12, z13, z14),
			container.NewGridWithColumns(5, z20, z21, z22, z23, z24),
			container.NewGridWithColumns(5, z30, z31, z32, z33, z34),
		),
	))

	w.Show()
}

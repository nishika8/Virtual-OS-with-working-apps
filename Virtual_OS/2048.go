package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func show2048() {

	w := myApp.NewWindow("2048 ")
	w.Resize(fyne.NewSize(320, 400))

	zz00 := "0"
	zz01 := "0"
	zz02 := "0"
	zz03 := "0"
	zz10 := "0"
	zz11 := "0"
	zz12 := "0"
	zz13 := "0"
	zz20 := "0"
	zz21 := "0"
	zz22 := "2"
	zz23 := "0"
	zz30 := "0"
	zz31 := "2"
	zz32 := "0"
	zz33 := "0"
	HS := "0"
	S := "0"
	z00 := widget.NewLabel(zz00)
	z01 := widget.NewLabel(zz01)
	z02 := widget.NewLabel(zz02)
	z03 := widget.NewLabel(zz03)
	z11 := widget.NewLabel(zz10)
	z12 := widget.NewLabel(zz11)
	z13 := widget.NewLabel(zz12)
	z10 := widget.NewLabel(zz13)
	z20 := widget.NewLabel(zz20)
	z21 := widget.NewLabel(zz21)
	z22 := widget.NewLabel(zz22)
	z23 := widget.NewLabel(zz23)
	z30 := widget.NewLabel(zz30)
	z31 := widget.NewLabel(zz31)
	z32 := widget.NewLabel(zz32)
	z33 := widget.NewLabel(zz33)
	hs := widget.NewLabel("HIGH SCORE : " + HS)
	s := widget.NewLabel("SCORE : " + S)
	btn_restart := widget.NewButton("RESTART", func() {
		z := "0"
		t := "2"
		z00.SetText(z)
		z01.SetText(z)
		z02.SetText(z)
		z03.SetText(z)
		z11.SetText(z)
		z12.SetText(z)
		z13.SetText(z)
		z10.SetText(z)
		z20.SetText(z)
		z21.SetText(z)
		z22.SetText(t)
		z23.SetText(z)
		z30.SetText(z)
		z31.SetText(t)
		z32.SetText(z)
		z33.SetText(z)
	})

	btn_up := widget.NewButton("^", func() {

	})
	btn_left := widget.NewButton("<", func() {

	})
	btn_down := widget.NewButton("v", func() {

	})
	btn_right := widget.NewButton(">", func() {

	})

	w.SetContent(container.NewVBox(
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(3, hs, s, btn_restart),
			container.NewGridWithColumns(4, btn_left, btn_up, btn_down, btn_right),
			container.NewGridWithColumns(4, z00, z01, z02, z03),
			container.NewGridWithColumns(4, z10, z11, z12, z13),
			container.NewGridWithColumns(4, z20, z21, z22, z23),
			container.NewGridWithColumns(4, z30, z31, z32, z33),
		),
	))

	w.Show()
}

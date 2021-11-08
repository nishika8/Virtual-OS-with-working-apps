package main

import (
	//"fyne.io/fyne"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()

var myWindow fyne.Window = myApp.NewWindow("My OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget
var btn7 fyne.Widget
var img fyne.CanvasObject

var DeskBtn fyne.Widget
var panelContent *fyne.Container

func main() {
	myApp.Settings().SetTheme(theme.DarkTheme())
	img = canvas.NewImageFromFile("C:\\Users\\hp\\Desktop\\Pepc\\Virtual_OS\\4.jpg")
	// img.FillMode = canvas.ImageFillOriginal
	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		showWeatherApp(myWindow)
	})

	btn2 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		showCalc()
	})

	btn3 = widget.NewButtonWithIcon("Gallery", theme.StorageIcon(), func() {
		showGallery()
	})

	btn4 = widget.NewButtonWithIcon("Text Editor", theme.DocumentIcon(), func() {
		showEditor()
	})

	btn5 = widget.NewButtonWithIcon("Audio Player", theme.FileAudioIcon(), func() {
		showAudioPlayer()
	})

	btn6 = widget.NewButtonWithIcon("2048", theme.NavigateNextIcon(), func() {
		show2048()
	})

	btn7 = widget.NewButtonWithIcon("Memory Game", theme.SearchIcon(), func() {
		showMemoryGame()
	})
	DeskBtn = widget.NewButtonWithIcon("Desktop", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox(container.NewGridWithColumns(8, DeskBtn, btn1, btn2, btn3, btn4, btn5, btn6, btn7))

	myWindow.Resize(fyne.NewSize(700, 700))
	//myWindow.CentreOnScreen()

	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)

	myWindow.ShowAndRun()
}

package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var count int = 1

func showEditor() {
	a := app.New()

	w := a.NewWindow("Text Editor")

	w.Resize(fyne.NewSize(600, 600))

	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("My Text Editor"),
		),
	)

	content.Add(widget.NewButton("Add New File ", func() {
		content.Add(widget.NewLabel("New File " + strconv.Itoa(count)))
		count++
	}))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter Text ... ")
	input.Resize(fyne.NewSize(400, 400))

	saveBtn := widget.NewButton("Save text file", func() {
		saveFileDialogue := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)

				uc.Write(textData)
			}, w)
		saveFileDialogue.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")

		saveFileDialogue.Show()
	})

	openBtn := widget.NewButton("Open File ", func() {
		openFileDialogue := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				ReadData, _ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("New File", ReadData)

				viewData := widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName))
				w.SetContent(container.NewScroll(viewData))

				w.Resize(fyne.NewSize(400, 400))
				w.Show()

			}, w)
		//openFileDialogue.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openFileDialogue.Show()
	})

	w.SetContent(
		container.NewVBox(
			content,
			input,
			container.NewHBox(
				saveBtn,
				openBtn,
			),
		),
	)
	w.Show()
}

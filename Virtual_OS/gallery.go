package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/theme"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
)

func showGallery() {
	//a := app.New()
	w := myApp.NewWindow("Gallery")
	w.Resize(fyne.NewSize(720, 720))
	root_src := "C:\\Users\\hp\\Pictures\\Saved Pictures"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}
	tabs := container.NewAppTabs()

	for _, file := range files {
		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpg" {
				//fmt.Println(picsArr)
				image := canvas.NewImageFromFile(root_src + "\\" + file.Name())
				image.FillMode = canvas.ImageFillOriginal
				tabs.Append(container.NewTabItem(file.Name(), image))

			}
		}
	}

	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(container.NewVBox(tabs))
	w.Show()
}

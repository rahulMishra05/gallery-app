package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fmt"
	"io/ioutil"
	"log"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Gallery")
	w.Resize(fyne.NewSize(990, 660))

	root_src := "C:\\Users\\91982\\Pictures"

	files, err := ioutil.ReadDir(root_src)
	
	if err != nil {
        log.Fatal(err)
    }

	tabs := container.NewAppTabs()
	
	for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir(){
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpeg"{
				image := canvas.NewImageFromFile(root_src+"\\"+file.Name())
				image.FillMode = canvas.ImageFillOriginal
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
    }



	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)

	w.ShowAndRun()
}


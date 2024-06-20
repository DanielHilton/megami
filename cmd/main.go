package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/DanielHilton/megami/internal/domain/screen"
)

func main() {
	a := app.New()

	w := a.NewWindow("Megami - MyAnimeList GUI")
	w.Resize(fyne.NewSize(800, 600))

	s := screen.NewSplash(func() {
		fmt.Println("Login")
	})

	w.SetContent(s.Render())

	w.ShowAndRun()
	teardown()
}

func teardown() {
	println("まったね！")
}

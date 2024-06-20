package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"os"
)

type SplashConfig struct {
	LoginButtonFunc func()
}

func NewSplash(loginButtonFunc func()) *SplashConfig {
	return &SplashConfig{LoginButtonFunc: loginButtonFunc}
}

func (s *SplashConfig) Render() *fyne.Container {
	title := canvas.NewText("Megami", color.White)
	title.TextSize = 58
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	subtitle := canvas.NewText("MyAnimeList GUI", color.White)
	subtitle.TextSize = 24
	subtitle.Alignment = fyne.TextAlignCenter
	subtitle.TextStyle = fyne.TextStyle{Italic: true}

	c := container.NewCenter(
		container.NewVBox(
			title,
			subtitle,
			canvas.NewText(" ", color.Transparent),
			container.NewHBox(
				widget.NewButton("Login", s.LoginButtonFunc),
				layout.NewSpacer(),
				widget.NewButton("Exit", func() {
					os.Exit(0)
				}),
			),
		),
	)

	return c
}

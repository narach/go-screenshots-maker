package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/kbinani/screenshot"
)

func makeScreenshot(displayIndex int, dirName string) {
	bounds := screenshot.GetDisplayBounds(displayIndex)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	fileName := getScreenshotName()
	os.Chdir(dirName)
	file, _ := os.Create(fileName)
	defer file.Close()

	png.Encode(file, img)

	fmt.Printf("Display #%d : %v \"%s\"\n", displayIndex, bounds, fileName)
}

func getScreenshotName() string {
	curTime := time.Now()
	fName := fmt.Sprintf("%02d-%02d-%02d.png", curTime.Hour(), curTime.Minute(), curTime.Second())
	return fName
}

func startScreenCapturing() {
	// n := screenshot.NumActiveDisplays()
	t := time.Now()
	dirName := fmt.Sprintf("screenshots-%d-%02d-%02d-%02d-%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	os.Mkdir(dirName, os.ModePerm)
	tick := time.Tick(5 * time.Second)
	for range tick {
		makeScreenshot(0, dirName)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Screen capture")

	label := widget.NewLabel("Screen capture with Fyne!")
	w.SetContent(container.NewVBox(
		label,
		widget.NewButton("Start capture", func() {
			label.SetText("Screen capturing should start here")
		}),
	))

	w.ShowAndRun()
}
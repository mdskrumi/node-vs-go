package main

import (
	"fmt"
	"time"

	"github.com/fogleman/gg"

	_ "image/jpeg"

	_ "image/png"
)

func main() {
	startTime := time.Now()
	processImage("../watermark.png",
		"../input.jpg",
		"../go-output.png")
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Time taken: %s\n", duration)
}

func processImage(watermarkFilePath string, inputFilePath string, outputFilePath string) error {
	inputImage, err := gg.LoadJPG(inputFilePath)
	if err != nil {
		return err
	}
	width := inputImage.Bounds().Dx()
	height := inputImage.Bounds().Dy()
	dc := gg.NewContext(width, height)
	watermarkImage, err := gg.LoadPNG(watermarkFilePath)
	if err != nil {
		return err
	}
	dc.DrawImage(inputImage, 0, 0)
	dc.DrawImage(watermarkImage, 64, 64)
	err = dc.SavePNG(outputFilePath)
	if err != nil {
		return err
	}
	return nil
}

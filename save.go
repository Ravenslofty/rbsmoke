package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func Save(filename string, height, width int, img []color.NRGBA) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Couldn't open file for writing: ", err.Error())
		return
	}

	defer file.Close()

	render := image.NewNRGBA(image.Rect(0, 0, width, height))

	for index, colour := range img {
		point := FlatIndexToPoint(width, index)
		render.SetNRGBA(point.X, point.Y, colour)
	}

	err = png.Encode(file, render)

	if err != nil {
		fmt.Println("Couldn't encode PNG: ", err.Error())
	}
}

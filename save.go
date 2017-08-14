package main

import (
	"fmt"
	"image/png"
	"os"
)

func Save(filename string) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Couldn't open file for writing: ", err.Error())
		return
	}

	defer file.Close()

	err = png.Encode(file, img.SubImage(img.Rect))

	if err != nil {
		fmt.Println("Couldn't encode PNG: ", err.Error())
	}
}

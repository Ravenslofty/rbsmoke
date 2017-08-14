package main

import (
	"image"
)

func Neighbours(pos image.Point) []image.Point {
	var neighbours []image.Point

	for x := -1; x <= +1; x++ {
		for y := -1; y <= +1; y++ {
			new_pt := pos.Add(image.Pt(x, y))
			if !(x == 0 && y == 0) && new_pt.In(img.Rect) {
				neighbours = append(neighbours, new_pt)
			}
		}
	}

	return neighbours
}

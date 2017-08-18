package main

import (
	"image"
)

var neighbour_list [][]image.Point

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

func InitNeighbours() {
	h := *height
	w := *width
	size := h * w

	neighbour_list = make([][]image.Point, size)

	for i, _ := range neighbour_list {
		neighbour_list[i] = Neighbours(FlatIndexToPoint(w, i))
	}
}

func FlatIndexToPoint(width, index int) image.Point {
	return image.Point{index % width, index / width}
}

func PointToFlatIndex(width int, point image.Point) int {
	return point.X + point.Y*width
}

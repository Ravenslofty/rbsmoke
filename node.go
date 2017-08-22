package main

import (
	"image"
)

var neighbour_list [][]int

func Neighbours(point, height, width int) []int {
	var neighbours []int

	pos := FlatIndexToPoint(width, point)

	for x := -1; x <= +1; x++ {
		for y := -1; y <= +1; y++ {
			new_pt := pos.Add(image.Pt(x, y))
			if !(x == 0 && y == 0) && new_pt.X >= 0 && new_pt.X < width && new_pt.Y >= 0 && new_pt.Y < height {
				neighbours = append(neighbours, PointToFlatIndex(width, new_pt))
			}
		}
	}

	return neighbours
}

func InitNeighbours(height, width int) {
	size := height * width

	neighbour_list = make([][]int, size)

	for i, _ := range neighbour_list {
		neighbour_list[i] = Neighbours(i, height, width)
	}
}

func FlatIndexToPoint(width, index int) image.Point {
	return image.Point{index % width, index / width}
}

func PointToFlatIndex(width int, point image.Point) int {
	return point.X + point.Y*width
}

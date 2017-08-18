package main

import (
	"image"
	"image/color"
)

func Select(colour color.Color, unfilled []image.Point) int {
	var best_fitness int32 = 255 * 255 * 3 // Maximum possible RGB difference.
	var best_index int

	for index, point := range unfilled {
		fitness := ColourFitness(colour, point)
		if fitness < best_fitness {
			best_index = index
			best_fitness = fitness
		}
	}

	return best_index
}

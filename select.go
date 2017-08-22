package main

import (
	"image/color"
)

func SelectSmallestDifference(colour color.NRGBA, unfilled []int, width int) int {
	var best_fitness int32 = 255 * 255 * 3 // Maximum possible RGB difference.
	var best_index int

	for index, point := range unfilled {
		fitness := ColourFitness(colour, point, width)
		if fitness < best_fitness {
			best_index = index
			best_fitness = fitness
		}
	}

	return best_index
}

func SelectGreatestDifference(colour color.NRGBA, unfilled []int, width int) int {
	var best_fitness int32
	var best_index int

	for index, point := range unfilled {
		fitness := ColourFitness(colour, point, width)
		if fitness > best_fitness {
			best_index = index
			best_fitness = fitness
		}
	}

	return best_index
}

package main

import (
	"image/color"
)

// Method to select a place for the next colour to go.
// Valid functions are SelectSmallestDifference and SelectGreatestDifference.
// Greatest tends to run slower than smallest.
func Select(colour color.NRGBA, unfilled []int, width int) int {
	return SelectSmallestDifference(colour, unfilled, width)
}

// Method to compute the difference between a colour and its neighbours.
// Valid functions are ColourFitnessMinimum and ColourFitnessSum.
// Minimum produces smooth gradients with sharp edges.
// Sum produces very psychedelic and angular patterns.
func ColourFitness(pixel color.NRGBA, pos, width int) int32 {
	return ColourFitnessMinimum(pixel, pos, width)
}

// Method to compute the difference between two colours.
// Valid functions are ColourDiffRgb and ColourDiffLab.
// Rgb is fastest.
// Lab is slower, but produces smoother gradients.
func ColourDiff(a, b color.NRGBA) int32 {
	return ColourDiffRgb(a, b)
}

// Method to change the order the colours are travelled.
// Valid functions are SortNone, SortHcl and SortHsv.
// None is fastest (obviously).
// Hcl and Hsv produce a smoke-like effect, with different gradients.
func Sort(a []color.NRGBA) {
	SortHsv(a)
}

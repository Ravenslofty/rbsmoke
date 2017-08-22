package main

import (
	"image/color"
)

// Calculate 8-bit colour for limited colour space.
func MakeColour(c, colours int) uint8 {
	return uint8((c * 255) / (colours - 1))
}

func MakeNRGBA(r, g, b, colours int) color.NRGBA {
	return color.NRGBA{MakeColour(r, colours), MakeColour(g, colours),
		MakeColour(b, colours), 255}
}

func ColourDiff(a, b color.NRGBA) int32 {
	rdiff := int32(a.R) - int32(b.R)
	gdiff := int32(a.G) - int32(b.G)
	bdiff := int32(a.B) - int32(b.B)

	return rdiff*rdiff + gdiff*gdiff + bdiff*bdiff
}

func ColourFitness(pixel color.NRGBA, pos, width int) int32 {
	var min_diff int32 = 255 * 255 * 3 // Max RGB difference.

	for _, new_pt := range neighbour_list[pos] {
		diff := ColourDiff(pixel, img[new_pt])
		if diff < min_diff {
			min_diff = diff
		}
	}

	return min_diff
}

func NewColourList(colours int) []color.NRGBA {
	colour_list := make([]color.NRGBA, 0, colours*colours*colours)

	for r := 0; r <= colours; r++ {
		for g := 0; g <= colours; g++ {
			for b := 0; b <= colours; b++ {
				colour_list = append(colour_list, MakeNRGBA(b, g, r, colours))
			}
		}
	}

	return colour_list
}

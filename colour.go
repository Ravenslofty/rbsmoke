package main

import (
	"image"
	"image/color"
)

var fitness []int32
var fitness_ok []bool

// Calculate 8-bit colour for limited colour space.
func MakeColour(c, colours int) uint8 {
	return uint8((c * 255) / (colours - 1))
}

func MakeNRGBA(r, g, b, colours int) color.NRGBA {
	return color.NRGBA{MakeColour(r, colours), MakeColour(g, colours),
		MakeColour(b, colours), 255}
}

func ColourDiff(a, b color.Color) int32 {
	r1, g1, b1, _ := a.RGBA()
	r2, g2, b2, _ := b.RGBA()
	rdiff := int32(r1) - int32(r2)
	gdiff := int32(g1) - int32(g2)
	bdiff := int32(b1) - int32(b2)

	return rdiff*rdiff + gdiff*gdiff + bdiff*bdiff
}

func ColourFitness(pixel color.Color, pos image.Point) int32 {
	idx := PointToFlatIndex(*width, pos)

	if fitness_ok[idx] {
		return fitness[idx]
	}

	var diff int32

	for _, new_pt := range neighbour_list[PointToFlatIndex(*width, pos)] {
		diff += ColourDiff(pixel, img.NRGBAAt(new_pt.X, new_pt.Y))
	}

	fitness[idx] = diff
	fitness_ok[idx] = true

	return diff
}

func NewColourList(colours int) []color.NRGBA {
	colour_list := make([]color.NRGBA, 0, colours*colours*colours)

	for r := 0; r <= colours; r++ {
		for g := 0; g <= colours; g++ {
			for b := 0; b <= colours; b++ {
				colour_list = append(colour_list, MakeNRGBA(r, g, b, colours))
			}
		}
	}

	return colour_list
}

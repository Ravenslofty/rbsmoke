package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"image/color"
)

// Calculate 8-bit colour for limited colour space.
func MakeRGB255(c, colours int) uint8 {
	return uint8((c * 255) / (colours - 1))
}

func MakeNRGBA(r, g, b, colours int) color.NRGBA {
	return color.NRGBA{MakeRGB255(r, colours), MakeRGB255(g, colours),
		MakeRGB255(b, colours), 255}
}

func MakeColorful(a color.NRGBA) colorful.Color {
	return colorful.Color{float64(a.R) / 255.0, float64(a.G) / 255.0, float64(a.B) / 255.0}
}

func ColourDiffRgb(a, b color.NRGBA) int32 {
	rdiff := int32(a.R) - int32(b.R)
	gdiff := int32(a.G) - int32(b.G)
	bdiff := int32(a.B) - int32(b.B)

	return rdiff*rdiff + gdiff*gdiff + bdiff*bdiff
}

func ColourDiffLab(a, b color.NRGBA) int32 {
	// Ugly colour space hacking.

	a_ := MakeColorful(a)
	b_ := MakeColorful(b)

	diff := a_.DistanceLab(b_)

	return int32(65535.0 * diff)
}

func ColourFitnessMinimum(pixel color.NRGBA, pos, width int) int32 {
	var min_diff int32 = 255 * 255 * 3 // Max RGB difference.

	for _, new_pt := range neighbour_list[pos] {
		diff := ColourDiff(pixel, img[new_pt])
		if diff < min_diff {
			min_diff = diff
		}
	}

	return min_diff
}

func ColourFitnessSum(pixel color.NRGBA, pos, width int) int32 {
	var sum_diff int32

	for _, new_pt := range neighbour_list[pos] {
		sum_diff += ColourDiff(pixel, img[new_pt])
	}

	return sum_diff
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

        Sort(colour_list)

	return colour_list
}

package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"image/color"
	"math"
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
	return colorful.Color{R: float64(a.R) / 255.0, G: float64(a.G) / 255.0, B: float64(a.B) / 255.0}
}

func ColourDiffRgb(a, b color.NRGBA) int32 {
	rdiff := int32(a.R) - int32(b.R)
	gdiff := int32(a.G) - int32(b.G)
	bdiff := int32(a.B) - int32(b.B)

	return rdiff*rdiff + gdiff*gdiff + bdiff*bdiff
}

func ColourDiffLab(a, b color.NRGBA) int32 {
	// Ugly colour space hacking.

	// RGB255 -> sRGB -> Linear RGB
	// This step is needed because go-colorful uses floats internally
	// rather than uint8. Additionally, RGB is a relative measurement, so
	// cannot be used directly.
	ar, ag, ab := MakeColorful(a).FastLinearRgb()
	br, bg, bb := MakeColorful(b).FastLinearRgb()

	// Linear RGB -> CIE XYZ
	// Here we transform the relative RGB system to absolute XYZ co-ordinates.
	ax, ay, az := colorful.LinearRgbToXyz(ar, ag, ab)
	bx, by, bz := colorful.LinearRgbToXyz(br, bg, bb)

	// CIE XYZ -> CIE L*a*b*
	// And finally, here we transform absolute XYZ to L*a*b*, which is a
	// perception-based colour space.
	a_l, a_a, a_b := colorful.XyzToLab(ax, ay, az)
	b_l, b_a, b_b := colorful.XyzToLab(bx, by, bz)

	// And finally we can calculate the perceived difference in colour.
	diff := math.Sqrt((a_l-b_l)*(a_l-b_l) + (a_a-b_a)*(a_a-b_a) + (a_b - b_b))

	// Yep, we just went through four colour spaces to get what we wanted.

	return int32(65535.0 * diff)
}

func ColourFitnessMinimum(pixel color.NRGBA, pos, width int, neighbours []int, img []color.NRGBA) int32 {
	var min_diff int32 = 255 * 255 * 3 // Max RGB difference.

	for _, new_pt := range neighbours {
		diff := ColourDiff(pixel, img[new_pt])
		if diff < min_diff {
			min_diff = diff
		}
	}

	return min_diff
}

func ColourFitnessSum(pixel color.NRGBA, pos, width int, neighbours []int, img []color.NRGBA) int32 {
	var sum_diff int32

	for _, new_pt := range neighbours {
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

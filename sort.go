package main

import (
	"image/color"
	"sort"
)

func SortNone(colours []color.NRGBA) {
	// There's nothing to see here.

	// No really, there's nothing here.

	// Why are you still reading this?

	// Are you expecting a "none" sort to involve code?

	// Really?

	// Well, I don't have any code for you. But I do have a joke!

	// A butcher is seven feet tall and wears size 14 shoes. What does he weigh?

	// Sausages.
}

func SortHcl(colours []color.NRGBA) {
	sort.Slice(colours, func(i, j int) bool {
		h1, _, _ := MakeColorful(colours[i]).Hcl()
		h2, _, _ := MakeColorful(colours[j]).Hcl()

		return h1 < h2
	})
}

func SortHsv(colours []color.NRGBA) {
	sort.Slice(colours, func(i, j int) bool {
		h1, _, _ := MakeColorful(colours[i]).Hsv()
		h2, _, _ := MakeColorful(colours[j]).Hsv()

		return h1 < h2
	})
}

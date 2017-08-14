package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"sort"
)

var img image.NRGBA
var unfilled []image.Point
var fitness []int
var fitness_ok []bool

func Save(filename string) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Couldn't open file for writing: ", err.Error())
		return
	}

	defer file.Close()

	err = png.Encode(file, img.SubImage(img.Rect))

	if err != nil {
		fmt.Println("Couldn't encode PNG: ", err.Error())
	}
}

// Based on the Rainbow Smoke algorithm by József Fejes.
func Render(x_size, y_size, colours int) {

	img = *image.NewNRGBA(image.Rect(0, 0, x_size, y_size))

	colour_list := NewColourList(colours)

	start_pt := image.Pt(x_size/2, y_size/2)

	unfilled = make([]image.Point, 0, x_size*y_size)
	unfilled_map := make(map[image.Point]int)
	filled_map := make(map[image.Point]bool)

	fitness = make([]int, x_size*y_size)
	fitness_ok = make([]bool, x_size*y_size)

	for i := 0; i < x_size*y_size; i++ {

		for j := 0; j < x_size*y_size; j++ {
			fitness_ok[j] = false
		}

		if i%256 == 0 {
			fmt.Printf("%d/%d done, slice: %d map: %d\n", i, x_size*y_size,
				len(unfilled), len(unfilled_map))
			Save(fmt.Sprintf("rbsmoke%08d.png", i))
		}

		var curr_pt image.Point

		if len(unfilled) == 0 {
			curr_pt = start_pt
		} else {
			// Expensive!
			sort.Slice(unfilled, func(j, k int) bool {
				return ColourFitness(colour_list[i], unfilled[j]) < ColourFitness(colour_list[i], unfilled[k])
			})
			curr_pt = unfilled[0]

			// Discard point
			unfilled[len(unfilled)-1], unfilled[0] = unfilled[0], unfilled[len(unfilled)-1]
			unfilled = unfilled[:len(unfilled)-1]
			delete(unfilled_map, curr_pt)
			filled_map[curr_pt] = true
		}

		img.SetNRGBA(curr_pt.X, curr_pt.Y, colour_list[i])

		for _, new_pt := range Neighbours(curr_pt) {
			_, present := unfilled_map[new_pt]
			if !present && !filled_map[new_pt] {
				unfilled = append(unfilled, new_pt)
				unfilled_map[new_pt] = len(unfilled) - 1
			}
		}
	}

	Save(fmt.Sprintf("rbsmoke%08d.png", x_size*y_size))

	fmt.Println("Done!")

}

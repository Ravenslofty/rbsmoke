package main

import (
	"fmt"
	"image"
	"sort"
	"time"
)

var img image.NRGBA

// Based on the Rainbow Smoke algorithm by József Fejes.
func Render(x_size, y_size, colours int) {

	img = *image.NewNRGBA(image.Rect(0, 0, x_size, y_size))

	colour_list := NewColourList(colours)

	start_pt := image.Pt(x_size/2, y_size/2)

	unfilled := make([]image.Point, 0, x_size*y_size)
	unfilled_map := make(map[image.Point]bool)
	filled_map := make(map[image.Point]bool)

	fitness = make([]int32, x_size*y_size)
	fitness_ok = make([]bool, x_size*y_size)

        InitNeighbours()

	start_time := time.Now()

	for i := 0; i < x_size*y_size; i++ {

		for j := 0; j < x_size*y_size; j++ {
			fitness_ok[j] = false
		}

		if i%256 == 255 {
			fmt.Printf("%d/%d done, open: %d, speed: %d px per sec\r", i, x_size*y_size,
				len(unfilled), int64(i*int(time.Second))/int64(time.Now().Sub(start_time)))
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
		}

		filled_map[curr_pt] = true

		img.SetNRGBA(curr_pt.X, curr_pt.Y, colour_list[i])

		for _, new_pt := range neighbour_list[PointToFlatIndex(*width, curr_pt)] {
			if !unfilled_map[new_pt] && !filled_map[new_pt] {
				unfilled = append(unfilled, new_pt)
				unfilled_map[new_pt] = true
			}
		}
	}

	Save(fmt.Sprintf("rbsmoke%08d.png", x_size*y_size))

	fmt.Println("Done!")

}

package main

import (
	"fmt"
	"image"
	"image/color"
	"time"
)

var img []color.NRGBA

// Based on the Rainbow Smoke algorithm by József Fejes.
func Render(height, width, colours int) {

	x_size := width
	y_size := height

	img = make([]color.NRGBA, x_size*y_size)

	colour_list := NewColourList(colours)

	start_pt := PointToFlatIndex(width, image.Pt(width/2, height/2))

	unfilled := make([]int, 0, x_size*y_size)
	unfilled_map := make(map[int]bool)
	filled_map := make(map[int]bool)

	InitNeighbours(height, width)

	start_time := time.Now()

	for i := 0; i < x_size*y_size; i++ {

		if i%256 == 255 {
			fmt.Printf("%.2f%%, open: %d, speed: %d px/sec\r", float64(100*i)/float64(x_size*y_size),
				len(unfilled), int64(i*int(time.Second))/int64(time.Now().Sub(start_time)))
			go Save(fmt.Sprintf("rbsmoke%08d.png", i), height, width)
		}

		var curr_pt int

		if len(unfilled) == 0 {
			curr_pt = start_pt
		} else {
			// Expensive!
			curr_pt_index := Select(colour_list[i], unfilled, width)
			curr_pt = unfilled[curr_pt_index]

			// Discard point
			unfilled[len(unfilled)-1], unfilled[curr_pt_index] = unfilled[curr_pt_index], unfilled[len(unfilled)-1]
			unfilled = unfilled[:len(unfilled)-1]
			delete(unfilled_map, curr_pt)
		}

		filled_map[curr_pt] = true

		img[curr_pt] = colour_list[i]

		for _, new_pt := range neighbour_list[curr_pt] {
			if !unfilled_map[new_pt] && !filled_map[new_pt] {
				unfilled = append(unfilled, new_pt)
				unfilled_map[new_pt] = true
			}
		}
	}

	Save(fmt.Sprintf("rbsmoke%08d.png", height*width), height, width)

	fmt.Println("Done!")

}

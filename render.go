package main

import (
    "fmt"
    "image"
    "image/color"
    "image/png"
    "os"
    "sort"
)

var img image.NRGBA
var unfilled []image.Point

// Calculate 8-bit colour for limited colour space.
func MakeColour(c, colours int) uint8 {
    return uint8((c * 255) / (colours - 1))
}

func MakeNRGBA(r, g, b, colours int) color.NRGBA {
    return color.NRGBA{MakeColour(r, colours), MakeColour(g, colours),
                        MakeColour(b, colours), 255}
}

func ColourDiff(a, b color.NRGBA) int {
    rdiff := int(a.R - b.R)
    gdiff := int(a.G - b.G)
    bdiff := int(a.B - b.B)

    return rdiff*rdiff + gdiff*gdiff + bdiff*bdiff
}

func ColourFitness(pixel color.NRGBA, pos image.Point) int {
    var diff int

    for x := -1; x <= +1; x++ {
        for y := -1; y <= +1; y++ {
            new_pt := pos.Add(image.Pt(x, y))
            if !(x == 0 && y == 0) && new_pt.In(img.Rect) {
                diff += ColourDiff(pixel, img.NRGBAAt(new_pt.X, new_pt.Y))
            }
        }
    }

    return diff
}

func NewColourList(colours int) []color.NRGBA {
    colour_list := make([]color.NRGBA, 0, colours*colours*colours)

    fmt.Println("Initialising...")
    for r := 0; r <= colours; r++ {
        for g := 0; g <= colours; g++ {
            for b := 0; b <= colours; b++ {
                colour_list = append(colour_list, MakeNRGBA(r, g, b, colours))
            }
        }
    }

    return colour_list
}

// Based on the Rainbow Smoke algorithm by József Fejes.
func Render(x_size, y_size, colours int) {

    img = *image.NewNRGBA(image.Rect(0, 0, x_size, y_size))

    colour_list := NewColourList(colours)

    start_pt := image.Pt(x_size / 2, y_size / 2)

    unfilled = make([]image.Point, 0, x_size*y_size)
    unfilled_map := make(map[image.Point]int)

    for i := 0; i < x_size*y_size; i++ {
        if i%256 == 0 {
            fmt.Printf("%d/%d done, %d elements in queue\n", i, x_size*y_size,
                    len(unfilled))
        }

        var curr_pt image.Point

        if len(unfilled) == 0 {
            curr_pt = start_pt
        } else {
            // Expensive!
            sort.Slice(unfilled, func(j, k int) bool { return ColourFitness(colour_list[i], unfilled[j]) > ColourFitness(colour_list[i], unfilled[k]) } )
            curr_pt = unfilled[0]

            // Discard point
            unfilled[len(unfilled)-1], unfilled[0] = unfilled[0], unfilled[len(unfilled)-1]
            unfilled = unfilled[:len(unfilled)-1]
            delete(unfilled_map, curr_pt)
        }

        img.SetNRGBA(curr_pt.X, curr_pt.Y, colour_list[i])

        for x := -1; x <= +1; x++ {
            for y := -1; y <= +1; y++ {
                new_pt := curr_pt.Add(image.Pt(x, y))
                _, present := unfilled_map[new_pt]
                if !present && !(x == 0 && y == 0) && new_pt.In(img.Rect) {
                    unfilled = append(unfilled, new_pt)
                    unfilled_map[new_pt] = len(unfilled)-1
                }
            }
        }
    }

    fmt.Println("Done!")

    file, err := os.Create("rbsmoke.png")

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


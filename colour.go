package main

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
    idx := pos.X + (*width)*pos.Y

    if fitness_ok[idx] {
        return fitness[idx]
    }

    var diff int

    for _, new_pt := range(Neighbours(pos)) {
        diff += ColourDiff(pixel, img.NRGBAAt(new_pt.X, new_pt.Y))
    }

    fitness[idx] = diff
    fitness_ok[idx] = true

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


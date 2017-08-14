package main

import (
    "fmt"
    "math"
)

func main() {
    // TODO: make these configurable.
    x_size := 64
    y_size := 64

    // Since we have a single pixel per image, we approximate the number
    // of colours per channel.
    colours := int(math.Ceil(math.Cbrt(float64(x_size) * float64(y_size))))

    fmt.Printf("Going to render a %dx%d image with %d colours per channel\n",
                x_size, y_size, colours)

    Render(x_size, y_size, colours)
}

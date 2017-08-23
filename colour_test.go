package main

import (
    "image/color"
    "math/rand"
    "testing"
)

func RandomColour() (a color.NRGBA) {
    x := rand.Uint32()

    a = color.NRGBA{uint8(x & 255), uint8((x >> 8) & 255), uint8((x >> 16) & 255), 255}
    return
}

func BenchmarkColourDiffRgb(b *testing.B) {
    rand.Seed(1)

    x, y := RandomColour(), RandomColour()

    for n := 0; n < b.N; n++ {
        ColourDiffRgb(x, y)
    }
}

func BenchmarkColourDiffLab(b *testing.B) {
    rand.Seed(1)

    x, y := RandomColour(), RandomColour()

    for n := 0; n < b.N; n++ {
        ColourDiffLab(x, y)
    }
}

func BenchmarkColourDiffLabNoConversion(b *testing.B) {
    rand.Seed(1)

    x, y := MakeColorful(RandomColour()), MakeColorful(RandomColour())

    for n := 0; n < b.N; n++ {
        x.DistanceLab(y)
    }
}



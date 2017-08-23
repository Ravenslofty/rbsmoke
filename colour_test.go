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

var x, y color.NRGBA = RandomColour(), RandomColour()
var bench_result int32

func BenchmarkColourDiffRgb(b *testing.B) {
    var r int32
    for n := 0; n < b.N; n++ {
        r = ColourDiffRgb(x, y)
    }
    bench_result = r
}

func BenchmarkColourDiffLab(b *testing.B) {
    var r int32
    for n := 0; n < b.N; n++ {
        r = ColourDiffLab(x, y)
    }
    bench_result = r
}

func BenchmarkColourDiffLabNoConversion(b *testing.B) {
    _x, _y := MakeColorful(x), MakeColorful(y)
    for n := 0; n < b.N; n++ {
        _x.DistanceLab(_y)
    }
}


package main

import (
    "github.com/lucasb-eyer/go-colorful"
    "image/color"
    "math"
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

func BenchmarkColourDiffRgbInt(b *testing.B) {
    var r int32
    for n := 0; n < b.N; n++ {
        r = ColourDiffRgb(x, y)
    }
    bench_result = r
}

func BenchmarkColourDiffRgbFloat(b *testing.B) {
    var r float64
    _x, _y := MakeColorful(x), MakeColorful(y)
    for n := 0; n < b.N; n++ {
        r = _x.DistanceRgb(_y)
    }
    bench_result = int32(r)
}

func BenchmarkColourDiffLab(b *testing.B) {
    var r int32
    for n := 0; n < b.N; n++ {
        r = ColourDiffLab(x, y)
    }
    bench_result = r
}

func BenchmarkColourDiffLabNoConversion(b *testing.B) {
    var r float64
    _x, _y := MakeColorful(x), MakeColorful(y)
    for n := 0; n < b.N; n++ {
        r = _x.DistanceLab(_y)
    }
    bench_result = int32(r)
}

func sq(x float64) float64 {
    return x*x
}

func BenchmarkColourDiffLabLinearRgb(b *testing.B) {
    var r float64
    _x, _y := MakeColorful(x), MakeColorful(y)
    for n := 0; n < b.N; n++ {
        xr, xg, xb := _x.LinearRgb()
        yr, yg, yb := _y.LinearRgb()
        xx, xy, xz := colorful.LinearRgbToXyz(xr, xg, xb)
        yx, yy, yz := colorful.LinearRgbToXyz(yr, yg, yb)
        l1, a1, b1 := colorful.XyzToLab(xx, xy, xz)
        l2, a2, b2 := colorful.XyzToLab(yx, yy, yz)
        r = math.Sqrt(sq(l1-l2) + sq(a1-a2) + sq(b1-b2))
    }
    bench_result = int32(r)
}

func BenchmarkColourDiffLabFastLinearRgb(b *testing.B) {
    var r float64
    _x, _y := MakeColorful(x), MakeColorful(y)
    for n := 0; n < b.N; n++ {
        xr, xg, xb := _x.FastLinearRgb()
        yr, yg, yb := _y.FastLinearRgb()
        xx, xy, xz := colorful.LinearRgbToXyz(xr, xg, xb)
        yx, yy, yz := colorful.LinearRgbToXyz(yr, yg, yb)
        l1, a1, b1 := colorful.XyzToLab(xx, xy, xz)
        l2, a2, b2 := colorful.XyzToLab(yx, yy, yz)
        r = math.Sqrt(sq(l1-l2) + sq(a1-a2) + sq(b1-b2))
    }
    bench_result = int32(r)
}


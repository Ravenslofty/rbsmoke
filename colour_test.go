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

var bench_result int32

func BenchmarkColourDiffRgbInt(b *testing.B) {
    var r int32
    for n := 0; n < b.N; n++ {
        r = ColourDiffRgb(RandomColour(), RandomColour())
    }
    bench_result = r
}

func BenchmarkColourDiffRgbFloat(b *testing.B) {
    var r float64
    rand.Seed(1)
    for n := 0; n < b.N; n++ {
        r = MakeColorful(RandomColour()).DistanceRgb(MakeColorful(RandomColour()))
    }
    bench_result = int32(r)
}

func BenchmarkColourDiffLab(b *testing.B) {
    var r int32
    rand.Seed(1)
    for n := 0; n < b.N; n++ {
        r = ColourDiffLab(RandomColour(), RandomColour())
    }
    bench_result = r
}

func BenchmarkColourDiffLabNoConversion(b *testing.B) {
    var r float64
    rand.Seed(1)
    for n := 0; n < b.N; n++ {
        r = MakeColorful(RandomColour()).DistanceLab(MakeColorful(RandomColour()))
    }
    bench_result = int32(r)
}

func sq(x float64) float64 {
    return x*x
}

func BenchmarkColourDiffLabLinearRgb(b *testing.B) {
    var r float64
    rand.Seed(1)
    for n := 0; n < b.N; n++ {
        _x, _y := MakeColorful(RandomColour()), MakeColorful(RandomColour())
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
    rand.Seed(1)
    for n := 0; n < b.N; n++ {
        _x, _y := MakeColorful(RandomColour()), MakeColorful(RandomColour())
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

func BenchmarkColourDiffLabFasterLinearRgb(b *testing.B) {
    var r float64
    rand.Seed(1)
    for n := 0; n < b.N; n++ {
        _x, _y := MakeColorful(RandomColour()), MakeColorful(RandomColour())
        xr, xg, xb := sq(_x.R), sq(_x.G), sq(_x.B)
        yr, yg, yb := sq(_y.R), sq(_y.G), sq(_y.B)
        xx, xy, xz := colorful.LinearRgbToXyz(xr, xg, xb)
        yx, yy, yz := colorful.LinearRgbToXyz(yr, yg, yb)
        l1, a1, b1 := colorful.XyzToLab(xx, xy, xz)
        l2, a2, b2 := colorful.XyzToLab(yx, yy, yz)
        r = math.Sqrt(sq(l1-l2) + sq(a1-a2) + sq(b1-b2))
    }
    bench_result = int32(r)
}



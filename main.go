package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
)

var width = flag.Int("w", 64, "Rendered image width")
var height = flag.Int("h", 64, "Rendered image height")

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	x_size := *width
	y_size := *height

	// Since we have a single pixel per image, we approximate the number
	// of colours per channel.
	colours := int(math.Ceil(math.Cbrt(float64(x_size) * float64(y_size))))

	fmt.Printf("Going to render a %dx%d image with %d colours per channel\n",
		x_size, y_size, colours)

	Render(x_size, y_size, colours)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}

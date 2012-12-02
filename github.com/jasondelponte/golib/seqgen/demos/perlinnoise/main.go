package main

import (
	"flag"
	"fmt"
	"github.com/jasondelponte/golib/seqgen/perlinnoise"
)

var height = flag.Int("h", 40, "Height of the display")
var width = flag.Int("w", 100, "Width of the display")
var seed = flag.Int64("s", -1, "Seed")

func main() {
	flag.Parse()

	h := *height
	w := *width

	var pn *perlinnoise.PerlinNoise = nil
	if *seed > 0 {
		pn = perlinnoise.New(*seed)
	} else {
		pn = perlinnoise.NewDefault()
	}

	for i := 0; i < h; i++ { // y
		line := ""
		y := float64(i) / float64(h)
		for j := 0; j < w; j++ { // x
			x := float64(j) / float64(w)

			n := pn.Noise(10*x, 10*y, 0.8)

			if n < 0.35 {
				line += "~"

			} else if n >= 0.35 && n < 0.5 {
				line += "."

			} else if n >= 0.5 && n < 0.6 {
				line += ","

			} else if n >= 0.5 && n < 0.8 {
				line += "#"

			} else {
				line += "S"
			}
		}
		fmt.Println(line)
	}
}

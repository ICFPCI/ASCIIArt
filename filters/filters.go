package filters

import (
	"image/color"
)

func UpsideDown(pixels [][]color.Color) {
	for x := 0; x < len(pixels); x++ {
		line := pixels[x]
		for y := 0; y < len(line)/2; y++ {
			k := len(line) - y - 1
			line[y], line[k] = line[k], line[y]
		}
	}
}

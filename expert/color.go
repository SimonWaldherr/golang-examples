package main

import (
	"fmt"
	"math"
)

func rgb2hsl(r int, g int, b int) (int, int, int) {
	var rf, gf, bf, max, min, l, d, s, h float64

	rf = math.Max(math.Min(float64(r)/255, 1), 0)
	gf = math.Max(math.Min(float64(g)/255, 1), 0)
	bf = math.Max(math.Min(float64(b)/255, 1), 0)
	max = math.Max(rf, math.Max(gf, bf))
	min = math.Min(rf, math.Min(gf, bf))
	l = (max + min) / 2

	if max != min {
		d = max - min
		if l > 0.5 {
			s = d / (2 - max - min)
		} else {
			s = d / (max + min)
		}
		if max == rf {
			if gf < bf {
				h = (gf-bf)/d + 6
			} else {
				h = (gf - bf) / d
			}
		} else if max == gf {
			h = (bf-rf)/d + 2
		} else {
			h = (rf-gf)/d + 4
		}
	} else {
		h = 0
		s = 0
	}

	return int(h * 60), int(s * 100), int(l * 100)
}

func main() {
	fmt.Println(rgb2hsl(121, 167, 22))
	fmt.Println(rgb2hsl(69, 209, 237))
	fmt.Println(rgb2hsl(254, 207, 37))
	fmt.Println(rgb2hsl(122, 167, 255))
	fmt.Println(rgb2hsl(255, 255, 255))
}

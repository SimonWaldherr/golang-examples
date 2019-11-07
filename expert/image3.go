// based on https://github.com/SimonWaldherr/bbmandelbrotGo
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"sync"
)

const (
	maxiteration = 192
)

var (
	zh float64
	zv float64
)

func init() {
	zh = 2.4
	zv = 2.4
}

func abs(z complex128) float64 {
	return math.Hypot(real(z), imag(z))
}

func mandel(c complex128) float64 {
	z := complex128(0)
	for i := 0; i < maxiteration; i++ {
		if abs(z) > 2 {
			return float64(i-1) / maxiteration
		}
		z = z*z + c
	}
	return 0
}

func pixelColor(x, y, width, height uint64, csr, csg, csb int) color.RGBA {
	xf := float64(x)/float64(width)*zv - (zv/2.0 + 0.5)
	yf := float64(y)/float64(height)*zh - (zh / 2.0)
	c := complex(xf, yf)
	calcval := int(mandel(c) * 255)

	return color.RGBA{
		uint8(int(csr) * calcval % 255),
		uint8(int(csg) * calcval % 255),
		uint8(int(csb) * calcval % 255),
		255,
	}
}

// Mandelbrot generates the Mandelbrot picture as *image.RGBA according to the parameters
func Mandelbrot(width, height, cx1, cx2, cy1, cy2 uint64, csr, csg, csb int) (*image.RGBA, string) {
	var wg sync.WaitGroup
	var fullHeight bool

	background := image.Rect(0, 0, int(cx2-cx1), int(cy2-cy1))
	img := image.NewRGBA(background)

	if height == cy2 && cy1 == 0 {
		fullHeight = true
		cy2 = cy2 / 2
	}

	for x := cx1; x < cx2; x++ {
		wg.Add(1)
		go func(x uint64) {
			defer wg.Done()
			if fullHeight {
				for y := cy1; y < cy2+1; y++ {
					colval := pixelColor(x, y, width, height, csr, csg, csb)
					img.Set(int(x)-int(cx1), int(y), colval)
					img.Set(int(x)-int(cx1), int(height)-int(y), colval)
				}
			} else {
				for y := cy1; y < cy2; y++ {
					colval := pixelColor(x, y, width, height, csr, csg, csb)
					img.Set(int(x)-int(cx1), int(y)-int(cy1), colval)
				}
			}
		}(x)
	}

	wg.Wait()

	return img, ""
}

func main() {
	img, _ := Mandelbrot(900, 900, 0, 900, 0, 900, 2, 3, 1)

	file, err := os.OpenFile("./images/mandelbrot.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()

	if err != nil {
		log.Fatalf("Error opening file: %s\n", err)
	}

	err = png.Encode(file, img)
	if err != nil {
		log.Fatalf("Error encoding image: %s\n", err)
	}
}

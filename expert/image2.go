package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type Circle struct {
	X, Y, Radius float64
}

func (c *Circle) Brightness(x, y float64) uint8 {
	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.Radius
	if d > 1 {
		return 0
	} else {
		return uint8((1 - math.Pow(d, 12)) * 255)
	}
}

func main() {
	var width, height int = 300, 300
	var hwidth, hheight float64 = float64(width / 2), float64(height / 2)
	var radius float64 = 42

	p := 2 * math.Pi / 3
	circleRed := &Circle{hwidth - radius*math.Sin(0), hheight - radius*math.Cos(0), 60}
	circleGreen := &Circle{hwidth - radius*math.Sin(p), hheight - radius*math.Cos(p), 60}
	circleBlue := &Circle{hwidth - radius*math.Sin(-p), hheight - radius*math.Cos(-p), 60}

	m := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			col := color.RGBA{
				circleRed.Brightness(float64(x), float64(y)),
				circleGreen.Brightness(float64(x), float64(y)),
				circleBlue.Brightness(float64(x), float64(y)),
				255,
			}
			m.Set(x, y, col)
		}
	}

	f, err := os.OpenFile("./images/rgb.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	png.Encode(f, m)
}

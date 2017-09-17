package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

type ImageSet interface {
	Set(x, y int, c color.Color)
}

func main() {
	file, err := os.Open("./selfcss.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", "./selfcss.png", err)
	}

	b := img.Bounds()

	imgSet := img.(ImageSet)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			oldPixel := img.At(x, y)
			r, g, b, a := oldPixel.RGBA()
			fmt.Println(r, g, b, a)
			pixel := color.RGBA{uint8(g), uint8(g), uint8(g), uint8(a)}
			imgSet.Set(x, y, pixel)
		}
	}

	fd, err := os.Create("./gray.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(fd, img)
	if err != nil {
		log.Fatal(err)
	}

	err = fd.Close()
	if err != nil {
		log.Fatal(err)
	}

	file, err = os.Open("./selfcss.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err = png.Decode(file)
	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", "./selfcss.png", err)
	}

	b = img.Bounds()

	imgSet = img.(ImageSet)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			oldPixel := img.At(x, y)
			r, g, b, a := oldPixel.RGBA()
			//fmt.Println(r, g, b, a)
			r = 65535 - r
			g = 65535 - g
			b = 65535 - b
			pixel := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			imgSet.Set(x, y, pixel)
		}
	}

	fd, err = os.Create("./inv.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(fd, img)
	if err != nil {
		log.Fatal(err)
	}

	err = fd.Close()
	if err != nil {
		log.Fatal(err)
	}
}

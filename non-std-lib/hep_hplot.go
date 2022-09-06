package main

import (
	"image/color"
	"log"

	"go-hep.org/x/hep/hbook"
	"go-hep.org/x/hep/hplot"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot/vg"
)

func main() {
	// define amount of datapoints
	const datapoints = 1024

	// Create a normal distribution
	distribution := distuv.Normal{
		Mu:    0,
		Sigma: 2,
		Src:   rand.New(rand.NewSource(0)),
	}

	// Draw some random values from the standard
	// normal distribution into a 1 dimensional histogram
	histogram := hbook.NewH1D(40, -4, +4)
	for i := 0; i < datapoints; i++ {
		v := distribution.Rand()
		histogram.Fill(v, 1)
	}

	// normalize histogram
	area := 0.0
	for _, bin := range histogram.Binning.Bins {
		area += bin.SumW() * bin.XWidth()
	}
	histogram.Scale(1 / area)

	// Make a new plot and set its title text and axis description
	plot := hplot.New()
	plot.Title.Text = "Histogram"
	plot.X.Label.Text = "X"
	plot.Y.Label.Text = "Y"

	// Create a new histogram and add previously defined values and texts
	h := hplot.NewH1D(histogram, hplot.WithHInfo(hplot.HInfoSummary), hplot.WithYErrBars(true), hplot.WithBand(true))
	h.YErrs.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	plot.Add(h)

	// Add a red line with the normal distribution function
	norm := hplot.NewFunction(distribution.Prob)
	norm.Color = color.RGBA{R: 255, A: 255}
	norm.Width = vg.Points(2)
	plot.Add(norm)

	// Save the plot as PNG
	if err := plot.Save(20*vg.Centimeter, -1, "hep-hplot-output.png"); err != nil {
		log.Fatalf("error saving plot: %v\n", err)
	}
}

package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func solveGasUpProblem(mpg float64, cityPump, cityDistance []float64) int {
	travelData := make([]float64, 1, len(cityPump))
	// prevMiles := 0.0
	// for ind, gallons := range cityPump {
	// 	travelData[ind] = gallons * mpg
	// }
	city := 0
	cityRemainingGallons := 50.0
	travelData[0] = cityRemainingGallons
	remainingGallons := 0.0
	for i := 1; i < len(cityDistance); i++ {
		// travelData[i] = travelData[i] - cityDistance[i-1] + travelData[i-1]
		remainingGallons += cityPump[i-1] - cityDistance[i-1]/mpg
		travelData = append(travelData, remainingGallons)
		if remainingGallons < cityRemainingGallons {
			city = i
			cityRemainingGallons = remainingGallons
		}
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Remaining Gallons"
	p.X.Label.Text = "Gallons"
	p.Y.Label.Text = "City"

	err = plotutil.AddLinePoints(p,
		"Remaining Gallons", dataPoints(travelData))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "gallonsLeftAtEachCity.png"); err != nil {
		panic(err)
	}

	travelData = append(travelData, 0, 0, 0)
	travelData = append([]float64{0, 0, 0}, travelData...)

	graph := asciigraph.Plot(travelData, asciigraph.Caption("Remaining gallons in each city"), asciigraph.Height(30), asciigraph.Width(100))
	fmt.Println(graph)
	fmt.Println()

	return city
}

func dataPoints(travelData []float64) plotter.XYs {
	pts := make(plotter.XYs, len(travelData))
	for i, v := range travelData {
		pts[i].X = float64(i)
		pts[i].Y = v
	}
	return pts
}

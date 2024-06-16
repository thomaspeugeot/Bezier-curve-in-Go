package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math"
)

var x = []float64{40, 3.5, 25, 25, -5, -5, 15, -0.5, 19.5, 7, 1.5}
var y = []float64{0, 36, 25, 1.5, 3, 33, 11, 35, 15.5, 0, 10.5}

func Newton(n int, k int) float64 {
	var rez float64 = 1
	for i := 1; i < k; i++ {
		rez = rez * float64(n-i+1.0) / float64(i)
	}
	return rez
}

func Bernstein(n int, i int, t float64) float64 {
	var rez = math.Pow(float64(t), float64(i)) * math.Pow(float64(1-t), float64(n-i))
	return Newton(n, i) * rez
}

var nbPlots = 200

func Points() plotter.XYs {
	var n int = 10
	pts := make(plotter.XYs, nbPlots+1)

	for z := 0; z <= nbPlots; z++ {
		var current_x, current_y, t float64 = 0.0, 0.0, float64(z) / float64(nbPlots)
		var temp float64
		var denominator float64 = 0.0

		for i := 0; i <= n; i++ {
			temp = x[i] * Bernstein(n, i, t)
			current_x += temp
			temp = y[i] * Bernstein(n, i, t)
			current_y += temp
			temp = Bernstein(n, i, t)
			denominator += temp
		}

		pts[z].X = current_x / denominator
		pts[z].Y = current_y / denominator
	}

	return pts
}
func main() {

	p := plot.New()

	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"Bézier Curve", Points())
	if err != nil {
		panic(err)
	}

	if err := p.Save(8*vg.Inch, 8*vg.Inch, "example.png"); err != nil {
		panic(err)
	}
}

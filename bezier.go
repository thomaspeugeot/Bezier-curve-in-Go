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

type Point struct {
	x float64
	y float64
}

var points__ = []Point{
	{x: x[0], y: y[0]},
	{x: x[1], y: y[1]},
	{x: x[2], y: y[2]},
	{x: x[3], y: y[3]},
	{x: x[4], y: y[4]},
	{x: x[5], y: y[5]},
	{x: x[6], y: y[6]},
	{x: x[7], y: y[7]},
	{x: x[8], y: y[8]},
	{x: x[9], y: y[9]},
	{x: x[10], y: y[10]},
}

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

func Points(points []Point, nbPlots int) []Point {
	n := len(points) - 1
	pts := make([]Point, nbPlots+1)

	for z := 0; z <= nbPlots; z++ {
		var current_x, current_y, t float64 = 0.0, 0.0, float64(z) / float64(nbPlots)
		var temp float64
		var denominator float64 = 0.0

		for i := 0; i <= n; i++ {
			temp = points[i].x * Bernstein(n, i, t)
			current_x += temp
			temp = points[i].y * Bernstein(n, i, t)
			current_y += temp
			temp = Bernstein(n, i, t)
			denominator += temp
		}

		pts[z].x = current_x / denominator
		pts[z].y = current_y / denominator
	}

	return pts
}

func PointsToPlotter(points []Point) plotter.XYs {

	pts := make(plotter.XYs, len(points))

	for i, point := range points {
		pts[i].X = point.x
		pts[i].Y = point.y
	}

	return pts
}

func main() {

	p := plot.New()

	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"Bézier Curve", PointsToPlotter(Points(points__, 30)))
	if err != nil {
		panic(err)
	}

	if err := p.Save(8*vg.Inch, 8*vg.Inch, "example.png"); err != nil {
		panic(err)
	}
}

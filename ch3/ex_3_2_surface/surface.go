// Surface computes SVG-representation from three-dimension graph of function.
// go run surface.go eggtray > example.svg
// go run surface.go saddle > example.svg
package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 640, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.04
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	switch {
	case len(os.Args) == 1:
		fmt.Fprintln(os.Stderr, "usage: surface eggtray|saddle")
		os.Exit(1)
	case os.Args[1] == "eggtray":
		printSVG(eggtray)
	case os.Args[1] == "saddle":
		printSVG(saddle)
	}
}

func printSVG(f func(x, y float64) float64) {
	fmt.Fprintf(os.Stdout,
		"<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>",
		width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)

			if isNaN(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			fmt.Fprintf(os.Stdout,
				"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintln(os.Stdout, "</svg>")
}

func corner(i, j int, f func(x, y float64) float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func eggtray(x, y float64) float64 {
	return math.Cos(x) + math.Cos(y)
}

func saddle(x, y float64) float64 {
	return y*y/100 - x*x/36
}

func isNaN(args ...float64) bool {
	for _, arg := range args {
		if math.IsNaN(arg) {
			return true
		}
	}
	return false
}

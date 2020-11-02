// Surface computes SVG-representation from three-dimension graph of function.
package main

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 640, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.3
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func printSVG(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w,
		"<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>\n",
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

			fmt.Fprintf(w,
				"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, f func(x, y float64) float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func isNaN(args ...float64) bool {
	for _, arg := range args {
		if math.IsNaN(arg) {
			return true
		}
	}
	return false
}

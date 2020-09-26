// Surface computes SVG-representation from three-dimension graph of function.
// go run surface.go > example.svg
package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type poligon struct {
	ax, ay, bx, by, cx, cy, dx, dy, z float64
}

func main() {
	printSVG(os.Stdout)
}

func printSVG(w io.Writer) {
	fmt.Fprintf(w,
		"<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: black; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>",
		width, height)

	plgs := poligons()
	min, max := minmaxz(&plgs)

	for _, p := range plgs {
		fmt.Fprintf(w,
			"<polygon "+
				"style='fill: #%06x' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			color(p.z, min, max), p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy)
	}

	fmt.Fprintln(os.Stdout, "</svg>")
}

func poligons() []poligon {
	var p []poligon

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, za := corner(i+1, j)
			bx, by, zb := corner(i, j)
			cx, cy, zc := corner(i, j+1)
			dx, dy, zd := corner(i+1, j+1)

			if isNaN(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			p = append(p, poligon{ax, ay, bx, by, cx, cy, dx, dy, avg(za, zb, zc, zd)})
		}
	}

	return p
}

func minmaxz(poligons *[]poligon) (float64, float64) {
	var min, max = math.MaxFloat64, math.SmallestNonzeroFloat64

	for _, poligon := range *poligons {
		if min > poligon.z {
			min = poligon.z
		}
		if max < poligon.z {
			max = poligon.z
		}
	}

	return min, max
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func isNaN(args ...float64) bool {
	for _, arg := range args {
		if math.IsNaN(arg) {
			return true
		}
	}

	return false
}

func avg(args ...float64) float64 {
	var a float64

	for _, arg := range args {
		a += arg
	}

	return a / float64(len(args))
}

func color(z, min, max float64) int {
	min = math.Abs(min)
	max = math.Abs(max)

	if z > 0 {
		return int(math.Exp(z)/math.Exp(max)*255) << 16
	} else {
		return int(math.Exp(-z) / math.Exp(min) * 255)
	}
}

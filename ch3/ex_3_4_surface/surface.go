// Surface computes SVG-representation from three-dimension graph of function.
// go run surface.go
// http://localhost:8000/?width=600&width=320&color=green
// http://localhost:8000/?width=600&width=320&color=%23ff00ff
package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
)

var angle, sin30, cos30 = math.Pi / 6, math.Sin(angle), math.Cos(angle)

type config struct {
	width, height int
	color         string
	cells         int
	xyrange       float64
	xyscale       float64
	zscale        float64
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		w.Header().Set("Content-Type", "image/svg+xml")
		printSVG(w, newConfig(r))
	})
	http.ListenAndServe(":8000", nil)
}

func newConfig(r *http.Request) config {
	c := config{width: 600, height: 320, color: "grey"}

	if s, ok := r.Form["width"]; ok {
		if v, err := strconv.Atoi(s[0]); err == nil {
			c.width = v
		}
	}

	if s, ok := r.Form["height"]; ok {
		if v, err := strconv.Atoi(s[0]); err == nil {
			c.height = v
		}
	}

	if s, ok := r.Form["color"]; ok {
		c.color = s[0]
	}

	c.cells = 100
	c.xyrange = 30.0
	c.xyscale = float64(c.width) / 2 / c.xyrange
	c.zscale = float64(c.height) * 0.4

	return c
}

func printSVG(w io.Writer, c config) {
	fmt.Fprintf(w,
		"<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: %s; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>\n",
		c.color, c.width, c.height)

	for i := 0; i < c.cells; i++ {
		for j := 0; j < c.cells; j++ {
			ax, ay := corner(i+1, j, c)
			bx, by := corner(i, j, c)
			cx, cy := corner(i, j+1, c)
			dx, dy := corner(i+1, j+1, c)

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

func corner(i, j int, c config) (float64, float64) {
	x := c.xyrange * (float64(i)/float64(c.cells) - 0.5)
	y := c.xyrange * (float64(j)/float64(c.cells) - 0.5)
	z := f(x, y)
	sx := float64(c.width)/2 + (x-y)*cos30*c.xyscale
	sy := float64(c.height)/2 + (x+y)*sin30*c.xyscale - z*c.zscale
	return sx, sy
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

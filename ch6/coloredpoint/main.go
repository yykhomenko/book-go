package main

import (
	"fmt"
	"image/color"

	"github.com/yykhomenko/book-gopl/ch6/geometry"
)

type ColoredPoint struct {
	geometry.Point
	color color.Color
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)
	fmt.Println()

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	p := ColoredPoint{geometry.Point{1, 1}, red}
	q := ColoredPoint{geometry.Point{5, 4}, blue}

	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

	q.Point = p.Point
	fmt.Println(p.Point, q.Point)
}

package display

import (
	"image/color"
	"testing"

	"github.com/yykhomenko/book-gopl/ch6/geometry"
)

func TestDisplayMapStruct(t *testing.T) {
	v := make(map[geometry.Point]color.Color)
	v[geometry.Point{1, 2}] = color.RGBA{255, 128, 255, 255}
	Display("map", v)
}

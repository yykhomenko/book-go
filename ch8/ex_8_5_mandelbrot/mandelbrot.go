// Mandelbrot prints Mandelbrot's fractal to standard output.
// go run mandelbrot.go > example.png
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, heigth          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, heigth))

	for py := 0; py < heigth; py++ {
		y := float64(py)/heigth*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		r := real(v)
		i := imag(v)
		if r*r+i*i > 4 {
			return color.Gray{255 - contrast*n}
		}
	}

	return color.Black
}

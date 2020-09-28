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
		width, heigth          = 8000, 8000
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
	var v complex128

	for n := 0; n < iterations; n++ {
		v = v*v + z
		r := real(v)
		i := imag(v)
		if r*r+i*i > 4 {
			return toColor(n, iterations)
		}
	}

	return color.Black
}

func toColor(n, nmax int) color.Color {
	const contrast = 20

	deep := float64(n) / float64(nmax)
	y := 255 * deep * contrast
	if y > 255 {
		y = 255
	}

	switch {
	case deep > 0.2:
		return color.RGBA{
			R: uint8(y),
			A: 255,
		}
	case deep > 0.1:
		return color.White
	default:
		return color.RGBA{
			B: uint8(y),
			A: 255,
		}
	}
}

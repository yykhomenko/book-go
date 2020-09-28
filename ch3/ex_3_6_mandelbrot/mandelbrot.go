// Mandelbrot prints Mandelbrot's fractal to standard output. Supersampled.
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
		width, heigth          = 2048, 2048
	)

	img := image.NewRGBA(image.Rect(0, 0, width, heigth))

	for py := 0; py < heigth; py++ {
		y := float64(py)/heigth*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			const resY = 1.0 / heigth
			const resX = 1.0 / width

			c1 := mandelbrot(complex(x+resX, y+resY))
			c2 := mandelbrot(complex(x+resX, y-resY))
			c3 := mandelbrot(complex(x-resX, y-resY))
			c4 := mandelbrot(complex(x-resX, y+resY))

			c := color.Gray{uint8((c1 + c2 + c3 + c4) / 4)}

			img.Set(px, py, c)
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) int {
	const contrast = 15
	const iterations = 200
	var v complex128

	for n := 0; n < iterations; n++ {
		v = v*v + z
		r := real(v)
		i := imag(v)
		if r*r+i*i > 4 {
			return 255 - contrast*n
		}
	}

	return 0
}

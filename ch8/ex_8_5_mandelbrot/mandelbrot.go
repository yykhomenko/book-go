// Mandelbrot prints Mandelbrot's fractal to standard output.
// go run mandelbrot.go > example.png
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"sync"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, heigth          = 10000, 10000
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, heigth))
	rows := make(chan int)
	go generate(rows, heigth)

	wg := &sync.WaitGroup{}
	for i := 0; i < runtime.GOMAXPROCS(-1); i++ {
		wg.Add(1)
		go setRows(img, wg, rows)
	}
	wg.Wait()

	png.Encode(os.Stdout, img)
}

func generate(rows chan<- int, n int) {
	defer close(rows)
	for row := 0; row < n; row++ {
		rows <- row
	}
}

func setRows(img *image.RGBA, wg *sync.WaitGroup, rows <-chan int) {
	defer wg.Done()
	for row := range rows {
		y := float64(row)/heigth*(ymax-ymin) + ymin
		for col := 0; col < width; col++ {
			x := float64(col)/width*(xmax-xmin) + xmin
			img.Set(col, row, mandelbrot(complex(x, y)))
		}
	}
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

// Lissajous generate animation GIF file from random Lissajous figure.
// go build lissajous.go && ./lissajous > example.gif
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{R: 255, A: 255},
	color.RGBA{G: 255, A: 255},
	color.RGBA{B: 255, A: 255},
	color.RGBA{R: 255, G: 255, A: 255},
	color.RGBA{R: 255, B: 255, A: 255},
	color.RGBA{G: 255, B: 255, A: 255},
	color.White,
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Number of total oscillation x
		res     = 0.001 // Angle's resolution
		size    = 100   // Size of canvas
		nFrames = 64    // Number of GIF frames
		delay   = 8     // Delay between frames (1 is 10ms)
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0

	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				index(t),
			)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}

func index(t float64) uint8 {
	switch {
	case t < math.Pi/8:
		return 1
	case t < math.Pi/4:
		return 2
	case t < math.Pi/2:
		return 3
	case t < math.Pi:
		return 4
	case t < math.Pi*2:
		return 5
	case t < math.Pi*4:
		return 6
	default:
		return uint8(rand.Intn(len(palette)))
	}
}

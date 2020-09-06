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

var palette = []color.Color{color.Black, color.RGBA{G: 255, A: 255}}

const (
	blackIndex = 0
	greenIndex = 1
)

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
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}

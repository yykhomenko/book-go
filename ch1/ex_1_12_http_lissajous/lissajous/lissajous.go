// Lissajous generate animation GIF file from random Lissajous figure.
// go run Lissajous.go > example.gif
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
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

type Config struct {
	Cycles  int     // Number of total oscillation x
	Res     float64 // Angle's resolution
	Size    int     // Size of canvas
	NFrames int     // Number of GIF frames
	Delay   int     // Delay between frames (1 is 10ms)
}

func setDefaults(conf Config) Config {
	if conf.Cycles == 0 {
		conf.Cycles = 5
	}

	if conf.Res == 0.0 {
		conf.Res = 0.001
	}

	if conf.Size == 0 {
		conf.Size = 100
	}

	if conf.NFrames == 0 {
		conf.NFrames = 64
	}

	if conf.Delay == 0 {
		conf.Delay = 8
	}

	return conf
}

func Lissajous(out io.Writer, conf Config) {
	conf = setDefaults(conf)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: conf.NFrames}
	phase := 0.0

	for i := 0; i < conf.NFrames; i++ {
		rect := image.Rect(0, 0, 2*conf.Size+1, 2*conf.Size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < 2*math.Pi*float64(conf.Cycles); t += conf.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				conf.Size+int(float64(conf.Size)*x+0.5),
				conf.Size+int(float64(conf.Size)*y+0.5),
				index(t),
			)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, conf.Delay)
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

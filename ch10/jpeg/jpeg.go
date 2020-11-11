package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
)

func main() {
	if err := toJpeg(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
	}
}

func toJpeg(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

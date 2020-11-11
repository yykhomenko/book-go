// Imgconv converts picture.
// ./imgconv -f=png|jpg|gif <INPUT >OUTPUT
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	format := flag.String("f", "jpg", "output format [png, gif, jpg]")
	flag.Parse()

	info, _ := os.Stdout.Stat()
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Fprintln(os.Stderr, "Refusing to write to character device. Redirect output to a pipe or regular file.")
		os.Exit(1)
	}

	if err := toFormat(os.Stdin, os.Stdout, *format); err != nil {
		fmt.Fprintf(os.Stderr, "imgconv: %v\n", err)
		os.Exit(1)
	}
}

func toFormat(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "convert %s to %s\n", kind, format)
	switch format {
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, nil)
	case "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	default:
		return fmt.Errorf("unknown output format %s", format)
	}
}

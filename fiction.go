package main

import (
	"image"
	"sync"

	"image/png"

	"image/color"

	"fmt"

	"math/cmplx"

	"os"

	"flag"

	"strconv"

	"time"

	"strings"

	"encoding/hex"

	"errors"
)

func main() {

}

func draw() {
	w, h := parseArgs()
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	out := time.Now().Format("20060102T150405") + ".png"
	f, err := os.Create(out)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer f.Close()
	drawMandelbrot(img)
	png.Encode(f, img)
}

func parseArgs() (int, int) {

	var w, h int

	var err error

	flag.Parse()

	args := flag.Args()

	len := len(args)

	switch {

	case len > 0:

		w, err = strconv.Atoi(args[0])

		h = w

	case len > 1:

		h, err = strconv.Atoi(args[1])

	default:

		w, h = 800, 600

	}

	if w <= 0 || h <= 0 || err != nil {

		fmt.Println("error arguments, ", w, h)

		os.Exit(1)

	}

	return w, h

}

func drawMandelbrot(img *image.RGBA) {

	const (
		MIN_Y = -1.6

		MAX_Y = 1.6

		MIN_X = -2.2

		MAX_X = 1.2
	)

	height := img.Bounds().Dy()

	width := img.Bounds().Dx()

	q := make(chan func(), 4)
	w := sync.WaitGroup{}

	startTaskPool(q, &w)

	for y := 0; y < height; y++ {

		yy := y
		task := func() {

			for x := 0; x < width; x++ {

				zx := mapValue(float64(x), MIN_X, MAX_X, 0.0, float64(width))

				zy := mapValue(float64(y), MIN_Y, MAX_Y, 0.0, float64(height))

				drawMandelbrotPoint(img, image.Point{x, yy}, complex(zx, zy))

			}
		}
		q <- task

	}
	close(q)
	w.Wait()
}

func startTaskPool(q chan func(), w *sync.WaitGroup) {
	for i := 0; i < 4; i++ {
		go startTaskProcessor(q, w)
	}
}

func startTaskProcessor(q chan func(), w *sync.WaitGroup) {
	w.Add(1)
	for task := range q {
		task()
	}
	w.Done()
}

func drawMandelbrotPoint(img *image.RGBA, p image.Point, c complex128) {

	var z = c

	bg, _ := fromColorString("#00CED1")

	for i := uint8(0); i < 255; i++ {

		z = z*z + c

		if cmplx.Abs(z) > 2 {

			cl := color.RGBA{i % 8 * 32, i % 12 * 16, i % 16 * 8, 255}

			img.Set(p.X, p.Y, cl)

			return

		}

	}

	img.Set(p.X, p.Y, bg)

}

func mapValue(x float64, minsrc float64, maxsrc float64, mintar float64, maxtar float64) float64 {

	return x*(maxsrc-minsrc)/(maxtar-mintar) + minsrc

}

func fromColorString(s string) (color.RGBA, error) {

	black := color.RGBA{0, 0, 0, 255}

	if strings.HasPrefix(s, "#") && len(s) == 7 {

		r, err := hex.DecodeString(s[1:])

		if err != nil {

			return black, err

		}

		return color.RGBA{r[0], r[1], r[2], 255}, nil

	} else {

		return black, errors.New("arguments error")

	}

}

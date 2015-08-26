// This program reads a squid image, finds the tentacles, and outputs the sparse transition matrix.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"

	png "image/png"
)

var in = flag.String("i", "squids_1.png", "file name of input image")
var out = flag.String("o", "output.png", "file name of output image")

func main() {
	flag.Parse()

	reader, err := os.Open(*in)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	i, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	b := i.Bounds()
	m := image.NewGray(b) // monochrome
	draw.Draw(m, b, i, b.Min, draw.Src)

	o := image.NewGray(image.Rect(0, 0, 256, 256))
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}
	black := color.RGBA{0x00, 0x00, 0x00, 0xff}
	draw.Draw(o, o.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	delta := []int{-16, -1, 1, 16}
	ud := []int{-6, 0, 0, 6}
	lr := []int{0, -6, 6, 0}
	for j := 0; j < 16; j++ {
		for i := 0; i < 16; i++ {
			cx := int(19 + 20.5333*float64(i))
			cy := int(19 + 20.5333*float64(j))
			s := []uint32{0, 0, 0, 0}
			best := uint32(50000)
			bestK := 0
			for k := 0; k < 4; k++ {
				for sx := -1; sx < 2; sx++ {
					for sy := -1; sy < 2; sy++ {
						r, g, b, _ := m.At(cx+lr[k]+sx, cy+ud[k]+sy).RGBA()
						s[k] += (r + g + b) / 3
					}
				}
				if k == 0 || s[k] < best {
					best = s[k]
					bestK = k
				}
			}

			src := 16*j + i
			dst := (src + delta[bestK] + 256) % 256
			fmt.Printf("%v\n", dst)
			o.Set(src, dst, black)
		}
	}

	writer, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(writer, o)
	writer.Close()
}

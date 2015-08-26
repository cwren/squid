// This program multiplies sparse permutation matricies.
// With one input it computes powers until it finds the identity.
// With two inputs it mutliplies the matricies and prints the result.
// In either case it generates PNGs of the resulting matrices.
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
	"strconv"

	png "image/png"
)

var aFile = flag.String("a", "A.txt", "left input matrix")
var bFile = flag.String("b", "", "right input matrix")
var output = flag.String("o", "step%06d.png", "output file pattern")

func main() {
	flag.Parse()

	a, err := readMatrix(aFile)
	if err != nil {
		log.Fatal(err)
	}

	if *bFile == "" {
		findIdentityPower(a)
	} else {
		b, bErr := readMatrix(bFile)
		if bErr != nil {
			log.Fatal(bErr)
		}

		multiply(a, b)
	}
}

func findIdentityPower(a []int) {
	b := make([]int, len(a))
	c := make([]int, len(a))

	copy(b, a)
	step := 0
	isIdent := false
	for !isIdent {
		fmt.Printf("step %v\n", step)
		renderOut(output, step, a)
		isIdent = true
		for i := 0; i < len(a); i++ {
			c[i] = a[b[i]]
			isIdent = isIdent && c[i] == i
		}
		step++
		copy(a, c)
	}
	fmt.Printf("step %v\n", step)
	renderOut(output, step, a)
}

func multiply(a, b []int) {
	if len(a) != len(b) {
		log.Fatal(errors.New("inputs are different lengths"))
	}

	for i := 0; i < len(a); i++ {
		fmt.Printf("%v\n", a[b[i]])
	}
	renderOut(output, 0, a)
}

func readMatrix(filename *string) ([]int, error) {
	f, osErr := os.Open(*filename)
	if osErr != nil {
		return nil, osErr
	}
	defer f.Close()

	m := make([]int, 0, 256)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		m = append(m, x)
	}
	return m, scanner.Err()
}

func renderOut(pattern *string, step int, p []int) {
	filename := fmt.Sprintf(*pattern, step)

	writer, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	o := image.NewGray(image.Rect(0, 0, 256, 256))
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}
	black := color.RGBA{0x00, 0x00, 0x00, 0xff}
	draw.Draw(o, o.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	for i := 0; i < len(p); i++ {
		o.Set(i, p[i], black)
	}

	png.Encode(writer, o)
	writer.Close()
}

package main

import (
	"fmt"
	"image/png"
	"math"
	"os"

	"github.com/fogleman/gg"
)

func drawSquircle(width, height int, fileName string) error {
	dc := gg.NewContext(width, height)

	dc.Clear()

	n := 3.3
	rx := float64(width) / 2
	ry := float64(height) / 2
	centerX := float64(width) / 2
	centerY := float64(height) / 2
	step := 0.01

	dc.NewSubPath()
	for theta := 0.0; theta <= 2*math.Pi; theta += step {
		x := math.Pow(math.Abs(math.Cos(theta)), 2/n)*rx*math.Copysign(1, math.Cos(theta)) + centerX
		y := math.Pow(math.Abs(math.Sin(theta)), 2/n)*ry*math.Copysign(1, math.Sin(theta)) + centerY
		dc.LineTo(x, y)
	}
	dc.ClosePath()

	dc.SetRGB(0.2, 0.6, 0.86) // Corresponding to #3498db
	dc.Fill()

	// Save to file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := png.Encode(file, dc.Image()); err != nil {
		return err
	}

	fmt.Println("Squircle image saved as", fileName)
	return nil
}

func main() {
	width := 400
	height := 400
	fileName := "./squircle.png"

	if err := drawSquircle(width, height, fileName); err != nil {
		fmt.Println("Error saving squircle image:", err)
		return
	}
}

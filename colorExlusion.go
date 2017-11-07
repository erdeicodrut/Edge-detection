package main

import (
	"image/jpeg"
	"os"
	"image/png"
	"image"
	"image/color"
	"log"
)

func main() {
	photo, _ := os.Open("input.jpg")

	ext := photo.Name()[len(photo.Name())-3:]

	var pic image.Image

	if ext == "jpg" || ext == "peg" {
		pic, _ = jpeg.Decode(photo)
	}

	if ext == "png" {
		pic, _ = png.Decode(photo)
	}

	limit := pic.Bounds().Size()

	img := image.NewNRGBA(image.Rect(0, 0, limit.X, limit.Y))

	var maxIntenisty uint32

	for y := 0; y < limit.Y; y++ {
		for x := 0; x < limit.X; x++ {

			r, g, b, _ := pic.At(x, y).RGBA()

			intensity := (r + g + b) / 3

			if intensity > maxIntenisty {
				maxIntenisty = intensity
			}
		}
	}

	for y := 0; y < limit.Y; y++ {
		for x := 0; x < limit.X; x++ {

			r, g, b, _ := pic.At(x, y).RGBA()

			intensity := (r + g + b) / 3

			if intensity > maxIntenisty / 2 {
				setColor32(img, x, y, r, g, b)
			} else {
				setColor32(img, x, y, intensity, intensity, intensity)
			}

		}
	}

	f, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

func setColor32(img *image.NRGBA, x, y int, r, g, b uint32) {
	img.Set(x, y, color.NRGBA{
		R: uint8(r / 256),
		G: uint8(g / 256),
		B: uint8(b / 256),
		A: 255,
	})
}

func setColor8(img *image.NRGBA, x, y int, r, g, b uint8) {
	img.Set(x, y, color.NRGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	})
}

package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
)

const QE = 70

func main() {
	photo, _ := os.Open("input1.jpg")

	pic, _ := jpeg.Decode(photo)

	// Create a colored image of the given width and height.
	limit := pic.Bounds().Size()

	img := image.NewNRGBA(image.Rect(0, 0, limit.X, limit.Y))

	for y := 0; y < limit.Y; y++ {
		for x := 0; x < limit.X; x++ {

			pixelO := pic.At(x, y)
			x2, y2, z2, _ := pixelO.RGBA()

			img.Set(x, y, color.NRGBA{
				R: uint8((0) & 255),
				G: uint8((0) << 1 & 255),
				B: uint8((0) << 2 & 255),
				A: 255,
			})

			var sum float64
			for _, offsetX := range []int{-1, 0, 1} {
				for _, offsetY := range []int{-1, 0, 1} {
					if offsetX == offsetY {
						break
					}
					pixelN := pic.At(x+offsetX, y+offsetY)
					x1, y1, z1, _ := pixelN.RGBA()

					xSqr := (x1 - x2) * (x1 - x2)
					ySqr := (y1 - y2) * (y1 - y2)
					zSqr := (z1 - z2) * (z1 - z2)
					mySqr := float64(xSqr + ySqr + zSqr)
					dist := math.Sqrt(mySqr)
					sum += dist

				}
			}

			avg := sum / 8

			if avg < 65536/QE {
				img.Set(x, y, color.NRGBA{
					R: uint8((x2) & 255),
					G: uint8((y2) << 1 & 255),
					B: uint8((z2) << 2 & 255),
					A: 255,
				})
			}

		}
	}

	f, err := os.Create("output.png")
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

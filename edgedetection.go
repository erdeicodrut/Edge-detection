package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
)

func main() {
	photo, _ := os.Open("toDestroy.jpg")

	pic, _ := jpeg.Decode(photo)

	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(10)
	limit := pic.Bounds().Size()

	for QE := 10; QE < 200; QE += 20 {

		go func(QWE int) {
			img := image.NewNRGBA(image.Rect(0, 0, limit.X, limit.Y))
			// Create a colored image of the given width and height.

			for y := 0; y < limit.Y; y++ {
				for x := 0; x < limit.X; x++ {

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

							m.Lock()
							pixelN := pic.At(x+offsetX, y+offsetY)
							pixelO := pic.At(x, y)
							m.Unlock()

							x1, y1, z1, _ := pixelN.RGBA()

							x2, y2, z2, _ := pixelO.RGBA()

							xSqr := (x1 - x2) * (x1 - x2)
							ySqr := (y1 - y2) * (y1 - y2)
							zSqr := (z1 - z2) * (z1 - z2)
							mySqr := float64(xSqr + ySqr + zSqr)
							dist := math.Sqrt(mySqr)
							sum += dist

						}
					}

					avg := sum / 8

					if avg < 65536/float64(QWE) {
						img.Set(x, y, color.NRGBA{
							R: uint8((255) & 255),
							G: uint8((255) << 1 & 255),
							B: uint8((255) << 2 & 255),
							A: 255,
						})
					}

				}
			}
			name := photo.Name() + "output" + strconv.Itoa(QWE) + ".png"

			f, err := os.Create(name)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(name)

			if err := png.Encode(f, img); err != nil {
				f.Close()
				log.Fatal(err)
			}

			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}(QE)
	}
	wg.Wait()
}

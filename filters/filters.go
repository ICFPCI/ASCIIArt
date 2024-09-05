package filters

import (
	"fmt"
	"image"
	"image/color"
	"image_filters/utils"
)

func getTileIndex(gray uint8) int {
	normalized := float64(gray) / 255.0
	return int(normalized * 9)
}

func ASCIIArt(img image.Gray) (image.Image, error) {
	texture, _, err := utils.LoadImage("images/texture.png")

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	tileSize := 8

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := img.At(x, y)
			gray, _, _, _ := color.GrayModel.Convert(pixelColor).RGBA()

			grayValue := uint8(gray >> 8)
			tileIndex := getTileIndex(grayValue)

			for dy := 0; dy < tileSize; dy++ {
				for dx := 0; dx < tileSize; dx++ {
					tx := x + dx
					ty := y + dy
					// if tx < bounds.Max.X && ty < bounds.Max.Y {
					textureX := tileIndex*8 + dx
					fmt.Printf("x: %d, y: %d\n", tileIndex, dy)

					newImg.Set(tx, ty, texture.At(textureX+dx, dy))
					// newImg.Set(tx, ty, color.RGBA{R: 255, G: 0, B: 0, A: 255})
					// }
				}
			}
		}
	}
	return newImg, nil
}

func UpsideDown(pixels [][]color.Color) {
	for x := 0; x < len(pixels); x++ {
		line := pixels[x]
		for y := 0; y < len(line)/2; y++ {
			k := len(line) - y - 1
			line[y], line[k] = line[k], line[y]
		}
	}
}

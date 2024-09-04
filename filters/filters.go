package filters

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image_filters/utils"
)

func getTileIndex(gray uint8) int {
	normalized := float64(gray) / 255.0
	return int(normalized * 9)
}

func ASCIIArt(img image.Image) (image.Image, error) {
	texture, _, err := utils.LoadImage("./images/texture.jpg")

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error al cargar la imagen")
	}

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	tileSize := 8
	numTiles := 10

	tiles := make([]image.Image, numTiles)
	for i := 0; i < numTiles; i++ {
		// current := (i) * tileSize
		tileBounds := image.Rect(0, 0, tileSize, tileSize)
		tiles[i] = image.NewRGBA(tileBounds)
		draw.Draw(tiles[i].(*image.RGBA), tileBounds, texture, tileBounds.Min, draw.Src)
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := img.At(x, y)
			gray, _, _, _ := color.GrayModel.Convert(pixelColor).RGBA()

			// Convertir el valor de gris a uint8
			grayValue := uint8(gray >> 8)
			tileIndex := getTileIndex(grayValue)
			tile := tiles[tileIndex]

			// Poner el tile en la imagen nueva
			for dy := 0; dy < tileSize; dy++ {
				for dx := 0; dx < tileSize; dx++ {
					tx := x + dx
					ty := y + dy
					if tx < bounds.Max.X && ty < bounds.Max.Y {
						newImg.Set(tx, ty, tile.At(dx, dy))
					}
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

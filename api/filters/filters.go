package filters

import (
	"image"
	"image/color"
	"image_filters/api/utils"
)

func getTileIndex(gray uint8) int {
	normalized := float64(gray) / 255.0
	return int(normalized * 9)
}

func ASCIIArt(img image.Image) (image.Image, error) {
	texture, _, err := utils.LoadImage("api/textures/chars.png")

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	rect := image.Rect(0, 0, bounds.Dx()*8, bounds.Dy()*8)
	newImg := image.NewRGBA(rect)

	tileSize := 8

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := img.At(x, y)
			gray, _, _, _ := color.GrayModel.Convert(pixelColor).RGBA()

			grayValue := uint8(gray >> 8)
			tileIndex := getTileIndex(grayValue)

			for dy := 0; dy < tileSize; dy++ {
				for dx := 0; dx < tileSize; dx++ {
					tx := x*8 + dx
					ty := y*8 + dy
					textureX := tileIndex*8 + dx
					newImg.Set(tx, ty, texture.At(textureX, dy))
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

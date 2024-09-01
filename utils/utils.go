package utils

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
)

func ImageToTensors(img image.Image) [][]color.Color {
	size := img.Bounds().Size()
	var pixels [][]color.Color

	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		pixels = append(pixels, y)
	}
	return pixels
}

func TensorToImage(pixels [][]color.Color) image.Image {
	rect := image.Rect(0, 0, len(pixels), len(pixels[0]))
	nImg := image.NewRGBA(rect)

	for x := 0; x < len(pixels); x++ {
		for y := 0; y < len(pixels[0]); y++ {
			q := pixels[x]
			if q == nil {
				continue
			}
			p := pixels[x][y]
			if p == nil {
				continue
			}
			original, ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok {
				nImg.Set(x, y, original)
			}
		}
	}

	return nImg

}

func LoadImage(filePath string) (image.Image, string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, "", err
	}

	defer file.Close()

	img, imgType, err := image.Decode(file)

	if err != nil {
		return nil, "", err
	}

	return img, imgType, nil
}

func SaveImage(pixels [][]color.Color) error {
	file, err := os.Create("./result/image.png")

	if err != nil {
		return err
	}

	defer file.Close()

	img := TensorToImage(pixels)

	if err := png.Encode(file, img); err != nil {
		return err
	}

	return nil
}

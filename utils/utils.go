package utils

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"

	"golang.org/x/image/draw"
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

func SaveImage(img image.Image) error {
	file, err := os.Create("./result/image.png")

	if err != nil {
		return err
	}

	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return err
	}

	return nil
}

func SaveImageFromPixels(pixels [][]color.Color) error {
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

func DownSample(img image.Image, scaleFactor int) image.Image {
	bounds := img.Bounds().Size()
	width := int(bounds.X / scaleFactor)
	height := int(bounds.Y / scaleFactor)
	oldBounds := image.Rect(0, 0, bounds.X, bounds.Y)
	newBounds := image.Rect(0, 0, width, height)

	newImg := image.NewRGBA(newBounds)

	draw.NearestNeighbor.Scale(newImg, newBounds, img, oldBounds, draw.Over, nil)

	return newImg
}

func Resize(img image.Image, originalSize image.Rectangle) image.Image {
	newImg := image.NewRGBA(originalSize)

	// Redibujar la imagen reducida al tamaño original
	draw.NearestNeighbor.Scale(newImg, originalSize, img, img.Bounds(), draw.Over, nil)

	return newImg
}

func ConvertToGray(img image.Image) image.Image {
	bounds := img.Bounds()

	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := (r + g + b) / 3
			grayImg.Set(x, y, color.Gray{uint8(gray >> 8)})
		}
	}

	return grayImg
}

func QuantizeGray(img *image.Gray, numLevels int) *image.Gray {
	bounds := img.Bounds()

	quantizedImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayColor := img.At(x, y).(color.Gray)
			grayValue := grayColor.Y
			level := grayValue / uint8(256/numLevels)
			quantizedImg.Set(x, y, color.Gray{level * uint8(256/numLevels)})
		}
	}

	return quantizedImg
}

func SaveFile(text string) {
	file, err := os.Create("./result/image.txt")

	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Println("Texto guardado con éxito en ./result/image.txt")
}

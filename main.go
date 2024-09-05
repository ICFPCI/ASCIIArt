package main

import (
	"fmt"
	"image_filters/filters"
	"image_filters/utils"
)

func main() {
	img, imageType, err := utils.LoadImage("images/test.jpeg")

	if err != nil {
		fmt.Print("Error: ", err)
		return
	}

	fmt.Println(imageType)

	downscaledImg := utils.DownSample(img, 8)

	grayScaleImg := utils.ConvertToGray(downscaledImg)
	// grayScaleImg = utils.QuantizeGray(grayScaleImg, 10)

	asciiImg, err := filters.ASCIIArt(grayScaleImg)

	// asciiImg = utils.Resize(asciiImg, img.Bounds())

	if err != nil {
		fmt.Println(err)
		return
	}

	print(asciiImg)

	utils.SaveImage(asciiImg)

	fmt.Println("Save file to ./result/")

}

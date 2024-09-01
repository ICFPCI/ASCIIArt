package main

import (
	"fmt"
	"image_filters/filters"
	"image_filters/utils"
)

func main() {

	img, imageType, err := utils.LoadImage("/home/fabrizzio/Desktop/code/image_filters/test/test.png")

	if err != nil {
		fmt.Print("Error: ", err)
		return
	}

	fmt.Println(imageType)

	pixels := utils.ImageToTensors(img)

	// Applying filters
	filters.UpsideDown(pixels)

	if err := utils.SaveImage(pixels); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Save file to ./result/")

}

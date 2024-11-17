package handlers

import (
	"fmt"
	"image"
	"image/png"
	"image_filters/api/filters"
	"image_filters/api/utils"
	"net/http"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	img, _, err := image.Decode(r.Body)

	if err != nil {
		http.Error(w, "Error al leer el archivo", http.StatusBadRequest)
		return
	}

	downscaledImg := utils.DownSample(img, 8)

	grayScaleImg := utils.ConvertToGray(downscaledImg)

	asciiImg, err := filters.ASCIIArt(grayScaleImg)

	asciiImg = utils.Resize(asciiImg, img.Bounds())

	if err != nil {
		fmt.Println(err)
		return
	}

	if err := png.Encode(w, asciiImg); err != nil {
		http.Error(w, "Error al codificar el archivo", http.StatusBadRequest)
		return
	}
}

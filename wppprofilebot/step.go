package wppprofilebot

import (
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"os"
)

func DoNextStep() image.Image {

	imgRo := GetImage()
	img := image.NewRGBA(imgRo.Bounds())
	draw.Draw(img, imgRo.Bounds(), imgRo, image.ZP, draw.Src)

	data, err := ioutil.ReadFile("pos")
	if err != nil {
		log.Println("Generating new movement file")
		data = GenerateMovement(img, 16)
	} else {
		data[0], data[1] = MakeMovement(img, data[0], data[1], data[2])
		if len(data) == 3 {
			data = nil
		} else {
			data = append(data[:2], data[3:]...)
		}
	}
	SaveCurrent(img)
	if data == nil {
		log.Println("done, removing state files")
		_ = os.Remove("pos")
		_ = RemoveCurrent()
	} else {
		_ = ioutil.WriteFile("pos", data, os.ModePerm)
	}
	return img
}

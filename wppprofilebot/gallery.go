package wppprofilebot

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

const curFile = "current.jpg"

func getNewImage() (file *os.File) {
	available, _ := ioutil.ReadDir("gallery")
	if len(available) == 0 {
		fmt.Println("Not image provided")
		os.Exit(1)
	}
	s := rand.Intn(len(available))
	var err error
	fileName := "gallery/" + available[s].Name()
	file, err = os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not open %s", fileName)
	}
	return
}

func GetImage() image.Image {

	imgBuffer, err := os.Open(curFile)
	if err != nil {
		log.Println("Getting new image")
		imgBuffer = getNewImage()
	}
	imgRo, _, _ := image.Decode(imgBuffer)
	_ = imgBuffer.Close()

	return imgRo
}

func SaveCurrent(img image.Image) {
	buf, _ := os.OpenFile(curFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	_ = jpeg.Encode(buf, img, nil)
	_ = buf.Close()
}

func RemoveCurrent() error {
	return os.Remove(curFile)
}

package wppprofilebot

import (
	"image"
	"image/color"
	"image/draw"
)

/**
 * top = 0
 * right = 1
 * left = 2
 * bottom = 3
 */

func SetupBlackSquare(img draw.Image) {
	b := img.Bounds()
	width := b.Dx()
	height := b.Dy()
	wS := width / 4
	wH := height / 4
	rect := image.Rect(0, 0, wS, wH).Add(image.Point{X: wS * 3, Y: wH * 3})
	buf := image.NewUniform(color.Black)
	draw.Draw(img, rect, buf, image.ZP, draw.Src)
}

func MakeMovement(img draw.Image, posX, posY, dir byte) (byte, byte) {
	b := img.Bounds()
	width := b.Dx()
	height := b.Dy()
	wS := width / 4
	wH := height / 4
	posFX := posX
	posFY := posY
	switch dir {
	case 0:
		posFY--
		break
	case 1:
		posFX++
		break
	case 2:
		posFX--
		break
	case 3:
		posFY++
		break
	}
	rect := image.Rect(0, 0, wS, wH)
	p1 := image.Point{X: wS * int(posX), Y: wH * int(posY)}
	p2 := image.Point{X: wS * int(posFX), Y: wH * int(posFY)}
	buf := image.NewRGBA(rect)
	draw.Draw(buf, rect, img, p1, draw.Src)
	draw.Draw(img, rect.Add(p1), img, p2, draw.Src)
	draw.Draw(img, rect.Add(p2), buf, image.ZP, draw.Src)
	return posFX, posFY
}

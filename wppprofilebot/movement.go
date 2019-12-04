package wppprofilebot

import (
	"image/draw"
	"math/rand"
)

/**
 * top = 0
 * right = 1
 * left = 2
 * bottom = 3
 */

func GenerateMovement(image draw.Image, stepCount int) []byte {
	data := make([]byte, stepCount+2)
	steps := data[2:]
	posX := byte(3)
	posY := byte(3)
	SetupBlackSquare(image)
	last := -1
	for i := 0; i < stepCount; i++ {
		n := rand.Intn(4)
		switch n {
		case 3 - last:
			{
				i--
				continue
			}
		case 0:
			if posY == 0 {
				i--
				continue
			}
			break
		case 1:
			if posX == 3 {
				i--
				continue
			}
			break
		case 2:
			if posX == 0 {
				i--
				continue
			}
		case 3:
			if posY == 3 {
				i--
				continue
			}
		}
		last = n
		posX, posY = MakeMovement(image, posX, posY, byte(n))
		steps[stepCount-1-i] = byte(3 - n)
	}
	data[0] = posX
	data[1] = posY
	return data
}

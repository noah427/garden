package sprite

import (

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func CreateSprite(image string, height, width int, x, y float64) *Sprite {
	img, _, _ := ebitenutil.NewImageFromFile("./assets/plants/basil.png")
	return &Sprite{
		Image:  img,
		Height: height,
		Width:  width,
		X:      x,
		Y:      y,
	}
}

func UpgradeSprite(sprite *Sprite, frameCount int) *AnimeSprite{
	return &AnimeSprite{
		Sprite: sprite,
		FrameCount: frameCount,
	}
}


func InBounds(x,y float64, X,Y float64, height,width int) bool{
	if X <= x && X+float64(width) >= x { // within x bounds
		if Y <= y && Y+float64(height) >= y { // within y bounds
			return true
		}
	}
	return false
}
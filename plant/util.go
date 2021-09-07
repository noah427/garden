package plant

import (
	"noah/garden/sprite"
)

func CreatePlant(image string, height, width int, x, y float64, frameCount int, name string) *Plant {
	return &Plant{
		AnimeSprite: sprite.UpgradeSprite(sprite.CreateSprite(image, height, width, x, y), frameCount),
		Name:        name,
	}
}

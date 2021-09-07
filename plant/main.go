package plant

import (
	"noah/garden/sprite"

	// "github.com/hajimehoshi/ebiten/v2"
)

type AnimeSprite = sprite.AnimeSprite

type Plant struct {
	*sprite.AnimeSprite
	Name string
}

func (P *Plant) Update(tps float64)error {
	return P.AnimeSprite.Update(tps)
}
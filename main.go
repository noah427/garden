package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"noah/garden/plant"
	"noah/garden/pot"
	"noah/garden/sprite"
)

const (
	screenWidth  = 1680 / 3
	screenHeight = 1050 / 3
)

// I hope this causes a memory leak or some shit
type ECS struct {
	Sprites []sprite.AnySprite
}

type Game struct {
	Plants *ECS
	Pots   []*pot.Pot
}

func (G *Game) FindPot(x, y float64) *sprite.DropZone {
	for _, pot := range G.Pots {
		if pot.DropZone.InBounds(x, y) {
			return pot.DropZone
		}
	}
	return nil
}

func (G *Game) Draw(screen *ebiten.Image) {
	for _, component := range G.Plants.Sprites {
		component.Draw(screen)
	}

	for _, pot := range G.Pots {
		pot.Draw(screen)
		pot.DropZone.DebugRect(screen)
	}
}

func (G *Game) Update() error {
	tps := ebiten.CurrentTPS()
	for _, component := range G.Plants.Sprites {
		component.Update(tps)
	}
	for _, pot := range G.Pots {
		pot.Update(tps)
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	basil := plant.CreatePlant("./assets/plants/basil.png", 32, 32, 0, 0, 5, "Basil")
	basil.Draggable = true

	dp := &pot.Pot{
		Sprite: sprite.CreateSprite("./assets/pots/defaultpot.png", 32, 32, 100, 100),
	}

	dp.CreateDropZone(0, 0, 32, 32, 0, -16)

	Plants := &ECS{
		[]sprite.AnySprite{basil},
	}

	game := Game{Plants: Plants, Pots: []*pot.Pot{dp}}

	basil.RequestDropZone = game.FindPot

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Garden")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(&game); err != nil {
		fmt.Println(err)
	}

}

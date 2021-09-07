package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"noah/garden/plant"
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
	Ecs *ECS
}

func (G *Game) Draw(screen *ebiten.Image) {
	for _, component := range G.Ecs.Sprites {
		component.Draw(screen)
	}
}

func (G *Game) Update() error {
	tps := ebiten.CurrentTPS()
	for _, component := range G.Ecs.Sprites {
		component.Update(tps)
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	basil := plant.CreatePlant("./assets/plants/basil.png", 32, 32, 0, 0, 5, "Basil")
	basil.Draggable = true

	Ecs := &ECS{
		[]sprite.AnySprite{basil},
	}

	game := Game{Ecs: Ecs}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Garden")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(&game); err != nil {
		fmt.Println(err)
	}

}

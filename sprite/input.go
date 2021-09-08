package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (S *Sprite) IsClicked() bool {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return false
	}
	a, b := ebiten.CursorPosition()
	x := float64(a)
	y := float64(b)

	
	if InBounds(x, y, S.X, S.Y, S.Height, S.Width){
		if S.Draggable {
			S.Dragging = true
			S.Drag()
		}
	}

	return false
}

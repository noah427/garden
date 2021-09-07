package sprite 
import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (S *Sprite) Drag(){
	x, y := ebiten.CursorPosition()
	// subtract half of Width and Height so it's held by center
	S.X = float64(x) - float64(S.Width/2)
	S.Y = float64(y) - float64(S.Height/2)
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		S.Dragging = false
	}
}
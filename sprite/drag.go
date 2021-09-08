package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (S *Sprite) Drag() {
	x, y := ebiten.CursorPosition()
	// subtract half of Width and Height so it's held by center
	nx := float64(x) - float64(S.Width/2)
	ny := float64(y) - float64(S.Height/2)
	S.MoveTo(nx,ny)
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if S.RequestDropZone != nil {
			zone := S.RequestDropZone(S.X,S.Y)
			if zone != nil {
				S.MoveTo(zone.SnapX,zone.SnapY)
			}
		}
		S.Dragging = false
	}
}

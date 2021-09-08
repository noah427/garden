package sprite

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type RequestDropZone func(x, y float64) *DropZone
type DropZone struct {
	Height, Width int
	X, Y          float64
	SnapX, SnapY  float64
}

func (S *Sprite) CreateDropZone(x, y float64, height, width int, snapX, snapY float64) {
	x += S.X
	y += S.Y
	snapX += S.X
	snapY += S.Y

	S.DropZone = &DropZone{}
	S.DropZone.initiate(x, y, height, width, snapX, snapY)
}

func (S *Sprite) MoveDropZone(x, y float64) {
	relX := S.DropZone.X - S.X
	relY := S.DropZone.Y - S.Y

	nx := relX + x
	ny := relY + y
	snapX := relX + x
	snapY := relY + y

	S.DropZone.X = nx
	S.DropZone.Y = ny

	S.DropZone.SnapX = snapX
	S.DropZone.SnapY = snapY
}

func (D *DropZone) initiate(x, y float64, height, width int, snapX, snapY float64) {
	D.X = x
	D.Y = y
	D.Height = height
	D.Width = width
	D.SnapX = snapX
	D.SnapY = snapY
}

func (D DropZone) DebugRect(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, D.X, D.Y, D.X+float64(D.Width), D.Y, color.White)                                     // top
	ebitenutil.DrawLine(screen, D.X, D.Y+float64(D.Height), D.X+float64(D.Width), D.Y+float64(D.Height), color.White) // bottom
	ebitenutil.DrawLine(screen, D.X, D.Y, D.X, D.Y+float64(D.Height), color.White)                                    // left side
	ebitenutil.DrawLine(screen, D.X+float64(D.Width), D.Y, D.X+float64(D.Width), D.Y+float64(D.Height), color.White)  // right side
}

func (D DropZone) InBounds(x, y float64) bool {
	mx, my := ebiten.CursorPosition()
	return InBounds(x, y, D.X, D.Y, D.Height, D.Width) || InBounds(float64(mx), float64(my), D.X, D.Y, D.Height, D.Width)
}

package sprite

import (
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const FPS = 12

type AnySprite interface {
	Update(tps float64) error
	Draw(screen *ebiten.Image) error
}

type Sprite struct {
	Image           *ebiten.Image
	Height, Width   int
	X, Y            float64
	Draggable       bool
	Dragging        bool
	DropZone        *DropZone
	RequestDropZone RequestDropZone
}

func (S *Sprite) MoveTo(x, y float64) {
	if S.DropZone != nil {
		S.MoveDropZone(x, y)
	}
	S.X = x
	S.Y = y
}

func (S *Sprite) Update(tps float64) error {
	S.IsClicked()
	if S.Dragging {
		S.Drag()
	}
	return nil
}

func (S *Sprite) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(S.X, S.Y)
	screen.DrawImage(S.Image, op)
	return nil
}

type AnimeSprite struct {
	*Sprite
	Frame, FrameCount int
	Tick              int
}

// assume each update is a new frame
// adjust to 12 fps
func (A *AnimeSprite) Update(tps float64) error {
	if tps == 0 {
		return nil
	}

	A.Sprite.Update(tps)
	A.Tick = (A.Tick + 1) % int(tps)
	if A.Tick%FPS == 0 {
		A.Frame = (A.Frame + 1) % A.FrameCount
	}
	return nil
}

func (A *AnimeSprite) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(A.X, A.Y)
	sx := 0 + (A.Width * A.Frame)
	sy := 0
	ex := sx + A.Width
	ey := sy + A.Height
	screen.DrawImage(A.Image.SubImage(image.Rect(sx, sy, ex, ey)).(*ebiten.Image), op)
	return nil
}

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
	Image         *ebiten.Image
	Height, Width int
	X, Y          float64
	Draggable     bool
	Dragging      bool
	DropZone      *DropZone
}

func (S *Sprite) Update() {
	S.IsClicked()
	if S.Dragging {
		S.Drag()
	}
}

func (S *Sprite) Draw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	sx := 0
	sy := 0
	ex := sx + S.Width
	ey := sy + S.Height
	screen.DrawImage(S.Image.SubImage(image.Rect(sx, sy, ex, ey)).(*ebiten.Image), op)
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

	A.Sprite.Update()
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

package sprite

type DropZone struct {
	Used          bool
	Height, Width int
	X, Y          float64
	SnapX, SnapY  float64
}

func (D *DropZone) Initiate(x, y float64, height, width int, snapX, snapY float64) {
	D.Used = true
	D.Height = height
	D.Width = width
	D.SnapX = snapX
	D.SnapY = snapY
}

func (D DropZone) InBounds(x, y float64) bool {
	if !D.Used{
		return false // potentially add error response on this function
	}
	return InBounds(x, y, D.X, D.Y, D.Height, D.Width)
}

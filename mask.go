package imagedraw

import (
	"image"
	"image/color"
)

type CircleMask struct {
	P image.Point
	R int
}

func (c *CircleMask) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *CircleMask) Bounds() image.Rectangle {
	return image.Rect(c.P.X-c.R, c.P.Y-c.R, c.P.X+c.R, c.P.Y+c.R)
}

func (c *CircleMask) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.P.X)+0.5, float64(y-c.P.Y)+0.5, float64(c.R)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}
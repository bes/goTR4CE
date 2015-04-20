package renderer

import (
	"math"
)

type Color struct {
	r, g, b float64
}

func NewColorRgb(r, g, b float64) *Color {
	result := new(Color)
	result.r = r
	result.g = g
	result.b = b

	return result
}

func NewColor() *Color {
	return new(Color)
}

func (c *Color) AddColor(r, g, b float64) {
	c.r += r
	c.g += g
	c.b += b
}

func clamp(val float64) float64 {
	return math.Max(0, math.Min(val, 254))
}

func (c *Color) GetRed() uint8 {
	return uint8(clamp(c.r))
}

func (c *Color) GetGreen() uint8 {
	return uint8(clamp(c.g))
}

func (c *Color) GetBlue() uint8 {
	return uint8(clamp(c.b))
}

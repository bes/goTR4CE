package renderer

import ()

type Light interface {
	GetPos() *Point3D
	GetColor() *Color
	GetIntensity() float64
}

type StandardLight struct {
	pos       *Point3D
	color     *Color
	intensity float64
}

func NewStandardLight(x, y, z float64, c *Color, i float64) *StandardLight {
	return &StandardLight{NewPoint3D(x, y, z), c, i}
}

func (s *StandardLight) GetPos() *Point3D {
	return s.pos
}

func (s *StandardLight) GetColor() *Color {
	return s.color
}

func (s *StandardLight) GetIntensity() float64 {
	return s.intensity
}

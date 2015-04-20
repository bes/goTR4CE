package renderer

import ()

type Plane struct {
	p                                     *Point3D
	d, a, diff, s, reflection, refraction float64
	c                                     *Color
}

func NewPlane(p *Point3D, d, a, diff, s, reflection, refraction float64, c *Color) *Plane {
	return &Plane{p, d, a, diff, s, reflection, refraction, c}
}

func (pl *Plane) Intersects(r *Ray, w *World) *Point3D {
	var i *Point3D
	dp := pl.p.Dot(r.GetVector())
	if dp != 0 {
		t := -(pl.p.Dot(r.GetPoint()) + pl.d) / dp
		if t > 0 {
			i = r.GetPoint().Plus(r.GetVector().Scale(t))
		}
	}
	return i
}

func (pl *Plane) GetColor() *Color {
	return pl.c
}

func (pl *Plane) GetNormal(p *Point3D) *Point3D {
	return pl.p
}

func (pl *Plane) Ambient() float64 {
	return pl.a
}

func (pl *Plane) Diffuse() float64 {
	return pl.diff
}

func (pl *Plane) Specular() float64 {
	return pl.s
}

func (pl *Plane) Refraction() float64 {
	return pl.refraction
}

func (pl *Plane) Reflection() float64 {
	return pl.reflection
}

func (pl *Plane) HasInvertNormal(r *Ray, p *Point3D) bool {
	return false
}

func (pl *Plane) GetInvertNormal(p *Point3D) *Point3D {
	return nil
}

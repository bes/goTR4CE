package renderer

import ()

type Plane struct {
	p0                                 *Point3D // Point on the plane
	n                                  *Point3D // Normal to the plane
	a, diff, s, reflection, refraction float64
	c                                  *Color
}

func NewPlane(p0, n *Point3D, a, diff, s, reflection, refraction float64, c *Color) *Plane {
	return &Plane{p0, n.Normalize(), a, diff, s, reflection, refraction, c}
}

func (pl *Plane) Intersects(r *Ray, w *World) *Point3D {
	// N plane normal
	// Q point on plane
	// E eye point
	// D eye Vector

	// N *(Q - E) / N * D

	var i *Point3D
	dp := pl.n.Dot(r.GetVector())
	if dp != 0 {
		t := pl.p0.Minus(r.GetPoint()).Dot(pl.n) / dp
		if t >= 0 {
			i = r.GetPoint().Plus(r.GetVector().Scale(t))
		}
	}
	return i
}

func (pl *Plane) GetColor() *Color {
	return pl.c
}

func (pl *Plane) GetNormal(p *Point3D) *Point3D {
	return pl.n
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

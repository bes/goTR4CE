package renderer

import (
	//"fmt"
	"math"
)

type Sphere struct {
	pos *Point3D
	r   float64
	c   *Color

	a, d, s, reflect, refract float64
}

func NewSphere(p *Point3D, a, d, s, r, refract, reflect float64, c *Color) *Sphere {
	return &Sphere{p, r, c, a, d, s, reflect, refract}
}

func (s *Sphere) GetColor() *Color {
	return s.c
}

func (s *Sphere) Intersects(ray *Ray, w *World) *Point3D {
	unit := ray.GetVector().Normalize()
	p := ray.GetPoint()

	alpha := -p.Minus(s.pos).Dot(unit)
	q := p.Plus(unit.Scale(alpha))
	bSq := q.Minus(s.pos).AbsSquared()

	if bSq > s.r*s.r {
		return nil
	}

	a := math.Sqrt(s.r*s.r - bSq)

	//fmt.Println(alpha, ",", a)
	if alpha >= a {
		return q.Minus(unit.Scale(a))
	}

	if alpha+a > 0 {
		return q.Plus(unit.Scale(a))
	}

	return nil
}

func (s *Sphere) GetNormal(point *Point3D) *Point3D {
	return point.Minus(s.pos).Normalize()
}

func (s *Sphere) Ambient() float64 {
	return s.a
}

func (s *Sphere) Diffuse() float64 {
	return s.d
}

func (s *Sphere) Specular() float64 {
	return s.s
}

func (s *Sphere) Refraction() float64 {
	return s.refract
}

func (s *Sphere) Reflection() float64 {
	return s.reflect
}

func (s *Sphere) HasInvertNormal(ray *Ray, point *Point3D) bool {
	halfWay := ray.GetPoint().Plus(ray.GetVector().Scale(ray.GetPoint().Distance(point) / 2))
	if halfWay.Distance(s.pos) < s.r {
		return true
	}
	return false
}

func (s *Sphere) GetInvertNormal(point *Point3D) *Point3D {
	return s.pos.Minus(point).Normalize()
}

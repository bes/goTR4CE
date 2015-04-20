package renderer

import (
	"math"
)

type Point3D struct {
	point [3]float64
}

func CopyPoint3D(p *Point3D) *Point3D {
	result := new(Point3D)

	copy(result.point[:], p.point[:])
	return result
}

func NewPoint3D(x, y, z float64) *Point3D {
	result := new(Point3D)
	result.point[0] = x
	result.point[1] = y
	result.point[2] = z
	return result
}

func (p *Point3D) Dot(b *Point3D) float64 {
	return p.point[0]*b.point[0] + p.point[1]*b.point[1] + p.point[2]*b.point[2]
}

func (p *Point3D) Minus(b *Point3D) *Point3D {
	return NewPoint3D(p.point[0]-b.point[0], p.point[1]-b.point[1], p.point[2]-b.point[2])
}

func (p *Point3D) Plus(b *Point3D) *Point3D {
	return NewPoint3D(p.point[0]+b.point[0], p.point[1]+b.point[1], p.point[2]+b.point[2])
}

func (p *Point3D) Mul(b *Point3D) *Point3D {
	return NewPoint3D(p.point[0]*b.point[0], p.point[1]*b.point[1], p.point[2]*b.point[2])
}

func (p *Point3D) Scale(scale float64) *Point3D {
	return NewPoint3D(p.point[0]*scale, p.point[1]*scale, p.point[2]*scale)
}

func (p *Point3D) Abs() float64 {
	return math.Sqrt(p.AbsSquared())
}

func (p *Point3D) AbsSquared() float64 {
	return p.point[0]*p.point[0] + p.point[1]*p.point[1] + p.point[2]*p.point[2]
}

func (p *Point3D) Normalize() *Point3D {
	abs := p.Abs()
	return NewPoint3D(p.point[0]/abs, p.point[1]/abs, p.point[2]/abs)
}

func (p *Point3D) Cross(b *Point3D) *Point3D {
	return NewPoint3D(p.point[1]*b.point[2]-p.point[2]*b.point[1],
		p.point[2]*b.point[0]-p.point[0]*b.point[2],
		p.point[0]*b.point[1]-p.point[1]*b.point[0])
}

func (p *Point3D) Distance(b *Point3D) float64 {
	return p.Minus(b).Abs()
}

func (p *Point3D) GetX() float64 {
	return p.point[0]
}

func (p *Point3D) GetY() float64 {
	return p.point[1]
}

func (p *Point3D) GetZ() float64 {
	return p.point[2]
}

func (p *Point3D) RotateY90CCW() *Point3D {
	return NewPoint3D(p.point[2], p.point[1], -p.point[0])
}

func (p *Point3D) RotateX90CCW() *Point3D {
	return NewPoint3D(p.point[0], -p.point[2], p.point[1])
}

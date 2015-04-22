package renderer

import (
	"math"
)

type Raster struct {
	width, height             int
	fov                       float64
	pos, eyePos, eyeDirection *Point3D
	xV, yV                    *Point3D
}

// fov in radians
func NewRaster(width, height int, rasterPos, eyeDirection *Point3D, fov float64) *Raster {
	result := new(Raster)
	result.width = width
	result.height = height
	result.fov = fov

	dist := (float64(width) / 2) / math.Tan(fov/2)

	result.pos = rasterPos
	result.eyePos = rasterPos.Minus(eyeDirection.Scale(dist))
	result.eyeDirection = eyeDirection

	// TODO: I don't think this is correct
	result.xV = eyeDirection.RotateY90CCW()
	result.yV = eyeDirection.RotateX90CCW()

	return result
}

func (r *Raster) GetWidth() int {
	return r.width
}

func (r *Raster) GetHeight() int {
	return r.height
}

func (r *Raster) GetXV() *Point3D {
	return r.xV
}

func (r *Raster) GetYV() *Point3D {
	return r.yV
}

func (r *Raster) GetRay(x, y int) *Ray {
	xlen := (-float64(r.width) / 2) + float64(x)
	ylen := (-float64(r.height) / 2) + float64(y)
	pointPos := r.pos.Plus(r.xV.Scale(float64(xlen))).Plus(r.yV.Scale(float64(ylen)))
	rayDirection := pointPos.Minus(r.eyePos).Normalize()
	return NewRay(r.eyePos, rayDirection)
}

package renderer

import (
	"math"
)

type Raster struct {
	width, height int
	fov           float64
	eye           *Eye
	pos, xV, yV   *Point3D
}

// fov in radians
func NewRaster(width, height int, eye *Eye, fov float64) *Raster {
	result := new(Raster)
	result.width = width
	result.height = height
	result.eye = eye
	result.fov = fov

	dist := (float64(width) / 2) / math.Tan(fov/2)

	result.pos = eye.GetPos().Plus(eye.GetDirection().Scale(dist))

	result.xV = eye.GetDirection().RotateY90CCW()
	result.yV = eye.GetDirection().RotateX90CCW()

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

func (r *Raster) GetPoint(x, y int) *Point3D {
	xlen := (-float64(r.width) / 2) + float64(x)
	ylen := (-float64(r.height) / 2) + float64(y)
	return r.pos.Plus(r.xV.Scale(float64(xlen))).Plus(r.yV.Scale(float64(ylen)))
}

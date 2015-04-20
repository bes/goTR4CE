package renderer

import ()

type Ray struct {
	point, vector *Point3D
}

func NewRayValues(x, y, z, xV, yV, zV float64) *Ray {
	result := new(Ray)
	result.point = NewPoint3D(xV, yV, zV)
	result.vector = NewPoint3D(x, y, z)
	return result
}

func NewRay(point, vector *Point3D) *Ray {
	result := new(Ray)
	result.point = point
	result.vector = vector
	return result
}

func (r *Ray) GetVector() *Point3D {
	return r.vector
}

func (r *Ray) GetPoint() *Point3D {
	return r.point
}

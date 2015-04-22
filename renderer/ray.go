package renderer

import ()

type Ray struct {
	point, vector *Point3D
}

func NewRayValues(x, y, z, xV, yV, zV float64) *Ray {
	return &Ray{NewPoint3D(xV, yV, zV), NewPoint3D(x, y, z)}
}

func NewRay(point, vector *Point3D) *Ray {
	return &Ray{point, vector}
}

func (r *Ray) GetVector() *Point3D {
	return r.vector
}

func (r *Ray) GetPoint() *Point3D {
	return r.point
}

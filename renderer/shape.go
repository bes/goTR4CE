package renderer

import ()

type Shape interface {
	Intersects(r *Ray, w *World) *Point3D
	GetColor() *Color
	GetNormal(*Point3D) *Point3D
	Ambient() float64
	Diffuse() float64
	Specular() float64
	Refraction() float64
	Reflection() float64
	HasInvertNormal(*Ray, *Point3D) bool
	GetInvertNormal(*Point3D) *Point3D
}

package renderer

import ()

type Eye struct {
	pos, direction *Point3D
}

func NewEye(pos, direction *Point3D) *Eye {
	return &Eye{pos, direction}
}

func (e *Eye) GetPos() *Point3D {
	return e.pos
}

func (e *Eye) GetDirection() *Point3D {
	return e.direction
}

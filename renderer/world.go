package renderer

import (
	"math"
)

type World struct {
	shapes []Shape
	lights []Light

	raster *Raster
	eye    *Eye

	width, height, depth int
}

func NewWorld(width, height, depth int) *World {
	result := new(World)

	result.shapes = make([]Shape, 0, 10)
	result.lights = make([]Light, 0, 10)

	result.width = width
	result.height = height
	result.depth = depth

	return result
}

func (w *World) GetWidth() int {
	return w.width
}

func (w *World) GetHeight() int {
	return w.height
}

func (w *World) GetDepth() int {
	return w.depth
}

func (w *World) AddShape(s Shape) {
	w.shapes = append(w.shapes, s)
}

func (w *World) AddLight(l Light) {
	w.lights = append(w.lights, l)
}

func (w *World) SetRaster(r *Raster) {
	w.raster = r
}

func (w *World) SetEye(e *Eye) {
	w.eye = e
}

// rangeStart is inclusve, rangeEnd is exclusive
func (w *World) Render(ch chan *Color, rangeStart, rangeEnd int) {

	yStart := int(rangeStart / w.GetWidth())
	yEnd := int(rangeEnd / w.GetWidth())

	xStart := rangeStart % w.raster.GetWidth()
	xEnd := rangeEnd % w.raster.GetWidth()
	if xEnd > 0 {
		yEnd++
	}

	progress := rangeStart
	for y := yStart; y < yEnd; y++ {
		for x := xStart; x < w.raster.GetWidth(); x++ {
			xStart = 0
			rasterPos := w.raster.GetPoint(x, y)

			mc := NewColor()

			colorSet := false

			// This antialiasing should move into the trace method, since here it will only provide one level of AA
			startAa := -0.4
			endAa := 0.4
			stepAa := 0.4
			contribAa := ((endAa - startAa) / stepAa) * ((endAa - startAa) / stepAa)
			for i := startAa; i < endAa; i += stepAa {
				for j := startAa; j < endAa; j += stepAa {
					r := NewRay(w.eye.GetPos().Plus(w.raster.GetXV().Scale(i).Plus(w.raster.GetYV().Scale(j))),
						rasterPos.Minus(w.eye.GetPos()).Normalize())

					tc := w.trace(r, 3, 1)
					if tc != nil {
						colorSet = true
						mc.AddColor(tc.r/contribAa, tc.g/contribAa, tc.b/contribAa)
					}
				}
			}

			if colorSet {
				ch <- mc
			} else {
				ch <- nil
			}
			progress++
			if progress == rangeEnd {
				break
			}
		}
	}
}

func (w *World) trace(r *Ray, cutoff int, nju1 float64) *Color {
	if cutoff == 0 {
		return NewColorRgb(0, 0, 0)
	}
	s, point := w.closestIntersection(r)

	if s != nil {
		mc := NewColor()
		var N *Point3D
		hasInvertedNormal := s.HasInvertNormal(r, point)
		if hasInvertedNormal {
			N = s.GetInvertNormal(point)
		} else {
			N = s.GetNormal(point)
		}

		L := r.GetPoint().Minus(point).Normalize()
		R := N.Scale(N.Dot(L) * 2).Minus(L)

		var red, green, blue float64
		for _, l := range w.lights {
			ls := l.GetPos()
			rs := NewRay(point, ls.Minus(point).Normalize())

			if !w.hasIntersection(rs) {
				La := rs.GetVector().Normalize()
				LN := La.Dot(N)

				V := w.eye.GetPos().Minus(point).Normalize()

				rvAlphaPow := alphaPow(math.Max(0, R.Dot(V)), 8)

				red += (s.Diffuse()*LN*l.GetColor().r*l.GetIntensity() + s.Diffuse()*LN*s.GetColor().r + s.Specular()*rvAlphaPow*l.GetColor().r*l.GetIntensity())
				green += (s.Diffuse()*LN*l.GetColor().g*l.GetIntensity() + s.Diffuse()*LN*s.GetColor().g + s.Specular()*rvAlphaPow*l.GetColor().g*l.GetIntensity())
				blue += (s.Diffuse()*LN*l.GetColor().b*l.GetIntensity() + s.Diffuse()*LN*s.GetColor().b + s.Specular()*rvAlphaPow*l.GetColor().b*l.GetIntensity())
			}

		}

		if s.Reflection() > 0 {
			c := w.trace(NewRay(point, R), cutoff-1, 1)
			red += s.Reflection() * c.r
			green += s.Reflection() * c.g
			blue += s.Reflection() * c.b
		}

		if s.Refraction() > 0 {
			nju2 := s.Refraction()
			if hasInvertedNormal {
				nju2 = 1
				nju1 = s.Refraction()
			}

			nju := nju1 / nju2
			c1 := NewPoint3D(0, 0, 0).Minus(N).Dot(r.GetVector())
			c2 := math.Sqrt(1 - nju*nju*(1-c1*c1))

			// There are some other variants in the java code
			t := r.GetVector().Scale(nju).Plus(N.Scale(nju*c1 - c2))

			c := w.trace(NewRay(point.Plus(t.Scale(1)), t), cutoff-1, nju2)

			red += c.r
			green += c.g
			blue += c.b
		}

		mc.AddColor(s.Ambient()*s.GetColor().r+red,
			s.Ambient()*s.GetColor().g+green,
			s.Ambient()*s.GetColor().b+blue)

		return mc
	}

	return NewColorRgb(0, 0, 0)
}

func (w *World) hasIntersection(r *Ray) bool {
	for _, s := range w.shapes {
		if s.Intersects(r, w) != nil {
			return true
		}
	}
	return false
}

func (w *World) closestIntersection(r *Ray) (Shape, *Point3D) {
	var hit Shape
	var hitPoint *Point3D
	zd := math.MaxFloat64

	for _, s := range w.shapes {
		point := s.Intersects(r, w)
		if point != nil {
			tzd := point.Minus(r.GetPoint()).Abs()
			if zd > tzd {
				zd = tzd
				hit = s
				var normal *Point3D
				if s.HasInvertNormal(r, point) {
					normal = s.GetInvertNormal(point)
				} else {
					normal = s.GetNormal(point)
				}
				hitPoint = point.Plus(normal.Scale(0.00005))
			}
		}
	}
	return hit, hitPoint
}

func alphaPow(val float64, pow int) float64 {
	temp := val
	for i := 1; i < pow; i++ {
		temp *= val
	}
	return temp
}

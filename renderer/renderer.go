package renderer

import (
	//	"fmt"
	"math"
	"runtime"
)

func RunRender(colorChans chan *ColorData, depth, width, height int) {
	w := NewWorld(width, height, depth)

	//                                 Sphere: p, a, d, s, r, refract, reflect float64, c *Color
	w.AddShape(NewSphere(NewPoint3D(-80, 0, 700), 0.5, 0.3, 0.5, 50, 0, 0, NewColorRgb(255, 255, 0)))
	w.AddShape(NewSphere(NewPoint3D(80, 0, 800), 0.5, 0.3, 0.5, 50, 0, 0, NewColorRgb(255, 0, 255)))
	w.AddShape(NewSphere(NewPoint3D(0, 10, 400), 0, 0.1, 0.1, 60, 2, 0, NewColorRgb(255, 255, 255)))
	w.AddShape(NewSphere(NewPoint3D(80, 200, 1600), 0.1, 0.4, 0.2, 250, 0, 0.5, NewColorRgb(120, 120, 120)))

	w.AddShape(NewPlane(NewPoint3D(0, -50, 0), NewPoint3D(0, 1, 0), 0.4, 0.9, 0.3, 0, 0, NewColorRgb(100, 90, 100)))

	w.AddLight(NewStandardLight(0, 200, 500, NewColorRgb(255, 0, 0), 0.5))
	//w.AddLight(NewStandardLight(-100, 100, 50, NewColorRgb(255, 255, 255), 0.5))
	w.AddLight(NewStandardLight(-110, 90, 160, NewColorRgb(0, 0, 255), 0.2))
	w.AddLight(NewStandardLight(80, 100, 900, NewColorRgb(255, 255, 255), 0.1))
	w.AddLight(NewStandardLight(80, 30, 300, NewColorRgb(0, 200, 0), 0.2))
	w.AddLight(NewStandardLight(0, 10, 1000, NewColorRgb(0, 255, 255), 0.2))
	//w.AddLight(NewStandardLight(-80, 100, 100, NewColorRgb(255, 255, 255), 0.5))

	w.SetRaster(NewRaster(width, height, NewPoint3D(0, 100, 0), NewPoint3D(0, 0, 1).Normalize(), math.Pi/4))

	go func() {
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				go w.Render(colorChans, x, y)
			}
			runtime.Gosched()
		}
	}()
}

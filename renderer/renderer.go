package renderer

import (
	"math"
)

func RunRender(colorChans []chan *Color, width, height int) (int, int) {
	depth := 50
	w := NewWorld(width, height, depth)

	w.AddShape(NewSphere(NewPoint3D(0, 50, 360), 0, 0.3, 0, 100, 1.7, 0, NewColorRgb(255, 255, 255)))
	w.AddShape(NewSphere(NewPoint3D(-110, 15, 160), 0.1, 0.5, 0.2, 30, 0, 0.1, NewColorRgb(255, 255, 0)))
	w.AddShape(NewSphere(NewPoint3D(80, 30, 180), 0.1, 0.4, 0.2, 60, 0, 0, NewColorRgb(255, 255, 255)))

	w.AddShape(NewPlane(NewPoint3D(0, 1, 0), 50, 0, 0.9, 0.3, 0, 0, NewColorRgb(255, 0, 0)))

	w.AddLight(NewStandardLight(-110, 90, 160, NewColorRgb(0, 0, 255), 0.2))
	w.AddLight(NewStandardLight(80, 100, 900, NewColorRgb(255, 255, 255), 0.5))
	w.AddLight(NewStandardLight(80, 30, 300, NewColorRgb(255, 0, 0), 0.2))
	w.AddLight(NewStandardLight(0, 10, 1000, NewColorRgb(255, 255, 255), 0.2))
	w.AddLight(NewStandardLight(-80, 100, 100, NewColorRgb(255, 255, 255), 0.5))

	e := NewEye(NewPoint3D(0, 300, 2000), NewPoint3D(-0.2, -0.2, -1.8))
	w.SetEye(e)
	w.SetRaster(NewRaster(width, height, e, math.Pi/4))

	numChans := len(colorChans)
	totalPixels := width * height
	numPixelsPerChan := int(totalPixels / numChans)

	rangeStart := 0
	rangeEnd := numPixelsPerChan
	for i := 0; i < numChans-1; i++ {
		// rangeStart is inclusive
		// rangeEnd is exclusive
		go w.Render(colorChans[i], rangeStart, rangeEnd)
		rangeStart = rangeEnd
		rangeEnd += numPixelsPerChan
	}

	go w.Render(colorChans[numChans-1], rangeStart, totalPixels)

	return numChans, numPixelsPerChan
}

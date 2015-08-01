package renderer

type ColorData struct {
	color  *Color
	offset uint32
}

func NewColorData(color *Color, offset uint32) *ColorData {
	return &ColorData{color, offset}
}

func (cd *ColorData) GetColor() *Color {
	return cd.color
}

func (cd *ColorData) GetOffset() uint32 {
	return cd.offset
}

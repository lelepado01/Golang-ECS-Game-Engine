package colorlib

//ColorRGBA ...
type ColorRGBA struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

//CreateColor ...
func CreateColor(r uint8, g uint8, b uint8, a uint8) ColorRGBA {
	return ColorRGBA{
		r: r,
		g: g,
		b: b,
		a: a}
}

//GetValues ...
func (c *ColorRGBA) GetValues() (uint8, uint8, uint8, uint8) {
	return c.r, c.g, c.b, c.a
}

//GetValue ...
func (c *ColorRGBA) GetValue(i int) uint8 {
	switch i {
	case 0:
		return c.r
	case 1:
		return c.g
	case 2:
		return c.b
	case 3:
		return c.a
	default:
		return 0
	}
}

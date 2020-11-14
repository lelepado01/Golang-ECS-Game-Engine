package mathlib

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Rectangle is a struxt containing the origin and the dimensions of the rectangle
type Rectangle struct {
	origin Vector2
	size   Vector2
}

// CreateRectangle returns a new triangle initialized with the given parameters
func CreateRectangle(x, y, w, h float64) Rectangle {
	var t Rectangle
	t.origin = Vec2(x, y)
	t.size = Vec2(w, h)

	return t
}

// GetSDLRectangle returns the caller's corresponding sdl Rect
func (r *Rectangle) GetSDLRectangle() *sdl.Rect {
	return &sdl.Rect{
		X: r.origin.GetXasInt32(),
		Y: r.origin.GetYasInt32(),
		W: r.size.GetXasInt32(),
		H: r.size.GetYasInt32(),
	}
}

// Translate moves the rectangle by the coords given
func (r *Rectangle) Translate(x, y float64) {
	r.origin.Translate(x, y)
}

// GetCenterX returns the X coord at the center of the rectangle
func (r *Rectangle) GetCenterX() float64 {
	return r.origin.GetX() + r.size.GetX()/2
}

// GetCenterY returns the Y coord at the center of the rectangle
func (r *Rectangle) GetCenterY() float64 {
	return r.origin.GetY() + r.size.GetY()/2
}

// Contains returns true if the given point is inside the rectangle shape
func (r *Rectangle) Contains(x, y float64) bool {
	if x > r.origin.GetX() && x < r.origin.GetX()+r.size.GetX() {
		if y > r.origin.GetY() && y < r.origin.GetY()+r.size.GetY() {
			return true
		}
	}

	return false
}

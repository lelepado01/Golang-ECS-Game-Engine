package mathlib

// Vector2 is a struct containing two components, x and y
type Vector2 struct {
	x, y float64
}

// Vec2 returns a new Vector2 obj initialized with x and y
func Vec2(x float64, y float64) Vector2 {
	return Vector2{x: x, y: y}
}

// GetX returns the x component of the vector
func (v *Vector2) GetX() float64 {
	return v.x
}

// GetY returns the y component of the vector
func (v *Vector2) GetY() float64 {
	return v.y
}

// GetXasInt32 returns the x component of the vector as int32
func (v *Vector2) GetXasInt32() int32 {
	return int32(v.x)
}

// GetYasInt32 returns the y component of the vector as int32
func (v *Vector2) GetYasInt32() int32 {
	return int32(v.y)
}

// Translate moves the vector by the given coords
func (v *Vector2) Translate(x, y float64) {
	v.x += x
	v.y += y
}

// Mult multiplies the vector by the given x and y values correspondingly
func (v *Vector2) Mult(x, y float64) {
	v.x *= x
	v.y *= y
}

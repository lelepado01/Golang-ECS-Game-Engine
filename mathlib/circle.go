package mathlib

import "math"

// Circle is a struct containing useful methods for distance
// and collisions with circle shapes
type Circle struct {
	origin Vector2
	radius float64
}

// CreateCircle returns a new circle with the origin and the radius given
func CreateCircle(x, y, r float64) Circle {
	c := Circle{Vec2(x, y), r}

	return c
}

// DistanceFromOrigin returns the distance between the given pointand
// the center of the circle
func (c *Circle) DistanceFromOrigin(x, y float64) float64 {
	return math.Sqrt(math.Pow(c.origin.GetX()-x, 2) + math.Pow(c.origin.GetY()-y, 2))
}

// Contains returns true if the point given is inside the circle shape
func (c *Circle) Contains(x, y float64) bool {
	return c.DistanceFromOrigin(x, y) >= c.radius
}

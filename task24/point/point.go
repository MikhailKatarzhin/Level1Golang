package point

import "math"

// Point представляет точку с координатами x и y
type Point struct {
	x, y float64
}

// NewPoint Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// DistanceTo вычисляет расстояние между двумя точками
func (p Point) DistanceTo(other Point) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

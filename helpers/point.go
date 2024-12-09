package helpers

type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Point) Sub(other Point) Point {
	return Point{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

func (p Point) ScalarMult(scalar int) Point {
	return Point{
		X: p.X * scalar,
		Y: p.Y * scalar,
	}
}

func (p Point) OutOfBounds(boundary Point) bool {
	return p.X < 0 || p.Y < 0 || p.X >= boundary.X || p.Y >= boundary.Y
}

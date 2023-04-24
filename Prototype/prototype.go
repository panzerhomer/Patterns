package creational

// Figure is the prototype
type Figure interface {
	Clone() Figure
	GetSquare() float64
}

// Rectangle is the implementation of the Figure
type Rectangle struct {
	a, b float64
}

func (r *Rectangle) GetSquare() float64 {
	return r.a * r.b
}

func (r *Rectangle) Clone() Figure {
	return &Rectangle{r.a, r.b}
}

func NewRectangle(a, b float64) Figure {
	return &Rectangle{
		a: a,
		b: b,
	}
}

// Circle is the implementation of the Figure
type Circle struct {
	r float64
}

func (c *Circle) GetSquare() float64 {
	return c.r * c.r * 3.14
}

func (c *Circle) Clone() Figure {
	return &Circle{c.r}
}

func NewCircle(r float64) Figure {
	return &Circle{
		r: r,
	}
}

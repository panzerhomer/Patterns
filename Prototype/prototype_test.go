package creational

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	rectangle1 := NewRectangle(2, 4)
	rectangle2 := rectangle1.Clone()

	if rectangle1.GetSquare() != rectangle2.GetSquare() {
		t.Error("Expect name \"A\" to equal, but not equal.")
	}

	circle1 := NewCircle(2)
	circle2 := circle1.Clone()

	if circle1.GetSquare() != circle2.GetSquare() {
		t.Error("Expect name \"A\" to equal, but not equal.")
	}

}

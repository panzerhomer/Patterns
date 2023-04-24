package creational

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	swordman := NewCharacter(&SwordmanFactory{})

	swordman.Move()
	swordman.Hit()

	if swordman.Hit() == 0 {
		t.Errorf("Expect volume to %d, but got 0", swordman.Hit())
	}
}

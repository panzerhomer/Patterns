package creational

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	asserts := []string{"Mars", "KitKat", "Kinder"}

	factory := NewChocolateCreator()
	bars := []ChocolateBar{
		factory.CreateBar(Mars),
		factory.CreateBar(KitKat),
		factory.CreateBar(Kinder),
	}

	for i, bar := range bars {
		if name := bar.Name(); name != asserts[i] {
			t.Errorf("Expect action to %s, but %s.\n", asserts[i], name)
		}
	}
}

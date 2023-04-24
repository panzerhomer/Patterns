package creational

import "log"

type bar string

const (
	Mars   bar = "Mars"
	Kinder bar = "Kinder"
	KitKat bar = "KitKat"
)

type Creator interface {
	CreateBar(name bar) ChocolateBar
}

type ChocolateBar interface {
	Name() string
}

type MarsBar struct {
	name bar
}

func (m MarsBar) Name() string {
	return string(m.name)
}

type KitKatBar struct {
	name bar
}

func (k KitKatBar) Name() string {
	return string(k.name)
}

type KinderBar struct {
	name bar
}

func (k KinderBar) Name() string {
	return string(k.name)
}

type ChocolateCreator struct{}

func NewChocolateCreator() Creator {
	return &ChocolateCreator{}
}

func (c *ChocolateCreator) CreateBar(name bar) ChocolateBar {
	var product ChocolateBar

	switch name {
	case Mars:
		product = &MarsBar{name}
	case KitKat:
		product = &KitKatBar{name}
	case Kinder:
		product = &KinderBar{name}
	default:
		log.Println("This name is not implemented yet")
	}

	return product
}

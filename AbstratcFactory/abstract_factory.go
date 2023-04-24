package creational

import "fmt"

// Weapon is the abstract product
type Weapon interface {
	Hit() int
}

// Movement is the abstract product
type Movement interface {
	Move()
}

// Sword is implemented the product for Weapon
type Sword struct {
	damage int
}

func (s Sword) Hit() int {
	fmt.Println("Hit with the sword")
	return s.damage
}

// Bow is the implementation of the product for Weapon
type Bow struct {
	damage int
}

func (b Bow) Hit() int {
	fmt.Println("Shoot from the bow")
	return b.damage
}

// WalkMovement is the implementation of the product for Movement
type WalkMovement struct {
}

func (w WalkMovement) Move() {
	fmt.Println("The hero is walking")
}

// RunMovement is the implementation of the product for Movement
type RunMovement struct {
	speed int
}

func (r RunMovement) Move() {
	fmt.Println("The hero is running")
}

// CharacterFactory is the Abstract factory
type CharacterFactory interface {
	CreateMovement() Movement
	CreateWeapon() Weapon
}

// SwordmanFactory is the implementation of the abstract factory CharacterFactory
type SwordmanFactory struct {
}

func (s SwordmanFactory) CreateMovement() Movement {
	return RunMovement{speed: 20}
}

func (s SwordmanFactory) CreateWeapon() Weapon {
	return Sword{damage: 15}
}

// ArcherFactory is the implementation of the abstract factory CharacterFactory
type ArcherFactory struct {
}

func (a ArcherFactory) CreateMovement() Movement {
	return RunMovement{speed: 30}
}

func (a ArcherFactory) CreateWeapon() Weapon {
	return Bow{damage: 10}
}

// Character is the client of the abstract factory CharacterFactory
type Character struct {
	weapon   Weapon
	movement Movement
}

func (c *Character) Move() {
	c.movement.Move()
}

func (c *Character) Hit() int {
	return c.weapon.Hit()
}

func NewCharacter(factory CharacterFactory) *Character {
	character := Character{
		weapon:   factory.CreateWeapon(),
		movement: factory.CreateMovement(),
	}
	return &character
}

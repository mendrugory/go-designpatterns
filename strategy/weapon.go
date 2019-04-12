package strategy

import (
	"fmt"
)

// Weapon interface
type Weapon interface {
	Attack() string
}

// Sword struct
type Sword struct {
	Name string
}

// Axe struct
type Axe struct {
	Name string
}

// Poison struct
type Poison struct {
	Composition string
	Name        string
}

// Knife struct
type Knife struct {
	Name string
}

// Attack implementations for Sword
func (s Sword) Attack() string {
	return fmt.Sprintf("Sword %s !!!", s.Name)
}

// Attack implementations for Axe
func (a Axe) Attack() string {
	return fmt.Sprintf("Axe %s !!!", a.Name)
}

// Attack implementations for Poison
func (p Poison) Attack() string {
	return fmt.Sprintf("Poison %s !!!", p.Name)
}

// Attack implementations for Knife
func (k Knife) Attack() string {
	return fmt.Sprintf("Knife %s !!!", k.Name)
}

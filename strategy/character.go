package strategy

import (
	"fmt"
)

// Fighter interface
type Fighter interface {
	Fight() string
}

// Knight struct
type Knight struct {
	Name   string
	Weapon Weapon
}

// King struct
type King struct {
	Name   string
	Family string
	Weapon Weapon
}

// Queen struct
type Queen struct {
	Name   string
	Family string
	Weapon Weapon
}

// Sorcerer  struct
type Sorcerer struct {
	Name         string
	Congregation string
	Weapon       Weapon
}

// Fight implementation for Knife
func (k Knight) Fight() string {
	if k.Weapon == nil {
		return fmt.Sprintf("Kight %s does not have any weapon, provide one !!", k.Name)
	}
	return fmt.Sprintf("Kight %s fights with %s", k.Name, k.Weapon.Attack())
}

// SetWeapon sets the Weapon for Knight
func (k *Knight) SetWeapon(w Weapon) {
	k.Weapon = w
}

// Fight implementation for King
func (k King) Fight() string {
	if k.Weapon == nil {
		return fmt.Sprintf("King %s does not have any weapon, provide one !!", k.Name)
	}
	return fmt.Sprintf("King %s fights with %s", k.Name, k.Weapon.Attack())
}

// SetWeapon sets the Weapon for King
func (k *King) SetWeapon(w Weapon) {
	k.Weapon = w
}

// Fight implementation for Queen
func (q Queen) Fight() string {
	if q.Weapon == nil {
		return fmt.Sprintf("Queen %s does not have any weapon, provide one !!", q.Name)
	}
	return fmt.Sprintf("Queen %s fights with %s", q.Name, q.Weapon.Attack())
}

// SetWeapon sets the Weapon for Queen
func (q *Queen) SetWeapon(w Weapon) {
	q.Weapon = w
}

// Fight implementation for Sorcerer
func (s Sorcerer) Fight() string {
	if s.Weapon == nil {
		return fmt.Sprintf("Queen %s does not have any weapon, provide one !!", s.Name)
	}
	return fmt.Sprintf("Queen %s fights with %s", s.Name, s.Weapon.Attack())
}

// SetWeapon sets the Weapon for Sorcerer
func (s *Sorcerer) SetWeapon(w Weapon) {
	s.Weapon = w
}

// Fight a fighter attacks
func Fight(f Fighter) {
	fmt.Println(f.Fight())
}

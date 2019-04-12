package strategy

import (
	"fmt"
	"testing"
)

func TestKnightSetWeapon(t *testing.T) {
	knight := Knight{Name: "Rodri"}
	sword := Sword{Name: "Tizonadora"}
	knight.SetWeapon(sword)

	if knight.Weapon != sword {
		t.Errorf("The weapons are different.\n")
	}
}

func TestKnightSwitchWeapon(t *testing.T) {
	sword := Sword{Name: "Tizonadora"}
	knight := Knight{Name: "Rodri", Weapon: sword}
	knife := Knife{Name: "Cuchillito"}
	knight.SetWeapon(knife)

	if knight.Weapon != knife {
		t.Errorf("The weapons are different.\n")
	}
}

func TestQueenTryingWeapons(t *testing.T) {
	knife := Knife{Name: "Cuchillito"}
	queen := Queen{Name: "Isabelita", Weapon: knife}
	if queen.Weapon != knife {
		t.Errorf("The weapon is not the knife.\n")
	}

	sword := Sword{Name: "Rebanadora"}
	queen.SetWeapon(sword)
	if queen.Weapon != sword {
		t.Errorf("The weapon is not the knife.\n")
	}

	axe := Axe{Name: "Brutica"}
	queen.SetWeapon(axe)
	if queen.Weapon != axe {
		t.Errorf("The weapon is not the axe.\n")
	}

	poison := Poison{Name: "Doloroso"}
	queen.SetWeapon(poison)
	if queen.Weapon != poison {
		t.Errorf("The weapon is not the poison.\n")
	}
}

func TestKingFighting(t *testing.T) {
	kingName := "Fernan"
	axe := Axe{Name: "Brutica"}
	king := King{Name: kingName, Weapon: axe}
	if king.Fight() != fmt.Sprintf("King %s fights with %s", kingName, axe.Attack()) {
		t.Errorf("The King is not fighting in the right way.\n")
	}

	knife := Knife{Name: "Cuchillito"}
	king.SetWeapon(knife)
	if king.Fight() != fmt.Sprintf("King %s fights with %s", kingName, knife.Attack()) {
		t.Errorf("The King is not fighting in the right way.\n")
	}
}

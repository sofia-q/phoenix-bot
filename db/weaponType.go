package db

import (
	"errors"
)

type WeaponType int

const (
	SwordAndShield = iota
	DualBlades
	GreatSword
	LongSword
	Hammer
	HuntingHorn
	Lance
	GunLance
	SwitchAxe
	ChargeBlade
	InsectGlaive
	LightBowgun
	HeavyBowgun
	Bow
)

var weaponName = map[WeaponType]string{
	SwordAndShield: "Sword and Shield",
	DualBlades:     "Dual Blades",
	GreatSword:     "Greatsword",
	LongSword:      "Longsword",
	Hammer:         "Hammer",
	HuntingHorn:    "Hunting Horn",
	Lance:          "Lance",
	GunLance:       "Gunlance",
	SwitchAxe:      "Switch Axe",
	ChargeBlade:    "Charge Blade",
	InsectGlaive:   "Insect Glaive",
	LightBowgun:    "Light Bowgun",
	HeavyBowgun:    "Heavy Bowgun",
	Bow:            "Bow",
}

var weaponType = map[string]WeaponType{
	"Sword and Shield": SwordAndShield,
	"Dual Blades":      DualBlades,
	"Greatsword":       GreatSword,
	"Longsword":        LongSword,
	"Hammer":           Hammer,
	"Hunting Horn":     HuntingHorn,
	"Lance":            Lance,
	"Gunlance":         GunLance,
	"Switch Axe":       SwitchAxe,
	"Charge Blade":     ChargeBlade,
	"Insect Glaive":    InsectGlaive,
	"Light Bowgun":     LightBowgun,
	"Heavy Bowgun":     HeavyBowgun,
	"Bow":              Bow,
}

func (wt WeaponType) String() string {
	return weaponName[wt]
}

func (wt WeaponType) ParseStringToWeaponType(str string) (weapon WeaponType, err error) {
	if weapon, ok := weaponType[str]; ok {
		return weapon, nil
	}
	return 0, errors.New("invalid weapon type")
}

package model

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

func (wt WeaponType) String() string {
	return weaponName[wt]
}

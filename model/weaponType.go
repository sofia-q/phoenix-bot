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

var weaponHandle = map[WeaponType]string{
	SwordAndShield: "sns",
	DualBlades:     "db",
	GreatSword:     "gs",
	LongSword:      "ls",
	Hammer:         "ham",
	HuntingHorn:    "hh",
	Lance:          "lan",
	GunLance:       "gl",
	SwitchAxe:      "sa",
	ChargeBlade:    "cb",
	InsectGlaive:   "ig",
	LightBowgun:    "lbg",
	HeavyBowgun:    "hbg",
	Bow:            "bow",
}

func (wt WeaponType) String() string {
	return weaponName[wt]
}

func (wt WeaponType) GetWeaponHandle() string { return weaponHandle[wt] }

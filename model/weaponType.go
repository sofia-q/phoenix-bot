package model

type WeaponType string

const (
	SwordAndShield = "sns"
	DualBlades     = "db"
	GreatSword     = "gs"
	LongSword      = "ls"
	Hammer         = "ham"
	HuntingHorn    = "hh"
	Lance          = "lan"
	GunLance       = "gl"
	SwitchAxe      = "swaxe"
	ChargeBlade    = "cb"
	InsectGlaive   = "ig"
	LightBowgun    = "lbg"
	HeavyBowgun    = "hbg"
	Bow            = "bow"
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

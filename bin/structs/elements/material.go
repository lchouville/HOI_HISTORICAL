package hoiMod

import (
	s "hoiMod/bin/structs/systems"
)

type Material struct {
	// Input Value
	name string
	mohs float32
	// Output Value
	addStats s.Stats //Stat Additioned at Equipement stats
	MulStats s.Stats //Stat Multiplied at Equipement stats
}

func (material *Material) Calculation() {
	var armor_thinkness_M float32 //Store the multiplier Armor Stat
	armor_thinkness_M = 2400 / 6.75
	armor_thinkness_M = armor_thinkness_M * material.mohs
	armor_thinkness_M = 2400 / armor_thinkness_M
	armor_thinkness_M = 1 / armor_thinkness_M
	material.MulStats.SetArmorValue(armor_thinkness_M)
}
package hoiMod

import (
	s "hoiMod/bin/structs/systems"
)

/** Material Object -------------------------- **/

type Material struct {
	// Input Value
	name string
	mohs float32
	// Output Value
	addStats s.Stats //Stat Additioned at Equipement stats
	mulStats s.Stats //Stat Multiplied at Equipement stats
}


/** Builder --------------------------------- **/

func NewMaterial(_name string, _mohs float32)*Material{
	// create material with data gived on parameters
	var material = &Material{
		name: _name,
		mohs: _mohs,
	}
	// calculate material statistics
	material.calculation()
	return material
}


/** Accessors ------------------------------ **/

// Get the name of the material
func (material *Material) GetName() string{
	return material.name
}
// Set the Name to the material
func (material *Material) SetName(_name string){
	material.name = _name
}
// Get the mohs value of the material
func (material *Material) GetMohs() float32{
	return material.mohs
}
// Set the Moht value of the material
func (material *Material) SetMohs(_mohs float32){
	material.mohs = _mohs
}
// Get the adding effect with the material
func (material *Material) GetAddStats() s.Stats { 
	return material.addStats 
}
// Get the multiplier effect with the material
func (material *Material) GetMulStats() s.Stats { 
	return material.mulStats 
}


/** Method -------------------------------- **/

// Calculate The effect of the material on stats
func (material *Material) calculation() {
	var armor_thinkness_M float32 //Store the multiplier Armor Stat
	armor_thinkness_M = 2400 / 6.75
	armor_thinkness_M = armor_thinkness_M * material.mohs
	armor_thinkness_M = 2400 / armor_thinkness_M
	armor_thinkness_M = 1 / armor_thinkness_M
	material.mulStats.SetArmorValue(armor_thinkness_M)
}
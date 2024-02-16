package hoiMod

import s "hoiMod/bin/structs/systems"

/** Hull Object ------------------------------ **/

type Hull struct {
	// Input Value
	name string
	// size
	length int   // to mm
	width  int   // to mm
	hight  int   // to mm
	mass   int   // to kg
	armor  Armor // to mm
	// Output Value
	addStats s.Stats //Stat Additioned at Equipement stats
	mulStats s.Stats //Stat Multiplied at Equipement stats
}


/** Builder --------------------------------- **/

func NewHullE(name string) *Hull {
	var this Hull
	this.name = name
	return &this
}
func NewHullF(
	name string,
	length int, // to mm
	width int, // to mm
	hight int, // to mm
	mass int, // to kg
	armor Armor, // to mm
) *Hull {
	var this Hull
	this.name = name
	this.length = length
	this.width = width
	this.hight = hight
	this.mass = mass
	this.armor = armor
	calculate(&this)
	return &this
}


/** Accessors ------------------------------ **/

func (h *Hull) GetName() string {
	return h.name
}
func (h *Hull) GetSize() (lenght int, width int, hight int) {
	return h.length, h.width, h.hight
}
func (h *Hull) GetMass() int {
	return h.mass
}
func (h *Hull) GetArmor() Armor {
	return h.armor
}
// Get the adding effect with the material
func (hul *Hull) GetAddStats() s.Stats { 
	return hul.addStats 
}
// Get the multiplier effect with the material
func (hul *Hull) GetMulStats() s.Stats { 
	return hul.mulStats 
}

/** Method -------------------------------- **/

func calculate(h *Hull) {

}

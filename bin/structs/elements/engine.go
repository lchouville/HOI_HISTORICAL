package hoiMod

import (
	"fmt"
	s "hoiMod/bin/structs/systems"
)

/** Engine Object ---------------------------- **/

type Engine struct {
	// Input Value
	name string

	// Output Value
	addStats s.Stats //Stat Additioned at Equipement stats
	mulStats s.Stats //Stat Multiplied at Equipement stats
}

/** Builder ---------------------------------- **/

func NewEngine(_name string) *Engine {
	var eng Engine
	eng.name = _name

	return &eng
}

/** Accessors -------------------------------- **/

// Set the name of the engine
func (eng *Engine) SetName(_name string) { 
	eng.name = _name
}
// Get the name of the engine
func (eng *Engine) GetName() string { 
	if eng.name == "" {
		return "N/A"
	}
	return eng.name
}
// Get the adding effect with the material
func (eng *Engine) GetAddStats() s.Stats { 
	return eng.addStats 
}
// Get the multiplier effect with the material
func (eng *Engine) GetMulStats() s.Stats { 
	return eng.mulStats 
}

/** Method ----------------------------------- **/

// Print the engine
func (eng *Engine) Print() { 
	var text string
	// Create the text string
	text += "Name : " + eng.name + "\n"
	fmt.Println(text)
}

package hoiMod

import (
	"fmt"
)

type Engine struct {
	// Real Data
	name string

	// Game Data
	/* Misc Abilities */
	speed_modifier float32 // modifier for speed use for set the starting speed
	fuel_modifier  float32

	speed_multiplier float32 // multiplier for speed

	/* Defensive Abilities */
	/* Offensive Abilities */
	/* Space taken in convoy */

}

/*
Builder
*/
func NewEngineE(name string) *Engine {
	var this Engine
	this.name = name
	return &this
}
func NewEngineF(
	pName string,
	pSpeedModifier float32,
	pSpeedMultiplier float32,
	pFuelConsuming float32,
) *Engine {
	var this Engine
	this.name = pName
	this.speed_modifier = pSpeedModifier
	this.speed_multiplier = pSpeedMultiplier
	this.fuel_modifier = pFuelConsuming
	return &this
}

/*
Setter
*/
func (e *Engine) SetName(name string) { // set the name of the engine
	e.name = name
}

/*
Getter
*/
func (e *Engine) GetName() string { // get the name of the engine
	if e.name == "" {
		return "N/A"
	}
	return e.name
}

func (e *Engine) GetGameValues() {

}

/*
Print
*/
func (e *Engine) Print() { // Print the engine
	var text string
	// Create the text string
	text += "Name : " + e.name + "\n"
	fmt.Println(text)
}

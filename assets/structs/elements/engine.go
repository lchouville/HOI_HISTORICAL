package hoiMod

import "fmt"

type Engine struct {
	name string
	/* Misc Abilities */
	speed_modifier float32

	speed_multiplier float32

	fuel_consuming float32
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
	this.fuel_consuming = pFuelConsuming
	return &this
}

/*
Setter
*/
func (this Engine) SetName(pName string) { // set the name of the engine
	this.name = pName
}

/*
Getter
*/
func (this Engine) GetName() string { // get the name of the engine
	if this.name == "" {
		return "N/A"
	}
	return this.name
}

/*
Print
*/
func (this Engine) Print() { // Print the engine
	var text string
	// Create the text string
	text += "Name : " + this.name + "\n"
	fmt.Println(text)
}

package hoiMod

import (
	"fmt"
	d "hoiMod/bin/structs/display"
	e "hoiMod/bin/structs/elements"
	s "hoiMod/bin/structs/systems"
)

/** Tank Object ------------------------------ **/

type tank struct {
	// Real data
	name    string
	faction string
	date    string

	engine          e.Engine
	armor_thinkness int // the armor thickness applied to the Hull

	
	// Output Value
	stats s.Stats // Final stats of the tank

}


/** Builder --------------------------------- **/

func NewTankE(
	_name string,
	_faction string,
) *tank {
	var this tank
	this.name = _name
	this.faction = _faction
	return &this
}
func NewTankF(
	_name string,
	_faction string,
	_engine e.Engine,
) *tank {
	var this tank = *NewTankE(_name, _faction)
	this.engine = _engine
	return &this
}


/** Accessors ------------------------------ **/

func (tan *tank) SetName(_name string) {
	tan.name = _name
}
func (tan *tank) SetEngine(_engine e.Engine) {
	tan.engine = _engine
}

func (tan *tank) GetName() string {
	if tan.name == "" {
		return "N/A"
	}
	return tan.name
}
func (tan *tank) GetEngine() e.Engine {
	return tan.engine
}

// Get the Game statistics of the tank
func (tan *tank) GetStats() s.Stats { 
	return tan.stats 
}

/** Method -------------------------------- **/
func (tan *tank) calculation(){
	tan.stats.AddStats(tan.engine.GetAddStats())
	tan.stats.MulStats(tan.engine.GetAddStats())
}
func (tan *tank) Print() {
	var text string
	colorE := string(d.Color_Cyan)
	resetAll := string(d.ResetAll)
	// Create the text string
	text += (string(d.Text_Bold) +
		string(d.Color_Magenta) +
		tan.name +
		resetAll +
		"\n")
	// Engine
	text += (colorE + "↪ Engine : " + resetAll +
		"\t" + tan.engine.GetName() + "\n")
	fmt.Println(text)
}

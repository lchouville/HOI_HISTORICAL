package hoiMod

import (
	"fmt"
	d "hoiMod/assets/structs/display"
	e "hoiMod/assets/structs/elements"
)

type tank struct {
	// Real data
	name    string
	faction string
	date    string

	engine          e.Engine
	armor_thinkness int // the armor thickness applied to the Hull

	// Game data
	/* Misc Abilities */
	maximum_speed  float32 // modifier for downgrade the speed with the default mass of the hull
	reliability    float32
	fuel_consuming float32
	build_cost_ic  float32
	/* Defensive Abilities */
	defense      float32
	breakthrough float32
	hardness     float32
	armor_value  float32
	/* Offensive Abilities */
	soft_attack float32
	hard_attack float32
	ap_attack   float32
	air_attack  float32
	/* Space taken in convoy */
	lend_lease_cost float32
	// number of factories needed to build
	resources []Resource // resources needed to build

}

// Builder
func NewTankE(
	name string,
	faction string,
) *tank {
	var this tank
	this.name = name
	this.faction = faction
	return &this
}
func NewTankF(
	name string,
	faction string,
	engine e.Engine,
) *tank {
	var this tank = *NewTankE(name, faction)
	this.engine = engine
	return &this
}

// Setter
func (t *tank) SetName(name string) {
	t.name = name
}
func (t *tank) SetEngine(engine e.Engine) {
	t.engine = engine
}

// Getter
func (t *tank) GetName() string {
	if t.name == "" {
		return "N/A"
	}
	return t.name
}
func (t *tank) GetEngine() e.Engine {
	return t.engine
}

// Print
func (t *tank) Print() {
	var text string
	colorE := string(d.Color_Cyan)
	resetAll := string(d.ResetAll)
	// Create the text string
	text += (string(d.Text_Bold) +
		string(d.Color_Magenta) +
		t.name +
		resetAll +
		"\n")
	// Engine
	text += (colorE + "↪ Engine : " + resetAll +
		"\t" + t.engine.GetName() + "\n")
	fmt.Println(text)
}

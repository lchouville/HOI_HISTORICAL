package hoiMod

import (
	"fmt"

	d "hoiMod/assets/structs/display"
	e "hoiMod/assets/structs/elements"
)

type tank struct {
	name   string
	engine e.Engine
}

// Builder
func NewTankE(name string) *tank {
	var this tank
	this.name = name
	return &this
}
func NewTankF(name string, engine e.Engine) *tank {
	var this tank
	this.name = name
	this.engine = engine
	return &this
}

// Setter
func (this tank) SetName(name string) {
	this.name = name
}
func (this tank) SetEngine(engine e.Engine) {
	this.engine = engine
}

// Getter
func (this tank) GetName() string {
	if this.name == "" {
		return "N/A"
	}
	return this.name
}
func (this tank) GetEngine() e.Engine {
	return this.engine
}

// Print
func (this tank) Print() {
	var text string
	colorE := string(d.ColorCyan)
	resColor := string(d.ColorReset)
	// Create the text string
	text += "Name : " + this.name + "\n"

	// Engine
	text += (colorE + "↪ Engine : " + resColor +
		"\t" + this.engine.GetName() + "\n")
	fmt.Println(text)
}

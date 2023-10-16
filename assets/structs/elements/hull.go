package hoiMod

// s "hoiMod/assets/structs"

type Hull struct {
	// -- Real Data --
	name string
	// size
	length    int // to mm
	width     int // to mm
	hight     int // to mm
	mass      int // to kg
	thickness int // to mm
	// -- Game data --
	speed_modifier float32 // modifier with the hull mass at 1mm of armor
	// resources []s.Resource // resources needed to build
}

// Builder
// Getter
func (h *Hull) GetName() string {
	return h.name
}
func (h *Hull) GetSize() (lenght int, width int, hight int) {
	return h.length, h.width, h.hight
}
func (h *Hull) GetMass() int {
	return h.mass
}
func (h *Hull) GetThinckness() int {
	return h.thickness
}
func (h *Hull) GetGameValues() {

}

// Setter

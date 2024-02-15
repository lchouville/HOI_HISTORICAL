package hoiMod

type Hull struct {
	// -- Real Data --
	name string
	// size
	length int   // to mm
	width  int   // to mm
	hight  int   // to mm
	mass   int   // to kg
	armor  Armor // to mm
	// -- Game data --
	speed_modifier float32 // modifier with the hull mass at 1mm of armor
	armor_modifier int     // modifier with the hull
	// resources []s.Resource // resources needed to build
}

// Builder
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
func (h *Hull) GetArmor() Armor {
	return h.armor
}
func (h *Hull) GetArmorMod() int {
	return int(h.armor_modifier)
}

// Setter

// Methode

func calculate(h *Hull) {
	h.armor_modifier = ArmorCalculation(h.armor)

}

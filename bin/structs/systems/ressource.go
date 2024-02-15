package hoiMod

/** Resource Object **/
type Resource struct {
	name     string // Name of the resource
	quantity int    // Quantity of the resource
}

/** Builder **/

// NewResource creates a new Resource instance with provided parameters
func NewResource(_name string, _quantity int) *Resource {
	return &Resource{
		name:     _name,
		quantity: _quantity,
	}
}

/** Accessor  **/

// GetName gets the name of the resource
func (res *Resource) GetName() string {
	return res.name
}

// SetName sets the name of the resource
func (res *Resource) SetName(_name string) {
	res.name = _name
}

// GetQuantity gets the quantity of the resource
func (res *Resource) GetQuantity() int {
	return res.quantity
}

// SetQuantity sets the quantity of the resource
func (res *Resource) SetQuantity(_quantity int) {
	res.quantity = _quantity
}

// Actually in game
const (
	oil      = "Oil"
	steel    = "Steel"
	tungsten = "Tungsten"
	chromium = "Chromium"
)

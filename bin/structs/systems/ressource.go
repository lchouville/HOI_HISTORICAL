package hoiMod

/** Resource Object **/
type Resource struct {
	name     string // Name of the resource
	quantity float32    // Quantity of the resource
}

/** Builder **/

// Create a new Resource slice with a default capacity
func DefaultRessources(_value float32) []Resource {
	var ressources []Resource
	for _, r := range Resources {
		ressources = append(ressources, *NewResource(r,_value))
	}
	return ressources
}
// NewResource creates a new Resource instance with provided parameters
func NewResource(_name string, _quantity float32) *Resource {
	var resource = &Resource{
		name:     _name,
		quantity: _quantity,
	}
	Resources = append(Resources, resource.name)
	return resource
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
func (res *Resource) GetQuantity() float32 {
	return res.quantity
}

// SetQuantity sets the quantity of the resource
func (res *Resource) SetQuantity(_quantity float32) {
	res.quantity = _quantity
}



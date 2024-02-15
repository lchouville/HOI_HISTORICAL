package hoiMod

type Stats struct {
	/* Defensive Abilities */
	defense      float32 // defense capacity
	breakthrough float32 // breakthrough capacity
	hardness     float32 // capacity to resist attack (number attack can support)
	armor_value  float32 // armor value (represented the effective value)
	/* Offensive Abilities */
	soft_attack float32 // dommage make to an not armored enemy
	hard_attack float32 // dommage make to an armored enemy
	ap_attack   float32 // penetration of the equipement
	air_attack  float32 // dommage make to air enemy
	/* Misc Abilities */
	maximum_speed   float32 // maximum speed this equipement is allowed to reach
	reliability     float32 // reliability to the equipment
	fuel_consuming  float32 // Fuel consumption to the equipement
	build_cost_ic   float32 // number can make by 1 factory (depend of factory capacity)
	convert_cost_ic float32 // Factory needed to convert 1 equipment
	lend_lease_cost float32 // Place the equipement taked on convoy
	/* Material */
	resources_cost []Resource // Resources needed for producing the equipment
}

/** Stats Constructor **/
func NewStats(
	_defense float32,
	_breakthrough float32,
	_hardness float32,
	_armor_value float32,
	_soft_attack float32,
	_hard_attack float32,
	_ap_attack float32,
	_air_attack float32,
	_maximum_speed float32,
	_reliability float32,
	_fuel_consuming float32,
	_build_cost_ic float32,
	_convert_cost_ic float32,
	_lend_lease_cost float32,
	_resources []Resource,
) *Stats {
	stats := &Stats{
		defense:         _defense,
		breakthrough:    _breakthrough,
		hardness:        _hardness,
		armor_value:     _armor_value,
		soft_attack:     _soft_attack,
		hard_attack:     _hard_attack,
		ap_attack:       _ap_attack,
		air_attack:      _air_attack,
		maximum_speed:   _maximum_speed,
		reliability:     _reliability,
		fuel_consuming:  _fuel_consuming,
		build_cost_ic:   _build_cost_ic,
		convert_cost_ic: _convert_cost_ic,
		lend_lease_cost: _lend_lease_cost,
		resources_cost:  _resources,
	}
	return stats
}

/** Accessor **/
/* Defensive Abilities */

// SetDefense sets the defense stat
func (stats *Stats) SetDefense(_defense float32) {
	stats.defense = _defense
}

// GetDefense gets the defense stat
func (stats *Stats) GetDefense() float32 {
	return stats.defense
}

// SetBreakthrough sets the breakthrough stat
func (stats *Stats) SetBreakthrough(_breakthrough float32) {
	stats.breakthrough = _breakthrough
}

// GetBreakthrough gets the breakthrough stat
func (stats *Stats) GetBreakthrough() float32 {
	return stats.breakthrough
}

// SetHardness sets the hardness stat
func (stats *Stats) SetHardness(_hardness float32) {
	stats.hardness = _hardness
}

// GetHardness gets the hardness stat
func (stats *Stats) GetHardness() float32 {
	return stats.hardness
}

// SetArmorValue sets the armor value stat
func (stats *Stats) SetArmorValue(_armorValue float32) {
	stats.armor_value = _armorValue
}

// GetArmorValue gets the armor value stat
func (stats *Stats) GetArmorValue() float32 {
	return stats.armor_value
}

/* Offensive Abilities */
// SetSoftAttack sets the soft attack stat
func (stats *Stats) SetSoftAttack(_softAttack float32) {
	stats.soft_attack = _softAttack
}

// GetSoftAttack gets the soft attack stat
func (stats *Stats) GetSoftAttack() float32 {
	return stats.soft_attack
}

// SetHardAttack sets the hard attack stat
func (stats *Stats) SetHardAttack(_hardAttack float32) {
	stats.hard_attack = _hardAttack
}

// GetHardAttack gets the hard attack stat
func (stats *Stats) GetHardAttack() float32 {
	return stats.hard_attack
}

// SetApAttack sets the armor-piercing attack stat
func (stats *Stats) SetApAttack(_apAttack float32) {
	stats.ap_attack = _apAttack
}

// GetApAttack gets the armor-piercing attack stat
func (stats *Stats) GetApAttack() float32 {
	return stats.ap_attack
}

// SetAirAttack sets the air attack stat
func (stats *Stats) SetAirAttack(_airAttack float32) {
	stats.air_attack = _airAttack
}

// GetAirAttack gets the air attack stat
func (stats *Stats) GetAirAttack() float32 {
	return stats.air_attack
}

/* Misc Abilities */

// SetMaximumSpeed sets the maximum speed stat
func (stats *Stats) SetMaximumSpeed(_maximumSpeed float32) {
	stats.maximum_speed = _maximumSpeed
}

// GetMaximumSpeed gets the maximum speed stat
func (stats *Stats) GetMaximumSpeed() float32 {
	return stats.maximum_speed
}

// SetReliability sets the reliability stat
func (stats *Stats) SetReliability(_reliability float32) {
	stats.reliability = _reliability
}

// GetReliability gets the reliability stat
func (stats *Stats) GetReliability() float32 {
	return stats.reliability
}

// SetFuelConsuming sets the fuel consuming stat
func (stats *Stats) SetFuelConsuming(_fuelConsuming float32) {
	stats.fuel_consuming = _fuelConsuming
}

// GetFuelConsuming gets the fuel consuming stat
func (stats *Stats) GetFuelConsuming() float32 {
	return stats.fuel_consuming
}

// SetBuildCostIC sets the build cost in IC (Industrial Capacity) stat
func (stats *Stats) SetBuildCostIC(_buildCostIC float32) {
	stats.build_cost_ic = _buildCostIC
}

// GetBuildCostIC gets the build cost in IC (Industrial Capacity) stat
func (stats *Stats) GetBuildCostIC() float32 {
	return stats.build_cost_ic
}

// SetConvertCostIC sets the convert cost in IC (Industrial Capacity) stat
func (stats *Stats) SetConvertCostIC(_convertCostIC float32) {
	stats.convert_cost_ic = _convertCostIC
}

// GetConvertCostIC gets the convert cost in IC (Industrial Capacity) stat
func (stats *Stats) GetConvertCostIC() float32 {
	return stats.convert_cost_ic
}

// SetLendLeaseCost sets the lend lease cost stat
func (stats *Stats) SetLendLeaseCost(_lendLeaseCost float32) {
	stats.lend_lease_cost = _lendLeaseCost
}

// GetLendLeaseCost gets the lend lease cost stat
func (stats *Stats) GetLendLeaseCost() float32 {
	return stats.lend_lease_cost
}

// SetResources sets the resources cost
func (stats *Stats) SetResources(_resources []Resource) {
	stats.resources_cost = _resources
}

// GetResources gets the resources cost
func (stats *Stats) GetResources() []Resource {
	return stats.resources_cost
}
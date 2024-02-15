package hoiMod

type Stats struct {
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
	/* Misc Abilities */
	maximum_speed   float32
	reliability     float32
	fuel_consuming  float32
	build_cost_ic   float32
	convert_cost_ic float32
	lend_lease_cost float32
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
	}
	return stats
}

/** Accessor **/
/* Defensive Abilities */
func (stats *Stats) SetDefense(defense float32) {
	stats.defense = defense
}
func (stats *Stats) GetDefense() float32 {
	return stats.defense
}

func (stats *Stats) SetBreakthrough(breakthrough float32) {
	stats.breakthrough = breakthrough
}
func (stats *Stats) GetBreakthrough() float32 {
	return stats.breakthrough
}

func (stats *Stats) SetHardness(hardness float32) {
	stats.hardness = hardness
}
func (stats *Stats) GetHardness() float32 {
	return stats.hardness
}

func (stats *Stats) SetArmorValue(armorValue float32) {
	stats.armor_value = armorValue
}
func (stats *Stats) GetArmorValue() float32 {
	return stats.armor_value
}

// Offensive Abilities
func (stats *Stats) SetSoftAttack(softAttack float32) {
	stats.soft_attack = softAttack
}
func (stats *Stats) GetSoftAttack() float32 {
	return stats.soft_attack
}

func (stats *Stats) SetHardAttack(hardAttack float32) {
	stats.hard_attack = hardAttack
}
func (stats *Stats) GetHardAttack() float32 {
	return stats.hard_attack
}

func (stats *Stats) SetApAttack(apAttack float32) {
	stats.ap_attack = apAttack
}
func (stats *Stats) GetApAttack() float32 {
	return stats.ap_attack
}

func (stats *Stats) SetAirAttack(airAttack float32) {
	stats.air_attack = airAttack
}
func (stats *Stats) GetAirAttack() float32 {
	return stats.air_attack
}

// Misc Abilities
func (stats *Stats) SetMaximumSpeed(maximumSpeed float32) {
	stats.maximum_speed = maximumSpeed
}
func (stats *Stats) GetMaximumSpeed() float32 {
	return stats.maximum_speed
}

func (stats *Stats) SetReliability(reliability float32) {
	stats.reliability = reliability
}
func (stats *Stats) GetReliability() float32 {
	return stats.reliability
}

func (stats *Stats) SetFuelConsuming(fuelConsuming float32) {
	stats.fuel_consuming = fuelConsuming
}
func (stats *Stats) GetFuelConsuming() float32 {
	return stats.fuel_consuming
}

func (stats *Stats) SetBuildCostIC(buildCostIC float32) {
	stats.build_cost_ic = buildCostIC
}
func (stats *Stats) GetBuildCostIC() float32 {
	return stats.build_cost_ic
}

func (stats *Stats) SetConvertCostIC(convertCostIC float32) {
	stats.convert_cost_ic = convertCostIC
}
func (stats *Stats) GetConvertCostIC() float32 {
	return stats.convert_cost_ic
}

func (stats *Stats) SetLendLeaseCost(lendLeaseCost float32) {
	stats.lend_lease_cost = lendLeaseCost
}
func (stats *Stats) GetLendLeaseCost() float32 {
	return stats.lend_lease_cost
}

package hoiMod

import (
	"math"
)

type Armor struct {
	Front_thinckness int // front armor thickness
	Front_angle      int // front armor angle

	Right_thinckness int // right armor thickness
	Right_angle      int // right armor angle

	Left_thinckness int // left armor thickness
	Left_angle      int // left armor angle

	Rear_thinckness int // rear armor thickness
	Rear_angle      int // rear armor angle

	Top_thinckness int // top armor thickness
	Top_angle      int // top armor angle

	Bottom_thinckness int // bottom armor thickness
	Bottom_angle      int // bottom armor angle
}

func ArmorCalculation(armor Armor) int {
	var armor_value float64 //store the armor value calculated
	var slopTab []float64 = []float64{
		float64(armor.Front_angle),
		float64(armor.Right_angle),
		float64(armor.Left_angle),
		float64(armor.Rear_angle),
		float64(armor.Top_angle),
	}
	for i, v := range slopTab {
		tSlope := v * (math.Pi / 180)
		if i == 0 {
			armor_value += float64(1/math.Cos(float64(tSlope))) * 2
		} else {
			armor_value += float64(1 / math.Cos(float64(tSlope)))
		}
	}
	armor_value /= float64(len(slopTab))
	return int(math.Round(float64(armor_value)))
}
func ArmorMass(armor Armor, lenght int, width int, height int, mass int) (hullmass int) {
	hullmass = 0
	return
}

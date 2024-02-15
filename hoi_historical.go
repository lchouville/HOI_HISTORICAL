package main

import (
	"fmt"
	s "hoiMod/bin/structs"
	e "hoiMod/bin/structs/elements"
)

func main() {
	engine := e.NewEngineE("Moteur v12 Daimler-Benz MB509")
	tank := s.NewTankF("Maus", "GERMAN", *engine)
	tank.Print()

	var mausA e.Armor
	mausA.Front_thinckness = 220
	mausA.Front_angle = 35

	mausA.Right_thinckness = 185
	mausA.Right_angle = 0

	mausA.Left_thinckness = 185
	mausA.Left_angle = 0

	mausA.Rear_thinckness = 185
	mausA.Rear_angle = 30

	mausA.Top_thinckness = 105
	mausA.Top_angle = 9

	mausA.Bottom_thinckness = 105
	mausA.Bottom_angle = 0

	hull := e.NewHullF("maus", 9034, 3649, 2940, 128, mausA)
	fmt.Println(mausA.Front_thinckness, " : ", hull.GetArmorMod())

	tank2 := s.NewTankE("Panther", "GERMAN")
	tank2.Print()
	tank2.SetName("Panzer V")
	tank2.Print()
	// tank.GetEngine().Print()
}

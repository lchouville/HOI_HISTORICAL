package main

import (
	s "hoiMod/assets/structs"
	e "hoiMod/assets/structs/elements"
)

func main() {
	engine := e.NewEngineE("Moteur v12 Daimler-Benz MB509")
	tank := s.NewTankF("Maus", "GERMAN", *engine)
	tank.Print()

	tank2 := s.NewTankE("Panther", "GERMAN")
	tank2.Print()
	tank2.SetName("Panzer V")
	tank2.Print()
	// tank.GetEngine().Print()
}

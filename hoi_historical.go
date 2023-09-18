package main

import (
	s "hoiMod/assets/structs"
	e "hoiMod/assets/structs/elements"
)

func main() {
	engine := e.NewEngineE("Moteur v12 Daimler-Benz MB509")
	tank := s.NewTankF("Maus", *engine)
	tank.Print()

	tank2 := s.NewTankE("Panther")
	tank2.Print()
	// tank.GetEngine().Print()
}

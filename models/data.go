package models

import "log"

var DB = make([]Equation, 3)

type Equation struct {
	A      int `json:"A"`
	B      int `json:"B"`
	C      int `json:"C"`
	Nroots int `json:"Nroots"`
	solved bool
}

func AddEquation(eq Equation) {
	DB = append(DB, eq)
}

func solve(eq Equation) Equation {
	switch {
	case eq.solved:
		return eq // no calculation needed
	case eq.A == 0 && eq.B == 0: // y = const
		eq.Nroots = 0 // since it's guaranteed that at least one arg is non-zero
	case eq.A == 0: // linear and non-const
		eq.Nroots = 1
	default: // square
		{
			Dsqr := eq.B*eq.B - 4*eq.A*eq.C
			switch {
			case Dsqr > 0:
				eq.Nroots = 2
			case Dsqr == 0:
				eq.Nroots = 1
			default:
				eq.Nroots = 0
			}
		}
	}
	eq.solved = true
	return eq
}

func GetEquation() Equation {
	if len(DB) == 0 {
		log.Fatal("Broken promise: solve before grab")
	}
	return  solve(DB[len(DB)-1])
}
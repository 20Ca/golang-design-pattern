package main

import (
	"fmt"
	"math"
)

// An implementation of a finite state machine in Go,
// inspired by David Mertz's article "Charming Python: Using state machines"
// (http://www.ibm.com/developerworks/library/l-python-state/index.html)

type Handler func(interface{}) (string, interface{})

type Machine struct {
	Handlers   map[string]Handler
	StartState string
	EndStates  map[string]bool
}

func (machine *Machine) AddState(handlerName string, handlerFn Handler) {
	machine.Handlers[handlerName] = handlerFn
}

func (machine *Machine) AddEndState(endState string) {
	machine.EndStates[endState] = true
}

func (machine *Machine) Execute(cargo interface{}) {
	if handler, present := machine.Handlers[machine.StartState]; present {
		for {
			nextState, nextCargo := handler(cargo)
			_, finished := machine.EndStates[nextState]
			if finished {
				break
			} else {
				handler, present = machine.Handlers[nextState]
				cargo = nextCargo
			}
		}
	}
}
func do_math(i float64) float64 {
	return math.Abs(math.Sin(i) * 31.0)
}

func ones_counter() Handler {
	return func(val interface{}) (nextState string, nextVal interface{}) {
		nextState = ""
		nextVal = val.(float64)

		fmt.Printf("1s State:\t")
		for {
			switch {
			case (nextVal.(float64) <= 0 || nextVal.(float64) >= 30):
				nextState = "outofrange"
			case (nextVal.(float64) >= 20 && nextVal.(float64) < 30):
				nextState = "twenties"
			case (nextVal.(float64) >= 10 && nextVal.(float64) < 20):
				nextState = "tens"
			default:
				fmt.Printf(" @ %2.1f+", nextVal.(float64))
			}

			if len(nextState) > 0 {
				break
			}
			nextVal = do_math(nextVal.(float64))
		}
		fmt.Printf(" >>\n")
		return
	}
}

func tens_counter() Handler {
	return func(val interface{}) (nextState string, nextVal interface{}) {
		nextState = ""
		nextVal = val.(float64)

		fmt.Printf("10s State:\t")
		for {
			switch {
			case (nextVal.(float64) <= 0 || nextVal.(float64) >= 30):
				nextState = "outofrange"
			case (nextVal.(float64) >= 20 && nextVal.(float64) < 30):
				nextState = "twenties"
			case (nextVal.(float64) >= 1 && nextVal.(float64) < 10):
				nextState = "ones"
			default:
				fmt.Printf(" #%2.1f+", nextVal)
			}

			if len(nextState) > 0 {
				break
			}
			nextVal = do_math(nextVal.(float64))
		}
		fmt.Printf(" >>\n")
		return
	}
}

func twenties_counter() Handler {
	return func(val interface{}) (nextState string, nextVal interface{}) {
		nextState = ""
		nextVal = val.(float64)

		fmt.Printf("20s State:\t")
		for {
			switch {
			case (nextVal.(float64) <= 0 || nextVal.(float64) >= 30):
				nextState = "outofrange"
			case (nextVal.(float64) >= 10 && nextVal.(float64) < 20):
				nextState = "tens"
			case (nextVal.(float64) >= 1 && nextVal.(float64) < 10):
				nextState = "ones"
			default:
				fmt.Printf(" *%2.1f+", nextVal)
			}

			if len(nextState) > 0 {
				break
			}
			nextVal = do_math(nextVal.(float64))
		}
		fmt.Printf(" >>\n")
		return
	}
}

func main() {
	m := Machine{map[string]Handler{}, "ones", map[string]bool{}}
	m.AddState("ones", ones_counter())
	m.AddState("tens", tens_counter())
	m.AddState("twenties", twenties_counter())
	m.AddEndState("outofrange")

	m.Execute(1.0)
}

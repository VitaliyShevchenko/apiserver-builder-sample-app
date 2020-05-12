package controller

import (
	"gopath/src/pet-project/pkg/controller/actor"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, actor.Add)
}

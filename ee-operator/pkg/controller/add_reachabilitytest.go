package controller

import (
	"github.com/tanalam2411/eagle-eye/pkg/controller/reachabilitytest"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, reachabilitytest.Add)
}

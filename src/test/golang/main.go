package main

import (
	"os"

	"github.com/starter-go/rbac/modules/librbac"
	"github.com/starter-go/units"
)

func main() {

	a := os.Args
	m := librbac.ModuleForTest()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}

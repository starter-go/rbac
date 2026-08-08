package main

import (
	"os"

	"github.com/starter-go/rbac/modules/rbac"
	"github.com/starter-go/units"
)

func main() {

	a := os.Args
	m := rbac.ModuleForTest()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}

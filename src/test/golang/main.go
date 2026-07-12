package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/units"
)

func main() {

	a := os.Args
	m := starter.Module()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}

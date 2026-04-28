package main

import (
	"os"

	"github.com/starter-go/starter"
)

func main() {

	a := os.Args
	m := starter.Module()
	i := starter.Init(a)

	i.MainModule(m)
	i.WithPanic(true).Run()
}

package main

import (
	"fmt"
	"github.com/crit/critical-go/input"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()

	m.Use(input.MartiniMiddleware())

	m.Get("/", func(in input.Input) string {
		return fmt.Sprintf("We got: %+v", in.All())
	})

	m.Run()
}

package main

import (
	"fmt"
	"molcular-dynamics/domain"
)

func main() {
	md := domain.NewMdSystem()
	md.Run()
	fmt.Sprintln("main")
}

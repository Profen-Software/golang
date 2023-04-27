package main

import (
	"fmt"
	"main/conditions"
	"main/interfaces"
)

func main() {
	fmt.Println("Hello World")

	conditions.PrintConditions()

	interfaces.PrintInterface()
}

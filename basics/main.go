package main

import (
	"fmt"
	"main/conditions"
	"main/interfaces"
	"main/structs"
)

func main() {
	fmt.Println("Hello World")

	conditions.PrintConditions()

	interfaces.PrintInterface()

	structs.PrintStructs()
}

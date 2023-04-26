package conditions

func PrintConditions() {

	x := 10
	y := 20

	println("***  conditions ***")
	// if statement
	if x < y {
		println("x is less than y")
	} else if x > y {
		println("x is greater than y")
	} else {
		println("x is equal to y")
	}

	println("----------")

	// switch statement
	switch x {
	case 1:
		println("x is 1")
	case 2:
		println("x is 2")
	case 3:
		println("x is 3")
	default:
		println("x is not 1, 2 or 3")
	}

	println("***  end conditions ***")

}

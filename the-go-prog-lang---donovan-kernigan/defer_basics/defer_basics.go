package deferbasics

func Double(x int, y int) int {
	var result int

	defer func() { result *= result }()

	result = x * y

	return result
}

func DoubleBareReturn(x int, y int) (result int) { // named return catched by defer

	defer func() { result *= result }()

	return x * y
}

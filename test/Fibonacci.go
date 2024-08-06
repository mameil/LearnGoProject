package test

func fibonacciRecursive(idx int) int {
	if idx < 0 {
		return 0
	}

	if idx < 2 {
		return idx
	}

	return fibonacciRecursive(idx-1) + fibonacciRecursive(idx-2)
}

func fibonacciIterate(idx int) int {
	if idx < 0 {
		return 0
	}

	if idx < 2 {
		return idx
	}

	a, b := 0, 1
	for i := 2; i <= idx; i++ {
		a, b = b, a+b
	}
	return b
}

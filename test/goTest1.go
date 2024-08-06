package test

import "fmt"

func multiply(num int) int {
	return num * num
	//return 9 * 9
}

func main() {
	myNum := 9

	fmt.Printf("%d * %d = %d", myNum, myNum, multiply(myNum))
}

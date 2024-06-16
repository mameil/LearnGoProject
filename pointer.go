package main

import "fmt"

func main() {
	//포인터란.. 메모리 주소를 값으로 가지는 타입
	//변수를 선언하게되면 메모리 상의 어딘가에 저장이 되고 저장된 그 위치를 의마한다
	//포인터를 이용하면 하나의 변수를 다양한 포인터를 통해서 호출하는 것이 가능
	var myPointer *int
	myOriginal := 1234

	myPointer = &myOriginal
	fmt.Println("myOriginal 이라는 변수의 메모리 주소 값 : ", myPointer)
	fmt.Println("myPointer 가 가리키는 변수의 값 : ", *myPointer)
	fmt.Println("myPointer 가 가리키는 변수인 myOriginal 의 값 : ", myOriginal)

	//포인터의 값 비교를 해보자
	var a int = 20
	var b int = 30

	var p1 *int = &a
	var p11 *int = &a
	var p2 *int = &b

	fmt.Println(p1 == p11)
	fmt.Println(p11 == p2)

	//포인터의 기본 초기화 값은 nil 이다 >> 아무런 주소를 가리키지 않는 상황이라는 의미
	var p *int
	fmt.Println(p) //<nil>
}

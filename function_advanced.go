package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		함수의 고급 사용법
		1. 가변 인수 함수 >> 함수를 생성할 때 인수의 갯수를 고정하지 않고 받을 수 있음
		2. defer 지연 실행 >> defer 키워드를 사용하면 함수 종료 전에 반드시 실행해야 하는 코드를 실행할 수 있음
		3. 함수 타입 변수 >> 함수를 값으로 가지는 타입
		4. 함수 리터럴 >> 익명함수를 정의하고, 함수 타입 변수에 대입할 수 있음
	*/

	//1. 가변 인수 함수
	//가변 인수 함수는 말그대로 인자로 들어오는 값의 갯수를 가변적으로 설정해둔걸 의미
	//가장 큰 예시로는 fmt.Println 이 존재
	fmt.Println()
	fmt.Println(1, 2)
	fmt.Println(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	//가변 인수 함수를 직접 만들어보자(sum 함수 참고)
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) //55

	//가변 인수 함수를 직접 만들어보자(Print 함수 참고)
	Print(1, 2, 3, "a", "b", "c")

	//2. defer 지연 실행
	//코드를 짜는데 있어서 함수가 종료되기 직전에 실행해야 하는 코드가 있을 수 있음(예로는 file IO, DB 접속 종료 등)
	//이런 경우에 "defer 명령문" 이렇게 처리해주면, 함수가 다 수행되고 종료 전에 명령문을 수행해준다, 그리고 이걸 지연 실행이라고 명명한다
	f, err := os.Create("test.txt") // 파일을 만들고
	if err != nil {
		fmt.Println("failed to create test.txt")
		return
	}
	defer fmt.Println("종료되기 전에 defer 키워드가 붙은걸 확인해야함")
	defer f.Close()
	defer fmt.Println("성공적으로 파일을 닫았습니다")
	fmt.Println("파일에 글 쓰는중")
	fmt.Fprintln(f, "Hello, World!") //글을 쓴다

	//3. 함수 타입 변수
	//함수 역시도 포인터를 가지고 특정 메모리에 저장되어있다
	//함수 타입 변수 관련 예시를 보자(getOperator 참고)
	var operator func(int, int) int
	var operator2 func(int, int) int

	operator = getOperator("+")
	operator2 = getOperator("*")

	fmt.Println(operator(3, 4))  //7
	fmt.Println(operator2(3, 4)) //12

}

// 가변 인수 함수를 선언하는 방법은 인자를 설정하는 곳에서 ... 을 통해서 선언해두면 알아서 가변 인수 함수로 인식
func sum(nums ...int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}
	return sum
}

// 근데 가변 인수 함수를 만들 때 인자를 꼭 같은 타입으로만 설정해야 하는건 꼭 아님
// 아래 처럼 interface 를 선언해주면된다 >> 모든 타입은 "빈 인터페이스"를 포함하고 있기 때문에 인터페이스로 받고, 타입캐스팅을 통해서 각각에 맞춰서 처리해주면된다
func Print(args ...interface{}) {
	for _, arg := range args {
		switch f := arg.(type) {
		case int:
			fmt.Printf("%d\n", f)
		case string:
			fmt.Printf("%s\n", f)
		}
	}
}

// 함수 타입 변수를 활용한 예제
func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

// func(int, int) int 이렇게 함수가 기니까
// type opFunc func (int, int) int << 이렇게 별칭을 선언해두고
// func getOperator(op string) opFunc { << 이렇게 사용만해줘도 된다
func getOperator(op string) func(int, int) int {
	//op 라는 인자값에 따른 함수 타입 반환한다
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

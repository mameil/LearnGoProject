package main

import (
	"fmt"
)

func main() {
	//hello()
	//variable1()
	//variable2()
	//printTest()
	input()
}

func hello() {
	//기본적인 print 문은 fmt 패키지를 통해서 수행된다
	fmt.Println("=========================================")
	fmt.Println("===============Hello World===============")
	fmt.Println("=========================================")
}

func variable1() {
	//변수 놀이 1
	fmt.Println("=========================================")
	var first int = 10
	var second int = 20

	var sum = first + second
	sum += 10
	fmt.Println(sum)
	fmt.Println("=========================================")
}

func variable2() {
	//변수 놀이 2
	fmt.Println("=========================================")
	var a int = 3 //기본적인 선언(var + 변수명 + 변수의 타입 + 초기값)
	var b int     //(var + 변수명 + 변수의 타입 + (초기값이 없으면 타입별 기본값으로 대체)) >> int 의 초기값은 0 이구나
	var c = 4     //(var + 변수명 + (변수의 타입은 초기값을 기반으로 타입이 지정) + 초기값)
	d := 3        //:= 을 사용하면 var 없이 초기값만을 통해서 해당 변수의 타입, 초기값을 지정할 수 있음
	fmt.Println(a + b + c + d)
	fmt.Println("=========================================")

	/**
	정수타입의 초기값은 0 + :=(타입미지정)을 통하면 int타입으로 자동 지정
	실수타입의 초기값은 0.0 + :=(타입미지정)을 통하면 float64 타입으로 자동 지정
	boolean 의 초기값은 false
	string 의 초기값은 ""(공백) + :=(타입미지정)을 통하면 string 타입으로 자동 지정
	*/
	q := 3.14
	w := 3
	e := "hello world"
	fmt.Println(q, w, e)
}

func printTest() {
	myInt := 123
	myString := "my world"
	myFloat := 3.141592

	/**
	print > 그냥 입력값 출력
	println > 그냥 입력값 출력 + \n
	printf > 서식에 맞도록 입력값 출력 + \n
	*/
	//%v 가 각 변수에 맞는 타입에 맞춰서 자동으로 출력되는 방식(가장 범용성 높은듯)
	fmt.Printf("myInt: %v \nmyString: %v \nmyFloat: %v \n", myInt, myString, myFloat)
	fmt.Println("\n===================================================")
	//%d 가 정수, %s 가 스트링, %f 가 소수
	fmt.Printf("myInt: %d \nmyString: %s \nmyFloat: %f", myInt, myString, myFloat)
}

func input() {
	fmt.Println("==============================================")
	var a int
	var b int
	/**
	scan 함수를 사용해서 입력값을 전달받을 수 있음
	*/

	//scan 는 변수의 주소가 들어간다
	//하나 이상의 변수를 argument 로 설정해두면 입력 시 각 변수의 구분은 "공백"으로 구분해서 입력한다
	fmt.Scan(&a, &b)
	fmt.Println(a + b)

	fmt.Println("==============================================")
	//scanf 는 입력받을 format + 변수의 주소
	//scan 과는 다르게 입력받을 format 대로 입력해줘야 한다 > 아래를 예시로 들면 첫번쨰 입력한 정수가 a에 할당, \n 이 있으니 엔터치고, 그리고 두번쨰 입력한 정수가 b에 할당
	fmt.Scanf("%v\n%v", &a, &b)
	fmt.Println(a + b)

	fmt.Println("==============================================")
	//scanln 은 scan 과 같이 입력받고 무조건 엔터가 있어야한다
	//아래의 scan 을 보면 무조건 2개의 변수가 다 나올때까지 기다린다
	//아래를 실행하고 1을 입력하면 a 에 값이 할당되고 > 스페이스바를 누르든, 엔터를 누르든 뭔짓을 하든 무조건 2번째 변수값의 입력을 기다리는 개념
	fmt.Scan(&a, &b)
	fmt.Println(a + b)

	//아래의 scanln 을 보면 위의 scan 때와는 다르게 엔터가 들어오면 그냥 바로 종료해버리는 개념이다
	var c int
	//아래를 실행하고 1을 입력 후 엔터를 쳐보면 scanln 이 종료되고 입력받은 1은 a 에 할당되고 c는 할당받지 못한 상태로 그냥 기본값인 0 이 더해진다
	fmt.Scanln(&a, &c)
	fmt.Println(a + c)
}

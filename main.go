package main

import (
	"fmt"
)

func main() {
	//hello()
	//variable1()
	//variable2()
	//printTest()
	//input()
	//checkMethod()
	//multiReturn()
	//constTest()
	//ifTest()
	switchTest()
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
	var isAble bool = true
	fmt.Println(a + b + c + d)
	fmt.Println(isAble)
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

func checkMethod() {
	a := 3
	b := 4
	fmt.Println("=============================================")
	fmt.Println(myAdd(myAdd(a, b), 3))
	fmt.Println("=============================================")
}

func myAdd(a int, b int) int {
	return a + b
}

func multiReturn() {
	//go 에서는 함수에서 여러개의 값을 리턴하는 것이 가능하다..!
	sum, zeroAble := multiFun(1, -1)
	fmt.Println(sum, zeroAble)
}

// 리턴값이 여러개일 때는 리턴타입들을 소괄호로 묶어서 표현한다 @see multiFun
// 리턴하는 값에 변수명을 설정하는 것이 가능하다
// 리턴하는 값에 변수명을 붙힐 때는 붙힐꺼면 다 붙히고 안붙힐꺼면 다 안붙히고 둘중 하나만 해야함
func multiFun(a int, b int) (sr int, sb bool) {
	sr = a + b
	if a+b == 0 {
		sb = true
	} else {
		sb = false
	}
	return
}

func constTest() {
	//상수라는 개념이 존재한다
	const a = 1
	//a = 2 < compile error
	const b int = 2
	//b = 3 < compile error
	fmt.Println(a + b)

	//상수는 한꺼번에 선언이 가능하다(var 는 불가)
	//+iota 라는 키워드를 사용해서 괄호 범위 내에서 자동으로 증가하는 키워드를 사용할 수 있음(var는 불가)
	//+const() 을 사용해서 여러개의 상수를 선언할 때 첫번쨰 인자에만 변수를 적어주고 초기값이 동일하게
	//상수 같은 경우에도 var 에서 선언한 것 처럼 타입이 필수는 아님
	const (
		q = iota + 1
		w
		e
		r
	)
	fmt.Printf("q: %v \n", q)
	fmt.Printf("w: %v \n", w)
	fmt.Printf("e: %v \n", e)
	fmt.Printf("r: %v \n", r)
}

func ifTest() {
	//go 에서 if 문을 사용하는데 있어서 특이한점은 if 내에 들어가는 조건문이 간단하면 () 괄호 없이 if ~ 이렇게 만 작성하고도 수행이 가능하다
	if 2 > 1 {
		fmt.Println("IF 문 정상 수행")
	} else {
		fmt.Println("ELSE 문 정상 수행")
	}

	//go 에서 if 문의 특이한점은 if 문에 초기값을 넣어줄 수 있다는 점 >> 이해가 잘 안갔는데
	//간단하게 설명해보면 if 문을 특정 함수의 수행 여부에 따라서 수행되게 할 수 있다는 점
	if ipt, isPlus := checkPlus(9999); ipt == 9999 && isPlus {
		fmt.Println("ipt 값이 양수이고 그 값은 ", ipt, "이다")
	} else {
		fmt.Println("이상한 값")
	}

	if ipt, isPlus := checkPlus(-1234); isPlus {
		fmt.Println("ipt 값이 양수이고 그 값은 ", ipt, "이다")
	} else {
		fmt.Println("ipt 값이 음수이고 그 값은 ", ipt, "이다")
	}
}

func checkPlus(ipt int) (int, bool) {
	if ipt >= 0 {
		return ipt, true
	} else {
		return ipt, false
	}
}

func switchTest() {
	//go 에서 switch 문을 사용하는데 있어서 특이한 점은 if 문과 동일하게 조건문이 간단하면 () 괄호 없이 switch ~ 이렇게 만 작성하고도 수행이 가능하다
	//추가적으로 switch 에서도 초기값을 넣어주는 것이 가능하다
	//;(세미콜론) 전에 초기값을 설정하는 로직을 쭉 작성해주고 ; 으로 초기값 설정 로직 종료시키고
	//뒤에 변수를 입력하면 switch(변수)~~ 이렇게 변수에 대한 switch 가 수행되는거고
	//뒤에 변수를 입력하지 않으면 조건문으로 수행되는 switch 문이 수행된다고 보면 된다
	//추가로 switch 문을 사용하는데 있어서 break 을 걸지 않아도 자동으로 case 하나 수행하면 빠져나온다 > 자바에서는 케이스에 걸렸을 때 break 가 있으면 내려가면서 다 수행하는디
	switch age := getAge(); {
	case age >= 20:
		fmt.Println("사용자의 나이는 20대 입니다")
	case age >= 10:
		fmt.Println("사용자의 나이는 10대인데 switch 문 특성상 여기까진 않올 듯")
	default:
		{
			fmt.Println("여기까지 도달하다니.. 대단하군")
		}
	}
}

func getAge() int {
	return 28
}

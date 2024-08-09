package main

import (
	"bufio"
	"fmt"
	"os"
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
	//switchTest()
	//forTest()
	//forTest2()
	//forTest3()
	arrayTest()
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
	//변수 놀이 3
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
	//a = 3 < compile error
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

func forTest() {
	//for 문의 형식은 자바랑 많이 유사항 형태이다
	//다른 점은 역시.. for(int i=0; i<10; i++) { 이렇게 자바에서 하던걸 for i:=0; i<10; i++ { 이렇게 사용한다는 점이 다른점이다
	//약간 중괄호 혐이 있는 듯한 느낌이다
	//기본적인 수행
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	/**
	초기값, 조건문, 후처리 순서대로 각각을 빼고 수행하는 것도 가능하다
	초기문을 생략하는 경우
	후처리를 생략하는 경우
	조건문만 있는 경우
	이렇게 경우의 수가 몇가지가 있는데 각각의 경우는 필요시에 사용하면 될 것 같구 사용빈도가 굉장히 낮을 것으로 예상
	*/
}

func forTest2() {
	/*
		이외에도 다른 프로그램 언어와 동잃하게
			continue > 바로 다음 단계의 루프로 진행
			break > 진행하고 있는 루프를 종료
			이러한 키워드들이 존재
	*/

	//짝수를 받을때 까지 계속해서 입력을 받고,
	//숫자가 아닌 입력이 들어오면 "숫자를 입력하세요" 라는 문구가 나오고
	//홀수인 입력이 들어오면 "$입력값 은 홀수입니다" 라는 문구를 보여주고
	//짝수인 입력이 들어오면 "$입력값 은 짝수입니다" 라는 문구를 보여주고 루프를 종료하는 함수

	stdin := bufio.NewReader(os.Stdin)
	for { //for 문에서 따로 초기값, 조건값, 증감값 선언을 안해주면 기본적으로 무한루프로 판단
		fmt.Println("값을 입력해주세요")
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("숫자를 입력하셔야 합니다")
			//scan 과 같은 입력함수는 Buffer(입력 데이터를 임시로 저장하는 메모리 영역)에 저장되고 이걸 프로그램이 읽는 방식
			//현재 코드를 기준으로 scan 은 Buffer 에 올라가있는 값에서 숫자만 뽑아서 읽도록 되어있는데
			//숫자 + 엔터 를 통해서 입력을 했기 떄문에 버퍼에는 숫자 + 엔터로 구성되어있고 여기서 숫자만 뽑고 엔터가 남아있는 상황임
			//그래서 ReadString('\n') 을 해주지 않으면(아래에 실제로 테스트해봄) 자동으로 공백을 읽어가면서 수행된다
			stdin.ReadString('\n')
			continue
		}

		if num%2 == 0 {
			fmt.Println("입력하신 숫자", num, "은 짝수입니다")
			break
		} else {
			fmt.Println("입력하신 숫자", num, "은 홀수입니다")
		}
	}

	//버퍼에서 공백을 빼주는 작업이 없는 경우
	//for {
	//	fmt.Println("값을 입력해주세요")
	//	var num int
	//	_, err := fmt.Scan(&num)
	//	if err != nil {
	//		fmt.Println("숫자를 입력해주세요")
	//		continue
	//	}
	//
	//	if num%3 == 0 {
	//		fmt.Println("입력하신 숫자", num, "은 짝수입니다")
	//		break
	//	} else {
	//		fmt.Println("입력하신 숫자", num, "은 홀수입니다")
	//	}
	//}

	//값을 입력해주세요
	//vv
	//숫자를 입력해주세요
	//값을 입력해주세요
	//숫자를 입력해주세요
	//값을 입력해주세요

	//값을 vv+엔터 이렇게 입력하면
	//프로그램에서는 vv 을 도출해서 수행이되면서 "숫자를 입력해주세요" 가 나올꺼고
	//continue 를 통해서 다시 "값을 입력해주세요" 가 나올꺼고
	//버퍼에서 남은 공백을 자동으로 추출해서 프로그램이 수행되니까 "숫자를 입력해주세요" 가 나올꺼고
	//continue 를 통해서 다시 "값을 입력해주세요" 가 나올꺼고
}

func forTest3() {
	//이중 루프를 사용해서 별을 찍어본다
	//별을 기준으로 피라미드를 세워본다 입력값을 높이를 기준으로 수행됨 + 입력값은 숫자가 이쁠듯
	for {
		var num int
		fmt.Println("그리고 싶은 피라미드의 너비 사이즈를 입력하세요\n종료하시고 싶으시면 0 을 입력해주세요")
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("숫자를 입력해주셔야 합니다")
			continue
		} else if num == 0 {
			fmt.Println("루프를 종료합니다...")
			break
		} else if num%2 == 0 {
			fmt.Println("숫자는 짝수를 입력해주셔야 합니다")
			continue
		}

		for i := 1; i <= num; i += 2 {
			for j := 0; j < (num-i)/2; j++ {
				fmt.Print(" ")
			}
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
			for j := 0; j < (num-i)/2; j++ {
				fmt.Print(" ")
			}
			fmt.Println()
		}
	}
}

func arrayTest() {
	//배열의 생김새가 이상하다
	//[사이즈]타입{} 형식으로 배열을 초기화 시켜줄 수 있음
	var _ [3]int
	_ = [3]int{}

	//배열을 선언하면서 배열의 구성도 한꺼번에 구성할 수 있음
	//len() 함수를 사용해서 배열의 사이즈를 구하는 것이 가능
	var myArray = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	//이렇게 특정 index 부분에만 값을 넣고 나머지는 default 값으로 설정되게끔도 가능
	initArray := [3]string{0: "a", 2: "c"}
	fmt.Println(initArray[0]) //a
	fmt.Println(initArray[1]) //공백
	fmt.Println(initArray[2]) //c

	//배열의 사이즈를 정하지 않고 배열을 선언하는 것도 가능
	//but 초기화 시 배열을 구성한 값들을 기반으로 배열의 사이즈를 측정하기 떄문에 참고
	sizeArray := [...]int{0: 1, 1: 2, 3: 4}
	fmt.Println(sizeArray)      //배열의 print 가 이쁘게 잘나오네...
	fmt.Println(len(sizeArray)) //4

	//배열의 사이즈를 선언하는데 있어서 숫자를 변수로 설정할 수도 있지만 >>> 상수로 설정된 변수를 사이즈에 넣을 수 있음
	const TEST_SIZE = 3
	constArray := [TEST_SIZE]int{1, 2, 3}
	fmt.Println(constArray)

	//배열을 순회하는 방법으로 range 키워드를 사용할 수 있음
	//range 을 사용하게 되면 i(인덱스값), v(인덱스에 들어있는 값) 이렇게 도출된다
	for i, v := range myArray {
		fmt.Println(i, ":", v)
	}

	//만약에 인덱스 값이 필요없거나 인덱스에 맞는 value 값이 필요없으면 _ 을 통해서 무시해주는 것도 가능
	for _, v := range myArray {
		fmt.Println(v, ":", v)
	}

	multiArray := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(multiArray)
	//배열의 메모리 사이즈는 배열의 크기 * 배열의 구성 타입 사이즈
	//ex) multiArray > 3 * 3 * 8(int사이즈)

}

package main

import (
	"fmt"
)

func main() {
	//문자열은 큰 따옴표나 백쿼트(₩)로 묶어서 표시하는데 이 각각은 다른 쓰임새가 있다
	//₩ 으로 문자열을 묶으면 문자열 안의 특수문자가 일반 문자처럼 진행됨
	var str1 string = "Hello\tWorld!"
	fmt.Println(str1) //Hello   World!
	str2 := `Hello\tWorld!`
	fmt.Println(str2) //Hello\tWorld!
	//쓸일이 얼마나 있을지는 감이 안오지만.. 아무튼 그렇다는거

	//Go 언어는 UTF-8 문자 코드를 표준 문자코드로 사용한다 >> 영어, 숫자가 1바이트고 한글이 3~3 바이트

	//문자 하나를 표현할 때는 rune 타입을 사용해서 표현한다
	//UTF-8 을 기준으로 한 글자가 1~3 바이트의 사이즈를 차지하고 있는데 go 에서는 3바이트를 기준으로 타입을 가진게 없기 때문에 rune 타입을 사용해서 표현
	var testChar rune = '뀨'
	fmt.Printf("%c\n", testChar) //뀨

	//문자열의 길이를 len() 함수를 사용해서 측정할 수 있다
	str3 := "kdshim"
	fmt.Println(len(str3)) //6

	//문자열을 슬라이싱(ArrayList 와 같이 길이가 가변적인 배열)rune 타입으로 변환이 가능하다
	str3Cnt := []rune(str3)   //변환해서 길이 재보면 동하
	fmt.Println(len(str3))    //6
	fmt.Println(len(str3Cnt)) //6

	//문자열을 합치는 방법은 +, += 을 통해서 문자열을 연결하는 것이 가능
	connect1 := "Hello"
	connect2 := "World"
	fmt.Println(connect1 + " " + connect2) //Hello World
	connect1 += connect2
	fmt.Println(connect1 + connect2) //HelloWorldWorld

	//문자열을 비교하는 방법 : ==, != 을 통해서 문자열이 같은지 아닌지 검증 가능
	same1 := "iPhone"
	same2 := "IPhone"
	same3 := "iPhone"
	fmt.Printf("%s == %s : %v\n", same1, same2, same1 == same2)
	fmt.Printf("%s != %s : %v\n", same1, same2, same1 != same2)
	fmt.Printf("%s == %s : %v\n", same1, same3, same1 == same3)
	fmt.Printf("%s != %s : %v\n", same1, same3, same1 != same3)

	//문자열의 대소를 비교하는 방법 : <, <=, >, >= 을 통해서 문자열의 대소를 비교하는 것이 가능
	//문자열의 순서를 기준으로 비교된다 > "문자열의 UTF-8 변환 값" 을 기준으로 비교되는 점만 참고

	//C 에서 공부해보면 알다시피 string 이라는 놈은 char[] 으로 처리하기 힘들어서 언어에서 구현해서 제공하는 객체이다
	//역시 동일하게 go 에서도 string 의 내부 구조는 struct 으로 구성되어 있다
	//string 이라는 struct 내부에서는 2개의 필드가 존재한다 > Data 라는 포인터 + 문자열의 길이를 나타내는 int 타입의 길이 필드
	//여기에서 파생되는 이야기로는 String 의 특정 위치의 문자를 변환이 가능하냐는거다
	//ex) hella 라는 문자열에서 [4] 위치의 문자를 변경하고 싶다고 해서 단순 대입을 통해서는 불가능하다는거다

	//결론적으로 string 이라는 타입의 바이트 크기는 그럼 포인터 + int 이기 때문에 (8 + 8(int에서 따로 뒤에 숫자 안붙히면 기본적으로 8)) = 16 바이트가 된다

}

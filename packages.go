package main

/*
*
### Go Packages
go 에서는 기본적으로 라이브러리를 import 받아서 사용하는 구조인데 <br>
Spring 에서 사용하는 라이브러리 넥서스가 Maven Repository 였으면 <br>
여기는
링크를 통해 다운로드 받아서 사용하는 방식이다 <br>
유명한 곳
- https://golang.org/pkg/
- https://github.com/aveline/awesome-go

<br>
몇개는 정말 다양하게 활용되는 라이브러리 몇개를 리스트업 해봤다

- fmt : 표준 입출력 기능
- crypto : 암호화 기능
- net : 네트워크 기능
- http : 웹 프로토콜 기능
- math : 수학 관련 기능
*/

import (
	"fmt"
	"math/rand" //math 내부의 rand 패키지만 따로 import 가능

	_ "net" //goloang 의 컴파일러는 안쓰면 빨간줄을 매겨버리는 독한 아이이다... 그래도 import 에서 남겨야만 하는건 남겨야하는 패키지의 맨 앞에 _ 을 붙혀서 무시하도록 설정해주자
	//+ _ 을 통해서 unused import 작업을 해줘도 기본적으로 해당 패키지의 init() 이라는 함수는 실행되는 점을 기억하자
	//init() 메소드란, 반드시 입력 매개변수가 없어야 하고 반환값도 없어야 하는 함수이다
	//+ 추가 내용은 가장 아래에서 읽자=

	htemplate "html/template" //import 문에서 별칭을 설정해주는 것이 가능하다
	"text/template"           //이거랑 밑에랑 처럼 패키지 명이 겹치는 케이스에는

	"github.com/guptarohit/asciigraph" //외부 저장소에서 라이브러리를 땡겨온다(처음에 적으니까 빨간줄이 뜸 > 그래서 go mod tidy 을 go.mod 랑 같은 위치의 terminal 에서 수행해주니까 > 패키지를 다운받아서 external library 에 들어갔고 > go.mod 에 버전 정보가 추가됨(go.sum 추가됨)
)

func main() {
	fmt.Println("이게 바로 랜덤으로 뽑힌 숫자다 >>>> ", rand.Intn(100))

	//이렇게 패키지 명이 겹치는 경우에는 별칭을 설정해주는 것이 좋다
	template.New("test")
	htemplate.New("test")

	/**
	Go 모듈 > go build 을 수행하기 위해서는 go.mod 가 필요하다(약간 build.gradle 느낌)
	go.mod 에 작성되어있는 걸 보고 뭐 외부 저장소 패키지 같은 것들을 go.sum 을 통해서 외부패키지 + 내부패키지 해서 실행 파일을 만드는 구조이다

	go 모듈을 만드는 방법은 go mod init [모듈명] 을 통해서 생성이 가능하다
	*/

	//외부 라이브러리 사용법 확인
	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}

/*
*
패키지를 만들거나 패키지를 외부에 공개할때 규칙이 있음
- 패키지명은 소문자로 시작하도록 만들고, 최대한 간단하게 만들어라
- 아래 처럼 패키지의 전역으로 선언된 변수/상수/타입/함수/메소드 들 중 반드시 대문자로 시작하는 것들만 외부로 공개된다
*/
var Test1 = "Test1"
var test2 = "Test2"

/**
패키지가 초기화되는 순서가 있다
패키지를 입포트하면 임포트된 패키지가 수행되는데 그 수행의 순서는 아래와 같다
패키지 내부의 모든 전역 변수들의 초기화 > init() 메소드 함수 호출
*/

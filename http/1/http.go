package main

import (
	"fmt"
	"net/http"
)

/**
http 를 사용하기 위해서는 /net/http 패키지를 통해서 생성이 가능함
웹 서버를 만들기 위해서는 핸들러를 등록하고, 웹서버를 기동시키는 2가지의 단계가 필요함
*/

/*
*
핸들러 등록
핸들러란, http 요청 url 이 수신되었을 때 해당 요청을 처리해주는 함수 또는 객체
핸들러는 HandleFunc() 함수를 통해서 등록할 수 있으며 Handle() 함수는 http Handler 인터페이스를 구현한 객체를 등록할 수 있음
그러면 이제 서버를 띄웠을 때 url 경로에 해당하는 http 요청을 받으면 핸들러에 해당하는 함수를 호출하거나 http.Handler 객체의 인터페이스인 ServeHTTP() 메소드를 호출하여 요청에 따른 로직을 수행할 수 있음
아래 코드 참조
*/
func IndexPathHandler(w http.ResponseWriter, r *http.Request) {
	//todo do something ~

	//Request 객체 내부에는 Http 를 요청하는데 있어서 필요한 정보들이 들어있음
	//Http Method(GET, POST, PUT, DELETE 와 같은 http 메소드 정보)
	//Url
	//Http Protocol version 정보도 가지고 있음
	//Header 정보 > map[string][]string 형식으로 string - []string 형태로 header 의 값이 분산되어 Header 객체에 들어감

}

//func main() {
//	http.HandleFunc("/", IndexPathHandler) //이렇게 설정해두면 "/"라는 path 으로 url 호출이 들어왔을 때 IndexPathHandler 라는 함수를 호출한다는 의미임
//}

//웹 서버를 시작하는 방법
//각 url 에 대한 handler 를 처리해두면, 이젠 웹서버를 띄우고 이 웹서버를 통해서 url 을 받으면 되는데 이때는 ListenAndServe(addr string, handler Handler) 를 사용한다
//addr 은 HTTP 요청을 수신하는 주소를 나타냄, 포트번호를 적는 곳임, 그리고 handler 는 핸들러 인스턴스턴스를 넣어줘야하는데 nil 로 넣어주면 디폴트 핸들러가 들어감

func main() { //아니 패키지가 main 이여야만 실행이 가능하네 왜지
	http.HandleFunc("/", func(http http.ResponseWriter, request *http.Request) {
		fmt.Println("My First Web Server!!\nHELLO WORLD!!!")
	})

	err := http.ListenAndServe(":4000", nil) //오.. http://localhost:4000 으로 요청하니 구현한 handler 에 맞춰서 잘 노출되는구만
	if err != nil {
		fmt.Println(" This is Error")
		return
	}

	//결국 웹 서버를 구성하는 방법은 핸들러를 쭉~~~~ 등록해두고, ListenAndServe 을 통해서 포트를 열고 대기한다 그리고 해당 포트의 내용들을 구성한다
	//http 동작원리
	//우리가 웹에서 localhost:4000 땅 때렸다고 가정하면
	//기본적으로 웹 브라우저에서 DNS 시스템을 뒤지면서 도메인에 해당하는 ip 주소를 찾는다
	//localhost 이니 127.0.0.1 을 찾고, 해당 ip 의 포트번호에 맞는 곳으로 라우팅한다
	//만약에 http 로 그냥 요청하면서 포트를 설정하지 않는다면 기본적으로는 80 포트로 라우팅
	//만약에 https 로 그냥 요청하면서 포트를 설정하지 않는다면 기본적으로는 443 포트로 라우팅
	//그리고 요청했을 때 전달하는 게 이게 웹에서 제공하는 특정 데이터들이다, json 이될수도 있고, html 과 같은 문서가될 수도 있음
}

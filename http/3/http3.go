package main

import "net/http"

/**
기본적으로 웹을 서버를 수행할 때 static 패키지 내에 있는 이미지파일을 기반으로 서버를 실행시키면
static 패키지 내에 있는 이미지를 읽어서 보려고 함
*/

func main() {
	//기본적으로 서버를 구동시키면, 루트 디렉토리를 기준으로 static 을 찾아서 작업하는데
	//내가 만든건 http 패키지 밑에다가 static 을 만들어놔서 못찾는거같음 >> 명시적으로 http/static 으로 처리
	//handle 메소드에서 확실하게 FileServer 로 처리해주고, 어디 디렉토리르 기반으로 찾을것인지를 명시해줘야함
	http.Handle("/", http.FileServer(http.Dir("http/static")))

	http.ListenAndServe(":3000", nil)
}

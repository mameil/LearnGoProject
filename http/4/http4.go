package main

import (
	"fmt"
	"net/http"
	"strconv"
)

/**
http 요청을 보내는데 있어서 다양한 인자들을 넣어서 보내는 것이 가능하다
쿼리 파라미터라는 개념은 url 뒤에 넣어주는 인자를 의미함
예를 들어서 쿼리 파라미터로 name 이 kdshim 이라는 것을 넣어주고 싶다면
localhost:3000?name=kdshim
이렇게 넣어주면 됨 예시는 아래를 확인
*/

func queryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[REQ] STARED")
	values := r.URL.Query() //해당 메소드를 통해서 쿼리를 가져옴
	//id 랑 name 을 받아보쟈
	id, err := strconv.Atoi(values.Get("id"))
	if err != nil {
		fmt.Println("id 값에는 숫자를 넣어줘야 함!")
	} else {
		fmt.Printf("쿼리 파라미터로 전달받은 id 값은 %d 입니다! \n", id)
	}

	name := values.Get("name")
	if name == "" {
		fmt.Println("name 값에 공백이 들어왔습니다! ")
	} else {
		fmt.Printf("쿼리 파라미터로 전달받은 name 값은 %s 입니다! \n", name)
	}
	fmt.Println("[REQ] ENDED")
}

func main() {
	http.HandleFunc("/test", queryHandler)
	http.ListenAndServe(":3000", nil)
}

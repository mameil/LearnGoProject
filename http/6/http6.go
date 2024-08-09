package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
http 요청에 따라서 json 데이터를 내려주는 방식도 존재함

*/

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler6() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(Student{"kdshim", 28, 95}) //오.. 객체 to json 라이브러리..
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler6())
}

package main

import (
	"fmt"
	"net/http"
)

func MakeHandler2() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/qq", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, " /qq 경로롤 api 가 호출되었습니다")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "/bar 경로의 api 가 호출되었습니다")
	})

	return mux
}

func main() {
	http.ListenAndServe(":3000", MakeHandler2())
}

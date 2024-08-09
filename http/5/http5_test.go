package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/qq", nil)

	mux := MakeHandler2()
	mux.ServeHTTP(res, req) //이걸 통해서 req 를 날리고, res 를 받아서 넣어준다

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal(" /qq 경로롤 api 가 호출되었습니다", string(data))
}

func TestBarHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := MakeHandler2()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("/bar 경로의 api 가 호출되었습니다", string(data))
}

func TestErrorCase(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/qwer", nil)

	mux := MakeHandler2()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusNotFound, res.Code)
	fmt.Println(req)
	fmt.Println(res)

}

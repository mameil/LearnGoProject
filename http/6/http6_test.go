package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStudentHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil)

	mux := MakeWebHandler6()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	myjson, _ := io.ReadAll(res.Body)
	var parsedStudent Student
	json.Unmarshal(myjson, &parsedStudent)

	assert.Equal("kdshim", parsedStudent.Name)
	assert.Equal(28, parsedStudent.Age)
	assert.Equal(95, parsedStudent.Score)
}

package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
*goos: windows
goarch: amd64
pkg: learn-go-project/test
cpu: 12th Gen Intel(R) Core(TM) i7-1260P
BenchmarkFibonacciRecursive
BenchmarkFibonacciRecursive-16    	1000000000
BenchmarkFibonacciIterate
BenchmarkFibonacciIterate-16      	1000000000
PASS

아따 컴터 좋네
*/
func BenchmarkFibonacciRecursive(t *testing.B) {
	assert1 := assert.New(t)

	assert1.Equal(0, fibonacciRecursive(-1), "should be 0")
	assert1.Equal(0, fibonacciRecursive(0), "should be 0")
	assert1.Equal(1, fibonacciRecursive(1), "should be 1")
	assert1.Equal(2, fibonacciRecursive(3), "should be 3")
	assert1.Equal(233, fibonacciRecursive(13), "should be 223")
}

func BenchmarkFibonacciIterate(t *testing.B) {
	assert1 := assert.New(t)

	assert1.Equal(0, fibonacciIterate(-1), "should be 0")
	assert1.Equal(0, fibonacciIterate(0), "should be 0")
	assert1.Equal(1, fibonacciIterate(1), "should be 1")
	assert1.Equal(2, fibonacciIterate(3), "should be 3")
	assert1.Equal(233, fibonacciIterate(13), "should be 223")
}

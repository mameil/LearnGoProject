package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoTest1(t *testing.T) {
	testObj := multiply(9)
	if testObj != 81 {
		t.Errorf("multiply(9) should be 81 but multiply(9) returned %d", testObj)
	}
}

func TestGoTest2(t *testing.T) {
	testObj := multiply(3)
	if testObj != 9 {
		t.Errorf("multiply(3) should be 9 but multiply(3) returned %d", testObj)
	}
}

func TestGoTestTestifyVer(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(81, multiply(9), "multiply(9) should be 81")
}

func TestGoTestTestifyVer2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(10, multiply(3), "multiply(3) should be 9")
}

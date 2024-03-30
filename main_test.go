package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TetsCases = []struct {
	num int
	errorExpected bool
} {
	{num: 73, errorExpected: false},
	{num: 27, errorExpected: true},
	{num: 41, errorExpected: false},
	{num: 100, errorExpected: true}, 
}
func Test_alpha_IsPrime(t *testing.T)  {
	for _, testCase := range TetsCases {
		if isPrime(testCase.num) == "Is Prime" && testCase.errorExpected {
			t.Errorf("Wrong test case: %d is prime", testCase.num)
		} else if isPrime(testCase.num) != "Is Prime" && !testCase.errorExpected {
			t.Errorf("Wrong test case: %d is prime", testCase.num)
		}
	}
}

func Test_alpha_MainIsPrime(t *testing.T)  {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	n := 31
	isprime := isPrime(n)
	fmt.Println(isprime)

	w.Close()

	readAll, _ := io.ReadAll(r)
	res := string(readAll)

	os.Stdout = stdOut

	if !strings.Contains(res, "Is Prime") {
		t.Errorf("Failed test case")
	}

	// also another way
	assert.Equal(t, "Is Prime", isprime)
}

func Test_MainNotPrime(t *testing.T)  {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	n := 18
	isprime := isPrime(n)
	fmt.Println(isprime)

	w.Close()

	readAll, _ := io.ReadAll(r)
	res := string(readAll)

	os.Stdout = stdOut

	if !strings.Contains(res, "Not Prime") {
		t.Errorf("Failed test case")
	}

	// also another way
	assert.Equal(t, "Not Prime", isprime)
}
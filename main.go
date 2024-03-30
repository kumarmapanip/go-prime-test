package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	N = 105
	Prime = make([]int, N)
)

func init() {
	for i := 0; i < N; i++ {
		Prime[i] = 1
	}
	Prime[0] = 0
	Prime[1] = 0
	
	for i := 2; i*i < N; i++ {
		if Prime[i] == 1 {
			for j := 2*i; j < N; j+=i {
				Prime[j] = 0
			}
		}
	}
}

func isPrime(n int) string {
	if Prime[n] == 1 {
		return "Is Prime"
	}
	return "Not Prime"
}

func main() {
	quit := make(chan bool)

	fmt.Println("Execution Started")
	go solve(quit)

	// execution is blocked until channel recieves data
	<-quit
	fmt.Println("Execution Completed")
}

func solve(quit chan bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()

		text := scanner.Text()

		if strings.EqualFold(text, "quit") {
			quit <- true
			return
		}
		num, _ := strconv.Atoi(text)
		fmt.Println(isPrime(num))
	}
}
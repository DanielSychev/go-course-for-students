package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	// TODO тут напишите цикл с вызовом FizzBuzz
	for i := 1; i <= 100; i++ {
		fmt.Print(fizzbuzz.FizzBuzz(i), " ")
	}
}

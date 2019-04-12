package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closures() {
	// noew nextInt is a function with it's own i value
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println()

	newInts := intSeq()
	fmt.Println(newInts())
}

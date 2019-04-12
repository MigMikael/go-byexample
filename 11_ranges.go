package main

import "fmt"

func ranges() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	fmt.Println()

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cat"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Println()

	for k := range kvs {
		fmt.Println("key:", k)
		fmt.Println("value:", kvs[k])
	}
	fmt.Println()

	for i, c := range "go lang โก" {
		fmt.Println(i, c)
	}
	fmt.Println()
}

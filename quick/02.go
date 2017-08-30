package main

import "fmt"

func main() {
	var ori = []int{5, 4, 3, 2, 1}
	fmt.Println("before ori is")
	for item := range ori {
		fmt.Printf("%d \t", ori[item])
	}

	result := test(ori, 5)

	fmt.Println("\nafter ori is")
	for item := range ori {
		fmt.Printf("%d \t", ori[item])
	}

	fmt.Println("\nresult is")
	for item := range result {
		fmt.Printf("%d \t", result[item])
	}
}

func test(arr []int, size int) []int {
	var i int

	for i = 0; i < size; i++ {
		arr[i] = i
	}

	return arr
}

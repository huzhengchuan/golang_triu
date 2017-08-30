package main

import (
	"fmt"
)

func main() {
	arr := [5]int{5, 2, 3, 4, 5}

	for item, value := range arr {
		fmt.Printf("%d %d\r\n", arr[item], value)
	}

	fmt.Println("---------------------")
	for loop := 0; loop < len(arr); loop++ {
		fmt.Println(arr[loop])
	}

	fmt.Println("---------------------")
	sli := []string{"Red", "Blue", "Green"} //sli := make([]string, 3, 5)

	for item, value := range sli {
		fmt.Println(sli[item] + " " + value)
	}

	fmt.Println("---------------------")
	for loop := 0; loop < len(sli); loop++ {
		fmt.Println(sli[loop])
	}

	fmt.Println("---------------------")
	ma := make(map[string]string)
	ma["hello"] = "->hello"
	ma["world"] = "->world"
	for k, v := range ma {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}

}

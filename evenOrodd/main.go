package main

import "fmt"

func main() {
	sliceOfints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, ele := range sliceOfints {
		if ele%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}

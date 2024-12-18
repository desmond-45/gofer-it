package main

import "fmt"

func main() {
	intList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, ele := range intList {
		kind := "odd"
		if ele&1 == 0 {
			kind = "even"
		}
		fmt.Printf("Number %d is %s\n", ele, kind)
	}
}

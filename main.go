package main

import (
	"fmt"

	ch1 "github.com/ashcoder666/boot.dev.ds/CH1"
)

func main() {

	// res, res1 := ch1.Find_minimum([]int{2, 4, 6, 1, 3, 8})

	res := ch1.FindMedian([]int{1, 2, 3, 4})

	fmt.Println("res", res)
}

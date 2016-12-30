package main

import "fmt"

func main() {
	r := func(nums ...int) int {
		max := nums[0]
		for _, i := range nums {
			if i > max {
				max = i
			}
		}
		return max
	}(1, 2, 3, 4, 5)

	fmt.Println(r)
}

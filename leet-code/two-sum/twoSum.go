package main

import "fmt"

func main() {
	values := [4]int{2, 7, 11, 15}

	var nums []int
	var target int

	nums = values[0:]
	target = 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	var response []int
	dict := map[int]int{}

	for i, num := range nums {
		result := target - num

		if index, contains := dict[result]; contains {
			response = append(response, index, i)
			break
		} else {
			dict[num] = i
		}
	}

	return response
}

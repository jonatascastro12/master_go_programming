package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("%#v\n", nums)

	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))
	nums = append(nums, 1, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 4)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 100)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums[0:4], 200, 300, 400, 500, 600)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	fmt.Println(nums)

	letters := []string{"A", "B", "C", "D", "E", "F"}

	letters = append(letters[0:1], "X", "Y")
	fmt.Println(letters, len(letters), cap(letters))

	// fmt.Println(letters[4])
	fmt.Println(letters[4:5][0])
}

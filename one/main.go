package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const Sum int = 2020

func twoSum(nums []int, sum int) (int, int) {
	seen := make(map[int]bool)
	for _, num := range nums {
		if _, ok := seen[sum-num]; ok {
			return num, sum - num
		}
		seen[num] = true
	}
	return 0, 0
}

func threeSum(nums []int, sum int) (int, int, int) {
	for i, num := range nums {
		if Sum - num > 0  {
            rest := make([]int, len(nums)-1)
            rest = append(nums[:i])
            rest = append(nums[i+1:])
			num1, num2 := twoSum(rest, Sum-num)
			if num1 != 0 && num2 != 0 {
				return num, num1, num2
			}

		}
	}
	return 0, 0, 0
}

func main() {
	// read in file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// convert to slice of ints
	scanner := bufio.NewScanner(file)
	nums := make([]int, 0)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, i)
	}
	// solve first part
	num1, num2 := twoSum(nums, Sum)
	fmt.Println(num1 * num2)

	num1, num2, num3 := threeSum(nums, Sum)
	fmt.Println(num1 * num2 * num3)

}

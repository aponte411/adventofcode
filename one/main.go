package main

import (
	"fmt"
	"log"
    "strconv"
    "bufio"
    "os"
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
        i, _:= strconv.Atoi(scanner.Text())
        nums = append(nums, i)
    }
	// solve first part
	num1, num2 := twoSum(nums, Sum)
    fmt.Println(num1, num2)
	fmt.Println(num1 * num2)

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func IsValid(min, max int, char string, str string) bool {
	count := 0
	for _, ch := range str {
		if string(ch) == char {
			count += 1
		}
	}
	if count <= max && count >= min {
		return true
	} else {
		return false
	}
}

func IsValid2(min, max int, char string, password string) bool {
	return (char == string(password[min])) != (char == string(password[max]))
}

func StringToInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputs := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	numOfValidPasswords1 := 0
	numOfValidPasswords2 := 0

	for _, line := range inputs {
		// Get password and policy
		pol_and_pass := strings.Split(line, ":")
		policy, password := pol_and_pass[0], pol_and_pass[1]
		// parse policy
		res := strings.Split(policy, "-")
		min, rest := StringToInt(res[0]), res[1]
		tmp := strings.Split(rest, " ")
		max, char := StringToInt(tmp[0]), tmp[1]
		// Part 1, if valid count
		if IsValid(min, max, char, password) {
			numOfValidPasswords1 += 1
		}
        // Part 2
		if IsValid2(min, max, char, password) {
			numOfValidPasswords2 += 1
		}
	}
	fmt.Printf("Final number of valid passwords Part 1: %v, Out of: %v\n", numOfValidPasswords1, len(inputs))

	fmt.Printf("Final number of valid passwords Part 2: %v, Out of: %v\n", numOfValidPasswords2, len(inputs))
}

package main

import (
	"bufio"
    "strconv"
	"fmt"
	"log"
	"os"
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
    numOfValidPasswords := 0
    for _,line := range inputs {
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
            numOfValidPasswords += 1
        }

    }
    fmt.Printf("Final number of valid passwords: %v, Out of: %v\n", numOfValidPasswords, len(inputs))
}

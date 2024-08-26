package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error while reading file", err)
	}
	input := string(data)
	lines := strings.Split(input, "\n")
	result := 0
	copies := make([]int, len(lines))
	for i := range copies {
		copies[i] = 1
	}
	for i, line := range lines {
		res := 0
		if len(line) < 1{
			continue
		}
		list := strings.Split(line, ":")[1]
		numList := strings.Split(list, "|")
		wNumbers := strings.Split(numList[0], " ")

		numbers := strings.Split(numList[1], " ")
		
		for _, target := range wNumbers {
			if len(target) <= 0 {
				continue
			}
			for _, n := range numbers {
				if n == target {
					res++
					break
				}
			}

		}
		if res > 0 {
			for j := 0; j < res; j++ {
				copies[i+j+1] += copies[i]
			}
			
		}
	}
	for _, x := range copies{
		result += x
	}
	fmt.Println("result: ", result)

}
func power(base, exponent int) int {
	result := 1
	for exponent > 0 {
		if exponent%2 == 1 {
			result *= base
		}
		base *= base
		exponent /= 2
	}
	return result
}

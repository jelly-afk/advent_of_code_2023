package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	eleMap := make(map[string][]string)
	instructions := lines[0]
	var currEle []string
	for i, line := range lines {
		if i < 2 {
			continue
		}
		parts := strings.Split(line, "=")
		key := strings.TrimSpace(parts[0])
		if string(key[2]) == "A" {
			currEle = append(currEle, key)

		}

		parts[1] = strings.Trim(parts[1], " ()")
		values := strings.Split(strings.ReplaceAll(parts[1], " ", ""), ",")
		eleMap[key] = values

	}
	result := make([]int, len(currEle))
	for i, _ := range currEle {

		idx := 0
		for {
			currIns := instructions[idx]
			if string(currIns) == "L" {
				currEle[i] = eleMap[currEle[i]][0]
			} else {

				currEle[i] = eleMap[currEle[i]][1]
			}
			result[i]++
            if string(currEle[i][2]) == "Z"{
                break
            }
			idx++
			if idx == len(instructions) {
				idx = 0

			}
		}

	}
	println(lcmOfArray(result))

}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a*b) / gcd(a, b)
}
func lcmOfArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	lcmValue := numbers[0]
	for i := 1; i < len(numbers); i++ {
		lcmValue = lcm(lcmValue, numbers[i])
	}
	return lcmValue
}
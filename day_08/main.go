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
	currEle := "AAA"
	for i, line := range lines {
		if i < 2 {
			continue
		}
		parts := strings.Split(line, "=")
		key := strings.TrimSpace(parts[0])
		

		parts[1] = strings.Trim(parts[1], " ()")
		values := strings.Split(strings.ReplaceAll(parts[1], " ", ""), ",")
		eleMap[key] = values

	}
	idx := 0
	result := 0
	for {
		
		if currEle == "ZZZ" {
			break
		}
		currIns := instructions[idx]
		val := eleMap[currEle]
		if string(currIns) == "L" {
			currEle = val[0]

		} else if string(currIns) == "R" {
			currEle = val[1]
		} else {
            log.Fatal("ins err")
        }
		idx++
		if idx == len(instructions) {
			idx = 0
		}
        result++
	}
	println("Result: ", result)

}

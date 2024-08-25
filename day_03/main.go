package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error while reading file: ", err)
	}
	lines := strings.Split(string(data), "\n")
	result := 0
	added := make(map[string]bool)
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			c := rune(lines[i][j])
			if c != '.' && !unicode.IsDigit(c) {
				if j > 0 && unicode.IsDigit(rune(lines[i][j-1])) {
					n := findNum(lines[i], j, 0)
					_, exists := added[fmt.Sprintf("%d %d", i, n)]
					if !exists {

						num, err := strconv.Atoi(lines[i][n:j])
						if err != nil {
							log.Fatal("Error while converting string: ", num)
						}
						result += num
						added[fmt.Sprintf("%d %d", i, n)] = true
					}
				}
				if j < len(lines[i])-1 && unicode.IsDigit(rune(lines[i][j+1])) {
					_, exists := added[fmt.Sprintf("%d %d", i, j+1)]
					if !exists {
						n := findNum(lines[i], j, 1)
						num, err := strconv.Atoi(lines[i][j+1 : n])
						if err != nil {
							log.Fatal("Error while converting string: ", num)
						}
						result += num
						added[fmt.Sprintf("%d %d", i, n)] = true

					}
				}
				if i > 0 && unicode.IsDigit(rune(lines[i-1][j])) {
					s := findNum(lines[i-1], j, 0)
					_, exists := added[fmt.Sprintf("%d %d", i-1, s)]
					if !exists {
						e := findNum(lines[i-1], j, 1)
						num, err := strconv.Atoi(lines[i-1][s:e])
						if err != nil {
							log.Fatal("error while converting string")
						}
						result += num
						added[fmt.Sprintf("%d %d", i-1, s)] = true
					}
				} else {
					if i > 0 && j > 0 && unicode.IsDigit(rune(lines[i-1][j-1])) {
						n := findNum(lines[i-1], j, 0)
						_, exists := added[fmt.Sprintf("%d %d", i-1, n)]
						if !exists {

							num, err := strconv.Atoi(lines[i-1][n:j])
							if err != nil {
								log.Fatal("error while converting string: ", err)

							}
							result += num
							added[fmt.Sprintf("%d %d", i-1, n)] = true
						}
					}
					if i > 0 && j < len(lines[i])-1 && unicode.IsDigit(rune(lines[i-1][j+1])) {
						_, exists := added[fmt.Sprintf("%d %d", i-1, j+1)]
						if !exists {

							n := findNum(lines[i-1], j, 1)
							num, err := strconv.Atoi(lines[i-1][j+1 : n])
							if err != nil {
								log.Fatal("error while converting string: ", err)
							}
							result += num
							added[fmt.Sprintf("%d %d", i-1, j+1)] = true

						}
					}
				}
				if i < len(lines)-1 && unicode.IsDigit(rune(lines[i+1][j])) {
					s := findNum(lines[i+1], j, 0)
					_, exists := added[fmt.Sprintf("%d %d", i+1, s)]
					if !exists {
						e := findNum(lines[i+1], j, 1)
						num, err := strconv.Atoi(lines[i+1][s:e])
						if err != nil {
							log.Fatal("error while converting string")
						}
						result += num
						added[fmt.Sprintf("%d %d", i+1, s)] = true
					}
				} else {
					if i < len(lines)-1 && j > 0 && unicode.IsDigit(rune(lines[i+1][j-1])) {
						n := findNum(lines[i+1], j, 0)
						_, exists := added[fmt.Sprintf("%d %d", i+1, n)]
						if !exists {

							num, err := strconv.Atoi(lines[i+1][n:j])
							if err != nil {
								log.Fatal("error while converting string: ", err)

							}
							result += num
							added[fmt.Sprintf("%d %d", i+1, n)] = true
						}
					}
					if i < len(lines)-1 && j < len(lines[i]) && unicode.IsDigit(rune(lines[i+1][j+1])) {
						_, exists := added[fmt.Sprintf("%d %d", i+1, j+1)]
						if !exists {

							n := findNum(lines[i+1], j, 1)
							num, err := strconv.Atoi(lines[i+1][j+1 : n])
							if err != nil {
                                fmt.Println(lines[i-1])
                                fmt.Println(lines[i])
								log.Fatal("error while converting string: ", err)
							}
							result += num
							added[fmt.Sprintf("%d %d", i+1, j+1)] = true

						}
					}
				}

			}
		}
	}
	println("result: ", result)
}
func findNum(line string, current int, x int) int {

	if x == 0 {
		n := current - 1
		for n >= 0 {
			if unicode.IsDigit(rune(line[n])) {
				n--
			} else {
				break
			}
		}
		return n + 1

	}
	if x == 1 {
		n := current + 1
		for n < len(line) {
			if unicode.IsDigit(rune(line[n])) {
				n++
			} else {
				break
			}
		}
		return n
	}
	return -1

}

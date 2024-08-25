package main

import (
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
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			c := rune(lines[i][j])
			if c == '*' {
				var numbers []int
				if j > 0 && unicode.IsDigit(rune(lines[i][j-1])) {
					n := findNum(lines[i], j, 0)

					num, err := strconv.Atoi(lines[i][n:j])
					if err != nil {
						log.Fatal("Error while converting string: ", num)
					}
					numbers = append(numbers, num)
				}

				if j < len(lines[i])-1 && unicode.IsDigit(rune(lines[i][j+1])) {
					n := findNum(lines[i], j, 1)
					num, err := strconv.Atoi(lines[i][j+1 : n])
					if err != nil {
						log.Fatal("Error while converting string: ", num)
					}
					numbers = append(numbers, num)

				}

				if i > 0 && unicode.IsDigit(rune(lines[i-1][j])) {
					s := findNum(lines[i-1], j, 0)
					e := findNum(lines[i-1], j, 1)
					num, err := strconv.Atoi(lines[i-1][s:e])
					if err != nil {
						log.Fatal("error while converting string")
					}
					numbers = append(numbers, num)
				} else {
					if i > 0 && j > 0 && unicode.IsDigit(rune(lines[i-1][j-1])) {
						n := findNum(lines[i-1], j, 0)

						num, err := strconv.Atoi(lines[i-1][n:j])
						if err != nil {
							log.Fatal("error while converting string: ", err)

						}
						numbers = append(numbers, num)
					}
				
				if i > 0 && j < len(lines[i])-1 && unicode.IsDigit(rune(lines[i-1][j+1])) {

					n := findNum(lines[i-1], j, 1)
					num, err := strconv.Atoi(lines[i-1][j+1 : n])
					if err != nil {
						log.Fatal("error while converting string: ", err)
					}
					numbers = append(numbers, num)

				}
            }

				if i < len(lines)-1 && unicode.IsDigit(rune(lines[i+1][j])) {
					s := findNum(lines[i+1], j, 0)
					e := findNum(lines[i+1], j, 1)
					num, err := strconv.Atoi(lines[i+1][s:e])
					if err != nil {
						log.Fatal("error while converting string")
					}
					numbers = append(numbers, num)
				} else {
					if i < len(lines)-1 && j > 0 && unicode.IsDigit(rune(lines[i+1][j-1])) {
						n := findNum(lines[i+1], j, 0)

						num, err := strconv.Atoi(lines[i+1][n:j])
						if err != nil {
							log.Fatal("error while converting string: ", err)

						}
						numbers = append(numbers, num)
					}
				
				if i < len(lines)-1 && j < len(lines[i]) && unicode.IsDigit(rune(lines[i+1][j+1])) {

					n := findNum(lines[i+1], j, 1)
					num, err := strconv.Atoi(lines[i+1][j+1 : n])
					if err != nil {
						log.Fatal("error while converting string: ", err)
					}
					numbers = append(numbers, num)

				}
            }
              
              if len(numbers) == 2 {
                result = result + (numbers[0] * numbers[1])
              }
			}

		}

	}
    println("result: ",result)
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

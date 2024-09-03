package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	var cards [7][][]string
	cMap := map[string]int{"2": 0, "3": 1, "4": 2, "5": 3, "6": 4, "7": 5, "8": 6, "9": 7, "T": 8, "J": -1, "Q": 10, "K": 11, "A": 12}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		hand, bid := parts[0], parts[1]
		count := make(map[rune]int)
		jCount := 0
		for _, c := range hand {
			if c == 'J' {
				jCount++
				continue
			}
			count[c]++

		}
		maxC := 0
		for _, val := range count {
			if val > maxC {
				maxC = val
			}
		}
        cLen := len(count)
		idx := 0
		if cLen == 1 || cLen == 0{
			idx = 6
        } else if maxC == 4 {
            idx = 5
        } else if maxC == 3 {
            if jCount == 1 {
                idx = 5
            } else if cLen == 2 {
                idx = 4
            } else {
                idx = 3
            }
        } else if maxC == 2 {
            if jCount == 2 {
                idx = 5
            } else if jCount == 1 && cLen == 2 {
                idx = 4
            } else if jCount == 1 && cLen == 3 {
                idx = 3
            } else if cLen == 3 {
                idx = 2
            } else if cLen == 4 {
                idx = 1
            }
        } else if maxC == 1 {
            if jCount == 3 {
                idx = 5
            } else if jCount == 2 {
                idx = 3
            } else if jCount == 1 {
                idx = 1
            } else {
                idx = 0
            }
        }
	
		cards[idx] = append(cards[idx], []string{hand, bid})
	}
	for _, cardArr := range cards {
		sort.SliceStable(cardArr, func(i, j int) bool {
			iVal := cardArr[i][0]
			jVal := cardArr[j][0]
			for i, c := range iVal {
				if cMap[string(c)] > cMap[string(jVal[i])] {
					return false
				} else if cMap[string(c)] < cMap[string(jVal[i])] {
					return true

				}

			}
			return false

		})

	}
	result := 0
	i := 1
	for _, cArr := range cards {
		for _, arr := range cArr {
			bidAmt, err := strconv.Atoi(arr[1])
			if err != nil {
				log.Fatal(err)

			}
			result += (bidAmt * i)
			i++
		}
	}
	println("result: ", result)

}

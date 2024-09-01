package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("error while reading data from file", err)
	}
	input := string(data)
	maps := strings.SplitN(input, "\n\n", 2)
	seedStr := strings.Split(maps[0], ":")
	seeds := strings.Split(strings.TrimSpace(seedStr[1]), " ")
	maps = strings.Split(maps[1], "\n\n")
	var rangeArr []string
	result := math.MaxInt
	for i, seed := range seeds {
		if i%2 != 0 {
			continue
		}
		rangeArr = append(rangeArr, fmt.Sprintf("%s %s 0", seed, seeds[i+1]))
	}

	for len(rangeArr) > 0 {
		currRange := rangeArr[0]
		rangeArr = rangeArr[1:]
		seedParts := strings.Split(currRange, " ")
		seedRangeSt, err := strconv.Atoi(seedParts[0])
		found := false

		if err != nil {
			log.Fatal("error while reading data from file", err)
		}
		rangeLen, err := strconv.Atoi(seedParts[1])

		if err != nil {
			log.Fatal("error while reading data from file", err)
		}
		mapInd, err := strconv.Atoi(seedParts[2])
		lastIdx := mapInd == len(maps)-1
		if err != nil {
			log.Fatal("error while reading data from file", err)
		}

		mapStr := strings.TrimSpace(strings.Split(maps[mapInd], ":")[1])
		lines := strings.Split(strings.TrimSpace(mapStr), "\n")
		for _, line := range lines {
			values := strings.Split(line, " ")

			srcStart, err := strconv.Atoi(strings.TrimSpace(values[1]))
			if err != nil {
				log.Fatal(err)
			}
			dstStart, err := strconv.Atoi(strings.TrimSpace(values[0]))
			if err != nil {
				log.Fatal(err)
			}
			if seedRangeSt >= srcStart {
				valueRange, err := strconv.Atoi(values[2])
				if err != nil {
					log.Fatal(err)
				}
				diff := seedRangeSt - srcStart
				if diff < valueRange {
					found = true
					if seedRangeSt+rangeLen <= srcStart+valueRange {
						if lastIdx {
							if result > dstStart+diff {
								result = dstStart + diff
							}
						} else {
							rangeArr = append(rangeArr, fmt.Sprintf("%d %d %d", dstStart+diff, rangeLen, mapInd+1))
						}
					} else {
						firstLen := (srcStart+valueRange) - seedRangeSt 
						if lastIdx {
							if result > dstStart+diff {
								result = dstStart + diff
							}
						} else {
							rangeArr = append(rangeArr, fmt.Sprintf("%d %d %d", dstStart+diff, firstLen, mapInd+1))
						}
						rangeArr = append(rangeArr, fmt.Sprintf("%d %d %d", srcStart+valueRange, rangeLen-firstLen, mapInd))

					}
					break
				}

			}

		}
		if !found {
			if lastIdx {
				if result > seedRangeSt {
					result = seedRangeSt
				}
			} else {
				rangeArr = append(rangeArr, fmt.Sprintf("%d %d %d", seedRangeSt, rangeLen, mapInd+1))
			}
		}

	}
	println("result: ", result)
}

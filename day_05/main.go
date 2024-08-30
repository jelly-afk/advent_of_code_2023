package main

import (
	"log"
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
    seeds := strings.Split(strings.Trim(seedStr[1], " "), " ")
    maps = strings.Split(maps[1], "\n\n")
    prevs := make([]int, len(seeds))
    for i, seed := range seeds {
        prevs[i], err = strconv.Atoi(strings.Trim(seed, " "))
        if err != nil {
            log.Fatal("prevsErr:", err)
        }
        mapLoop:
        for _, m := range maps {
            mapStr := strings.Split(m, ":")[1]
            lines := strings.Split(mapStr, "\n")
            for _, line := range lines {
                if len(line) <= 0 {
                    continue
                }
                values := strings.Split(line, " ")


                srcStart, err := strconv.Atoi(strings.Trim(values[1], " "))
                if err != nil {
                    log.Fatal(err)
                }
                dstStart, err := strconv.Atoi(strings.Trim(values[0], " "))
                if err != nil {
                    log.Fatal(err)
                }
                if prevs[i] >= srcStart {
                    valueRange, err := strconv.Atoi(values[2])
                    if err != nil {
                        log.Fatal(err)
                    }   
                    diff :=  prevs[i]- srcStart
                    if diff < valueRange {
                        if i == 1 {
                        }
                        prevs[i] = dstStart+diff
                        continue mapLoop
                    }  

				}

			}


		}
	}
    result := prevs[0]
    for _, prev := range prevs{
        if prev < result{
            result = prev
        }
    }
    println("result: ", result)
}

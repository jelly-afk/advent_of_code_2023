package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
    data, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    input := string(data)
    parts := strings.Split(input, "\n")
    timeStr := strings.TrimSpace(strings.Split(parts[0], ":")[1])
    distStr := strings.TrimSpace(strings.Split(parts[1], ":")[1])
    timeL := strings.Split(timeStr, " ")
    distL := strings.Split(distStr, " ")
    result := 1
    for i, timeS := range timeL {
        time, err := strconv.Atoi(timeS)
        if err != nil {
            log.Fatal(err)
        }
        dist,err := strconv.Atoi(distL[i])
        if err != nil {
            log.Fatal(err)
        }
        left := 0
        right := time
        res := 1
        for right > left {
            mid := (right+left)/2
            prod := mid*(time-mid)
            if prod > dist {
                if (mid-1)*(time-mid+1) < dist {
                    res = mid
                    break
                }

                right = mid
            } else if prod < dist {
                left = mid+1
            } else {
                res = mid
                break
            }
            
        }
        result = result * ((time-res*2)+1)
        
    }
    println("result: ", result)


} 

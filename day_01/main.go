package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main () {
    data, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Println(os.Stderr, "error reading file: ", err)
    }
    input := string(data)
    lines := strings.Split(input, "\n")
    fmt.Println(lines)
    result := 0
    for _, line:= range lines{
        first := '0'
        last := '0'
        for _,c := range line {
            if c >= '0' && c <= '9' {
                first = c
                break
            }

        }
        runes := []rune(line)
        for i := len(runes) -1;i >= 0; i-- {
            if runes[i] >= '0' && runes[i] <= '9'{
                last = runes[i]
                break
            }
        }
        res := string(first) + string(last)
        temp,err :=  strconv.Atoi(res)
        if err != nil {
            log.Fatalf("error while ocnverting string", err)
        }
        result += temp 
        

    }
    println("result: ",result)
}

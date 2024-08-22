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
        log.Fatal(os.Stderr, "error reading file: ", err)
    }
    input := string(data)
    lines := strings.Split(input, "\n")
    result := 0
    numbers :=  []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    for _, line:= range lines{
        first := "0"
        last := "0"
        firstLoop:
        for i,c := range line {
            if c >= '0' && c <= '9' {
                first = string(c)
                break firstLoop
            }
            if (c == 'e' || c == 'o' || c == 'r' || c == 'x' || c == 'n' || c == 't'){
                for n, num := range numbers {
                    numLen := len(num)
                    if c == rune(num[numLen-1]) && i >= numLen-1{
                        if line[i - numLen+1:i+1] == num {
                            first = strconv.Itoa(n+1)
                            break firstLoop
                        }
                    } 
                } 
            }

        }
        lastLoop:
        for i := len(line) -1;i >= 0; i-- {
            c := string(line[i])
            if c >= "0" && c <= "9"{
                last = string(line[i])
                break lastLoop
            }
            if (c == "o" || c == "t" || c == "f" || c == "s" || c == "e" || c == "n"){
               for n, num := range numbers {
                    numLen := len(num)
                    if c == string(num[0]) && (len(line)-i) >= numLen{
                        if line[i:i+numLen] == num {
                            last = strconv.Itoa(n+1)
                            break lastLoop 
                        }
                    } 
                } 
 
            }
        }
        res := first + last
        fmt.Println("line: ", line)
        fmt.Println("res:", res)
        temp,err :=  strconv.Atoi(res)
        if err != nil {
            log.Fatalf("error while ocnverting string", err)
        }
        result += temp 
        

    }
    println("result: ",result)
}

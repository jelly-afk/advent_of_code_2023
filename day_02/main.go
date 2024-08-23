package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main (){
    data, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal("error while reading file: ", err)
    }
    input := string(data)
    lines := strings.Split(input, "\n")
    result := 0
    linesLoop:
    for _, line := range lines {
        parts := strings.Split(line, ":")
        if len(parts) < 2 {
            continue
        }
        gameId := strings.Split(parts[0], " ")[1]
        replacedStr := strings.ReplaceAll(parts[1], ",", "@")
        games := strings.Split(replacedStr ,";")
        gameId = strings.TrimSpace(gameId)

        for _, game := range games {
            rolls := strings.Split(game, "@")
                
            for _, roll := range rolls{
                roll = strings.TrimSpace(roll)
                rollParts := strings.Split(roll, " ")
                ballCount, err := strconv.Atoi(rollParts[0])
                if err != nil {
                    log.Fatal("Error while converting string")
                }
                switch rollParts[1] {
                case "red":
                    if ballCount > 12{
                        continue linesLoop
                    }
                case "green":
                    if ballCount > 13{

                        continue linesLoop
                    }
                case "blue":
                    if ballCount > 14{
                        continue linesLoop
                    }

                }
               
            }
            
        }
        res,err := strconv.Atoi(gameId)
        if err != nil {
            log.Fatal("error while converting to integer: ", err)     
        }
        result = result + res





    }
    fmt.Println("result: ", result)

}

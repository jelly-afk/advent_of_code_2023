package main

import (
	"log"
	"os"
	"strings"
)

func main (){
    data, err := os.ReadFile("input.txt") 
    if err != nil {
        log.Fatal(err)
    }
    lines := strings.Split(string(data), "\n")
    row, col, galaxies := make([]int, len(lines)), make([]int, len(lines[0])), make([][]int, 0)
    for i, line := range lines {
        for j, c := range line {
            if c == '#'{
                row[i]++
                col[j]++
                galaxies = append(galaxies, []int{i, j})

            }
        }
    }
    n := 0
    for i := range row {
        if row[i] == 0 {
            n++
        }
        row[i] = n
    }
    n = 0
    for i := range col {
        if col[i] == 0 {
            n++
        }
        col[i] = n
    }
    result := 0 
    for i, curr := range galaxies {
        for j := i+1;j < len(galaxies);j++ {
            next := galaxies[j]
            res := 0
            if curr[0] == next[0]{
                x := col[next[1]] - col[curr[1]]
                res = next[1] - curr[1] + x
            } else if curr[1] == next[1] {
                x := row[next[0]] - row[curr[0]]
                res = next[0] - curr[0] + x
            } else {

                x := diff(col[next[1]] , col[curr[1]])
                y := diff(row[next[0]] , row[curr[0]])
                res = diff(next[0] , curr[0]) + diff(next[1] , curr[1]) + x + y
            }

            result+=res
        }
    }
    println("result: ", result)

}
func diff (x, y int) int {
    if x > y {
        return x-y
    }
    return y-x

}

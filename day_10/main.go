package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
    data, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    } 
    lines := strings.Split(string(data), "\n")
    i, j := 0, 0
    mainLoop:
    for i < len(lines){
        j = 0
        for j < len(lines[0]){
            if string(lines[i][j]) == "S"{
                break mainLoop

            }
            j++
        }
        i++
    }
    dirs := map[string][][]string{
        "|": {[]string{}, []string{"7", "F", "|"}, []string{}, []string{"|", "J", "L"}}, 
        "-": {[]string{"-", "L", "F"}, []string{}, []string{"-", "J", "7"}, []string{}},
        "S":  {[]string{"-", "L", "F"}, []string{}, []string{"-", "J", "7"}, []string{}},
        "L": {[]string{}, []string{"7", "F", "|" }, []string{"-", "J", "7"},[]string{} },
        "J": {[]string{"-", "L", "F"}, []string{"7", "F", "|"}, []string{}, []string{}}   ,
        "7": {[]string{"-", "L", "F"}, []string{}, []string{}, []string{"|", "J", "L"}},
        "F": {[]string{}, []string{},  []string{"-", "J", "7"},  []string{"|", "J", "L"}},
        }

    q := [][]int{{i, j}, {-1, -1}}
    visited := make(map[string]bool)
    result := 0
    for len(q) > 0 {
        i, j := q[0][0], q[0][1]
        if len(q) == 1 {
            q = q[:0]
            
        } else {
            q = q[1:]
        }
        if i == -1 {
            if len(q) == 0 {
                break
            }
            result++
            q = append(q, []int{-1, -1})

            continue
        }
        curr := dirs[string(lines[i][j])]
        
        if j > 0 {
            ele := string(lines[i][j-1])
            if slices.Contains(curr[0], ele) && !(visited[fmt.Sprintf("%d %d", i, j-1)]){
                q = append(q, []int{i, j-1})
                visited[fmt.Sprintf("%d %d", i, j-1)] = true
            }
        }
        if i > 0 {
            ele := string(lines[i-1][j])
            if slices.Contains(curr[1], ele) && !(visited[fmt.Sprintf("%d %d", i-1, j)]){
                q = append(q, []int{i-1, j})
                visited[fmt.Sprintf("%d %d", i-1, j)] = true
            }
        }
        if j < len(lines[i])-1 {

            ele := string(lines[i][j+1])
            if slices.Contains(curr[2], ele) && !(visited[fmt.Sprintf("%d %d", i, j+1)]){
                q = append(q, []int{i, j+1})
                visited[fmt.Sprintf("%d %d", i, j+1)] = true
            }
        }
        if i < len(lines)-1 {
            ele := string(lines[i+1][j])
            if slices.Contains(curr[3], ele) && !(visited[fmt.Sprintf("%d %d", i+1, j)]){
                q = append(q, []int{i+1, j})
                visited[fmt.Sprintf("%d %d", i+1, j)] = true
            }
        }


    }
    fmt.Println("result: ", result)
}










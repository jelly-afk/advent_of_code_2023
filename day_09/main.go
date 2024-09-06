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
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	result := 0
	for _, line := range lines {
		vals := strings.Split(line, " ")
		lst, err := strconv.Atoi(vals[len(vals)-1])
		if err != nil {
			log.Fatal(err)
		}
		last := []int{lst}
		for {
			var newArr []string


			for i, val := range vals{	
				if i == len(vals)-1{
					continue
				}
				curr, err := strconv.Atoi(val)
				if err != nil {
					log.Fatal(err)
				}
				next, err := strconv.Atoi(string(vals[i+1]))
				if err != nil {
					log.Fatal(err)
				}
				newArr = append(newArr, strconv.Itoa(next-curr))

				if i == len(vals)-2 {
					last = append(last, next-curr)
				}
			}
			n := newArr[0]
			same := true
			for _, s := range newArr {
				if n != s {
					same = false
					break
				}

			}
			vals = newArr
			if same {
				break
			}
		}
		res := 0

		for _, n := range last {
			res += n
		}	
		result += res

	}
	println("result: ", result)

}

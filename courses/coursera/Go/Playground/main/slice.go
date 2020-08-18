package main

import (
	"fmt"
	"strconv"
	"sort"
)


func main ()  {
	sli := make([]int, 3)
	idx := 0
	for {
		var input string
		fmt.Printf("Please enter a number: ")
		fmt.Scan(&input)
		if input == "X" {
			break
		} else {
			num, _ := strconv.Atoi(input)
			if idx >= 3 {
				sli = append(sli, num)
			} else {
				if sli[0] == 0 {
					sli[0] = num
				} else if sli[1] == 0 {
					sli[1] = num
				} else {
					sli[2] = num
				}
			}
			idx += 1
			sort.Ints(sli)
			fmt.Println(sli)
		}
	}
}
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough argument.")
		return
	}

	ags := os.Args[1:]

	for _, n := range ags {
		for _, ch := range n {
			if ch < '0' && ch > '9' {
				fmt.Println("Please enter numbers only")
				return
			}
		}
	}

	mainSlice := []int{}

	for _, s := range ags {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error converting argument to string")
			return
		}
		mainSlice = append(mainSlice, num)
	}
	secHalf := secondHalf(mainSlice)

	fmt.Printf("Original slice: %v\n", mainSlice)
	fmt.Printf("Second Half slice: %v\n", secHalf)

}

// Function secondHalf() takes in a slice of string and returns its second half.
func secondHalf(s []int) []int {
	ns := []int{}
	l := float64(len(s))
	// fmt.Println(math.Ceil(float64(5) / 2))
	ln := math.Ceil(l / 2)
	indx := l-ln -1
	for i, n := range s {
		if i > int(indx) {
			ns = append(ns, n)
		}
	}
	return ns
}

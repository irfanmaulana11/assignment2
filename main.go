package main

import (
	"fmt"
)

func main() {
	PrintName()
}

// PrintName ...
func PrintName() {
	var names1 = []string{"irfan", "Aulia", "San", "Wicak", "Iqbal"}
	var names2 = []string{"Giva", "Fahmi", "Rizki", "Billy", "Fajar"}
	names1 = append(names1, names2...)

	for _, n := range names1 {
		fmt.Println(n)
	}
}

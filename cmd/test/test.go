package main

import "fmt"


func main() {
	a := []int{1,2,3}
	b := make([]int, 0)
	b = append(b, 2,3,4)
	fmt.Println(a[:])
	fmt.Println(b[:])

	
}
package main

import "fmt"

func add(x int, y int) int {
	var sum int
	sum = x + y
	return sum
}

func swap(x, y string) (string, string) {
	k := x
	x = y
	y = k
	return x, y
}

func main(){
	fmt.Println(add(4,5))
	fmt.Println(swap("a", "b"))
}


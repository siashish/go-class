package main

import "fmt"

// int x, int y

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(32, 45))
	fmt.Println(sub(20, 10))
}

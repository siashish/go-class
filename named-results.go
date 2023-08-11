package main

import "fmt"

func split(sum int) (x int, y int) {
	x = sum * 4 / 9 // 17 * 4/9
	y = sum - x
	return
}

func split1(sum int) (x, y int) {
	x = sum * 4 / 9 // 17 * 4/9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
	fmt.Println(split1(20))
}

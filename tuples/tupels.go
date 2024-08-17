package main

import "fmt"

func powerSeries(a int) (int, int, error) {
	square := a * a
	cube := square * a

	return square, cube, nil
}

func main() {
	var square int
	var cube int
	square, cube, err := powerSeries(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Square ", square, "Cube", cube)
}

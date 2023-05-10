package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	//your code here!

	// Check if position is "left"
	if position == "left" {
		data = append([]int{newData}, data...)
		return data
	}


	// if position == "left" {
	// 	data = append([]int{newData}, data...)
	// 	return data
	// }

	// Check if position is "right"
	if position == "right" {
		data = append(data, newData)
		return data
	}

	// Otherwise, return the original slice
	return data
}

func main() {
	// Example usage
	data := []int{1, 2, 3}
	newData := 4
	position := "right"
	result := AddElement(data, newData, position)
	fmt.Println(result) // Add elemen baru
}

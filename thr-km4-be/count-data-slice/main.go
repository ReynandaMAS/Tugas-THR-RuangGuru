package main

import "fmt"

func howManyElements(data []any) int {
	//your code here!

	return len(data)
}

func main() {
	// contoh pemanggilan fungsi
	data := []interface{}{"hello", 123, true}
	count := howManyElements(data)
	fmt.Println(count) // output: 3
}

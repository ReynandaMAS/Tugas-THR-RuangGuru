package main

func removeLeftRight(arr []any, position string) []any {
	//yout code here!

	var result []interface{}
    midIndex := len(arr) / 2
    
    if position == "left" {
        result = arr[midIndex:]
    } else if position == "right" {
        result = arr[:midIndex]
    }
    
    return result
}

func main() {}

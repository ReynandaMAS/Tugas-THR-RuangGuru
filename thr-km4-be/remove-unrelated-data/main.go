package main

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	//your code here
	delete(dataMap, key)
	return dataMap

}

func main() {}

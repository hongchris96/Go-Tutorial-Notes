package main

import (
	"fmt"
)

func main() {

	// Map
	fmt.Println("Map:")

	// literal syntax map[ keyType ]valueType{keyValue pairs}
	sampleMap := map[string]int{
		"Dio":   1234521,
		"Jojo":  3523521,
		"Wagon": 1234323,
		"Smoky": 4538903,
		"Lisa":  1122323,
	}
	fmt.Println(sampleMap)

	// array is a valid key type but slice is not
	mapWithArrayKey := map[[3]int]string{}
	fmt.Println(mapWithArrayKey)

	// making a map with make(map[ keyType ]valueType)
	sampleMap2 := make(map[string]int)
	fmt.Println(sampleMap2)
}

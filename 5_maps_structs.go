package main

import (
	"fmt"
	"reflect"
)

// Define a type here with struct
// capitalized name so can be exported (variable rule)
// other files and see Doctor struct but can't access what's inside because they are lowercase
// inside will be accessible if capitalized too
type Doctor struct {
	// field name and type
	number     int
	doctorName string
	companions []string
}

// Embedding
// Composition (similar to inheritance)
type Animal struct {
	// Tags uses backticks around, inside are key value pairs of chef's choice separated by space (key: "value")
	Name   string `required max:"100"` // it's required and max length of 100
	Origin string
}

// Bird has traits of Animal but is NOT an Animal because it's not inheritance
type Bird struct {
	Animal // embedding animal struct into bird struct
	Speed  float32
	CanFly bool
}

func main() {

	// Map
	// key value pairs are stored unordered
	fmt.Println("Map:")

	// literal syntax map[ keyType ]valueType{keyValue pairs}
	// after declaring the types, all the keys must have the same type, all the values must have the same type
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

	// making a map with make(map[ keyType ]valueType, optional second int variable with no effect)
	sampleMap2 := make(map[string]int)
	fmt.Println(sampleMap2)

	// reading value of key in map
	fmt.Println(sampleMap["Dio"])
	// adding new key and value pair
	sampleMap["Diavolo"] = 7234908
	fmt.Println(sampleMap)

	// delete entry from the map
	delete(sampleMap, "Diavolo")
	fmt.Println(sampleMap)

	// reading value of key that doesn't exist
	fmt.Println(sampleMap["x"]) // returns 0
	// destructure keyed value to two variables
	xValue1, xValue2 := sampleMap["x"] // returns 0, false (first value will be zero if doesn't exist)
	fmt.Println(xValue1, xValue2)
	dioValue1, dioValue2 := sampleMap["Dio"] // returns 1234521, true (first value returns the value if exist)
	fmt.Println(dioValue1, dioValue2)

	// length of map
	fmt.Println(len(sampleMap))

	// mutable reference
	copySample := sampleMap
	delete(copySample, "Lisa")
	fmt.Println(copySample)
	fmt.Println(sampleMap)

	// --------------------------------------------------------------------------

	// Struct (see above function main)
	// gather information together that are related to one concept
	// can have different types of field name or key
	fmt.Println("Struct:")
	// literal syntax declaration
	house := Doctor{
		number:     1000,
		doctorName: "Gregory House",
		// when setting something to a slice, must show the type for initializing
		companions: []string{
			"Robert Chase",
			"Eric Foreman",
			"Allison Cameron",
		},
	}
	fmt.Println(house)
	fmt.Println(house.companions) // grab value for the field with dot

	// positional syntax (not encouraged: when you alter the Struct backbone, position changes)
	// only used for short-lived struct (e.g. anonymous struct) where you probably won't alter the backbone
	hobo := Doctor{
		10,
		"Madao",
		[]string{
			"Gintoki",
			"Hijigada",
		},
	}
	fmt.Println(hobo)

	// Anonymouse struct
	// type created during variable declaration, cannot be used elsewhere
	// struct{definition}{initialization} struct{fieldname type}{fieldname: value}
	anon := struct{ name string }{name: "Guy"}
	fmt.Println(anon)

	// copying struct
	// independent dataset, another is a copy of the struct
	another := anon
	another.name = "Kakashi"
	fmt.Println(anon)
	fmt.Println(another)
	// pointer to the original struct
	anotherAnother := &anon
	anotherAnother.name = "Lee"
	fmt.Println(anon)
	fmt.Println(anotherAnother)

	// Embedding (see above function main)
	// Composition (similar to inheritance)
	wee := Bird{}
	wee.Name = "Emu" // Bird struct has field name from Animal, kinda like inheritance
	wee.Origin = "Australia"
	wee.Speed = 48
	wee.CanFly = false
	fmt.Println(wee.Name)

	// literal syntax of embedded struct
	// needs to be aware of embedding
	pee := Bird{
		// need to explicitly talk about Animal struct
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		Speed:  48,
		CanFly: false,
	}
	fmt.Println(pee.Name)

	// Tags (see above function main)
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field)
}

package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map[keyType]valueType
	// mapVariable := make(map[keyType]valueType)
	// using a map literal

	/*
		mapVariable = map[keyType]valueType{
			key1: value1,
			key2: value2,
		}
	*/

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["priyanshu"] = 10
	myMap["rounak"] = 923
	myMap["avnish"] = 49
	fmt.Println(myMap)
	fmt.Println(myMap["priyanshu"]) // if the key does not exist it returns default value of valueType
	delete(myMap, "priyanshu")
	fmt.Println(myMap)

	// clear(myMap) // this will wipe all the data in map
	// fmt.Println(myMap)

	_, ok := myMap["Priyanshu"]
	fmt.Println(ok) // unknownvalue is boolean value, meaning if the key exist or not

	myMap2 := map[string]int{
		"a": 1,
		"b": 2,
	}

	myMap3 := map[string]int{
		"a": 1,
		"b": 2,
	}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("true")
	}

	for _, v := range myMap2 {
		fmt.Println(v)
	}

	// zero value of map is nil

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value")
	} else {
		fmt.Println("dadaaadaadad")
	}

	val := myMap4["key"]
	fmt.Println(val)

	// myMap4["key"] = "value"
	// fmt.Println(myMap4)

	myMap4 = make(map[string]string)
	myMap4["key"] = "value"
	fmt.Println(myMap4)

	fmt.Println(len(myMap))

	// nested maps (multi-dimentional maps)

	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4

	fmt.Println(myMap5)
}

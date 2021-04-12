package main

import (
	"fmt"
	"time"
)

type Camera struct {
	name        string
	focalLength int
}

func initCamera(focalLength int, name string) *Camera {
	cam := Camera{name: name}
	cam.focalLength = focalLength
	return &cam
}

func main() {
	fmt.Println(time.Now().Date())
	tempState := func(i interface{}) {
		switch tp := i.(type) {
		case bool:
			fmt.Println("Bool")
		case int:
			fmt.Println("Int")
		default:
			fmt.Printf("Undefined %T\n", tp)
		}
	}
	tempState(true)

	var tempArray [5][3]int
	tempArray[3][1] = 100
	fmt.Println("Arr: ", tempArray)
	for _, row := range tempArray {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}

	// Slices
	var myString string = "Slices can be composed into multi-dimensional data structures."
	fmt.Println(string(myString[2:10]))
	testStr := make([]string, 0)
	testStr = append(testStr, "a")
	testStr2 := make([]string, len(testStr))
	copy(testStr2, testStr)
	fmt.Println("testString:", testStr)
	fmt.Println("testString2:", testStr2)

	// Maps
	testMap := make(map[string]int)
	testMap["test1"] = 101
	testMap["test2"] = 102
	testMap["test3"] = 103
	fmt.Println("Map size: ", mapTester(testMap))

	// Variadic functions
	test_str := []string{"Hello", "world", "!", "Hi!"}
	sumStr(test_str...)

	// Closures
	nextInt := intSequence()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Pointers
	value := 45
	testPointer(&value)
	fmt.Println(value, " ", &value)

	// Struct
	cam := initCamera(8, "IRZ")
	fmt.Println(cam.focalLength, " ", cam.name)
}

func mapTester(testMap map[string]int) int {
	delete(testMap, "test2")
	_, prs := testMap["test1"]
	fmt.Println("prs:", prs)
	fmt.Println(testMap)
	return len(testMap)
}

func sumStr(str ...string) {
	fmt.Println(str, " ")
	fmt.Println(len(str))
}

func intSequence() func() int {
	itr := 0
	return func() int {
		itr++
		return itr
	}
}

func testPointer(i *int) {
	*i++
}

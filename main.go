package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func sayHelloGreeting(name string) {
	fmt.Printf("Good Morning %v ! \n", name)
}
func sayGoodbyeFreeting(name string) {
	fmt.Printf("Goodbye %v ! \n", name)
}

func cycleNames(names []string, f func(string)) {
	for _, name := range names {
		f(name)
	}
}

func getInitials(n string) (string, string) {
	capatalised := strings.ToUpper(n)
	names := strings.Split(capatalised, " ")
	var initials []string
	for _, name := range names {
		initials = append(initials, name[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[len(initials)-1]
	}
	return initials[0], "_"
}

func cirecleArea(r float64) float64 {
	return math.Pi * r * r
}

func updateName(name *string) {
	*name = "updatedName"
}

func main() {

	fmt.Println("Hello, world")
	var name1 string = "Tamara"
	name2 := "Tenzin"
	var name3 string
	name3 = "hope"
	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)

	age := 35
	name := "Sean"

	// Print
	fmt.Println("Hello World!")
	fmt.Print("Hello! \n")

	// Printf with variables
	fmt.Printf("Hello! My age is %v and my name is %v. \n", age, name)
	fmt.Printf("Age is of type: %T and name is of type: %T. \n", age, name)
	fmt.Printf("You scored %f points! \n", 24.93)
	fmt.Printf("You scored %0.2f points! \n", 24.93)
	fmt.Printf("You scored %0.1f points! \n", 24.97)

	// Sprintf to save formatted strings as a new
	var srt = fmt.Sprintf("Hello! My age is %v and my name is %v. \n", age, name)
	fmt.Println(srt)

	// Arrays cannot change the length of an array in go
	var nums [5]int
	ages := [3]int{34, 87, 21}
	names := [4]string{"name1", "name2", "name3", "name4"}

	fmt.Println(nums, ages, names)
	fmt.Println(nums, len(nums))

	names[2] = "newName"
	fmt.Println(names)

	// Slices (uses arrays unbder the hood) but can change the length
	var scores = []int{100, 50, 50, 90}

	fmt.Println(scores)
	newScores := append(scores, 45)
	fmt.Println(newScores)

	// Slice Ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	appendedRange := append(rangeOne, "AppendedName")

	fmt.Println(appendedRange)

	// String modification, Strings package
	greeting := "Hello there friends!"

	fmt.Println(strings.Contains(greeting, "hi!"))           // returns true/ false
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "hi")) // does not modify original string
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.ToLower(greeting))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, " "))

	// Sorts Package
	ages2 := []int{1, 56, 42, 75, 34}
	sort.Ints(ages2) // modifies original slice

	fmt.Println(ages2)

	index := sort.SearchInts(ages2, 75)
	fmt.Println(index)

	names2 := []string{"Tim", "Sara", "Happy", "Lama", "Zed"}
	sort.Strings(names2)
	fmt.Println(names2)
	fmt.Println(sort.SearchStrings(names2, "Sara"))

	// LOOPS for fun for all of us
	x := 0
	for x < 5 {
		fmt.Println("the value of x is:", x)
		x++
	}
	for i := 8; i < 11; i++ {
		fmt.Println("The value of i is:", i)
	}
	names3 := []string{"Tim", "Sara", "Happy", "Lama", "Zed"}
	for i := 0; i < len(names3); i++ {
		fmt.Println(names3[i])
	}
	for index, value := range names3 {
		fmt.Printf("The index is %v and the name is %v \n", index, value)
	}
	for _, value := range names3 {
		fmt.Printf("The name is %v \n", value)
	}

	// CONDITIONALS
	age7 := 30
	if age7 < 30 {
		fmt.Println("age is less than 30")
	} else if age7 < 40 {
		fmt.Println("age is more than 30  and less an 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	// Nest loop statement
	for i, v := range names3 {
		if i == 1 {
			fmt.Println("continuing at pos", i, v)
			continue
		}
		if i > 2 {
			fmt.Println("breaking at pos", i, v)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", i, v)
	}

	cycleNames([]string{"n1", "n2", "n3", "n4", "n5"}, sayHelloGreeting)
	cycleNames([]string{"n1", "n2", "n3", "n4", "n5"}, sayGoodbyeFreeting)

	a1 := cirecleArea((10.5))
	fmt.Printf("The area of 10.5 is %v. \n", a1)
	fmt.Printf("the radius of the circle us %0.3f. \n", a1)

	fn1, sn1 := getInitials("tifa lockheart")
	fn2, sn2 := getInitials("happy")
	fn3, sn3 := getInitials("happy lama damas")

	fmt.Println(fn1, sn1)
	fmt.Println(fn2, sn2)
	fmt.Println(fn3, sn3)

	// MAPS: all of the keys must be the same type and all of the values same type
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          8.99,
		"toffee pudding": 9.99,
	}
	fmt.Println(menu)
	fmt.Println(menu["soup"])

	// looping maps
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// ints as key type
	phonebook := map[int]string{
		1234567: "lemon",
		4535544: "monnn",
		2347672: "lenons",
	}
	fmt.Println(phonebook)
	phonebook[4535544] = "camel"

	fmt.Println(phonebook)
	fmt.Println(phonebook[1234567])
	for key, value := range phonebook {
		fmt.Println(key, value)
	}

	// //POINTERS

	mary := "Mary"
	m := &mary// create a pointer
	fmt.Println("name before update:", *m) // dereference pointer
	updateName(m)
	fmt.Println(*m)
	fmt.Println(m)
	fmt.Println("Mary vari:", mary)

	// STRUTS
	 

}

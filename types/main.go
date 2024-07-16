package main

import "fmt"

func main() {
	// boolean
	var isReady bool = true
	fmt.Println(isReady)

	// numeric: integer
	var age int = 30
	fmt.Println(age)

	// numeric: floating-point
	var pi float64 = 3.14
	fmt.Println(pi)

	// numeric: complex
	var c complex128 = 1 + 2i
	fmt.Println(c)

	// string
	var name string = "John"
	fmt.Println(name)

	// composite types: array
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	// composite types: slice
	var slice []int = []int{1, 2, 3}
	fmt.Println(slice)

	// composite types: map
	var m map[string]int = map[string]int{"one": 1, "two": 2}
	fmt.Println(m)

	// composite types: struct
	type Person struct {
		name string
		age  int
	}
	var person Person = Person{name: "Alice", age: 25}
	fmt.Println(person)

	// composite types: pointer
	var x int = 10
	var p *int = &x
	fmt.Println(*p)

	// composite types: function
	var f func(int, int) int = add
	fmt.Println(f(2, 3))

	// reference types: channel
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)

	// reference types: interface
	var s Speaker = Dog{}
	fmt.Println(s.Speak())

	// special types: rune (alias for int32)
	var r rune = 'a'
	fmt.Println(r)

	// special types: byte (alias for uint8)
	var b byte = 'a'
	fmt.Println(b)

	// type aliases
	type MyInt int
	var z MyInt = 10
	fmt.Println(z)
}

func add(a int, b int) int {
	return a + b
}

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

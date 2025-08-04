package main

import "fmt"

func main() {
	var arith int
	var a int
	var b int
	fmt.Println("👋 Hi there! Ready to do some quick math?")
	fmt.Println("What arithmetic do you want to perform today?")
	fmt.Println("{1 for Add} {2 for Subtract} {3 for Multiply} {4 for Division}?")
	fmt.Scanln(&arith)
	fmt.Println("Add the two numbers you want to work with")
	fmt.Println("First Number")
	fmt.Scanln(&a)
	fmt.Println("Second Number")
	fmt.Scanln(&b)
	

	func add(a int, b int)(int){
		return a + b
	}



}

package main

import "fmt"

func add(a int, b int)(int){
		return a + b
	}

	func minus(a int, b int)(int){
		if a > b {
			return a - b
		}else{
			return b - a
		}
	}

	func multiply(a int, b int)(int){
		return a * b
	}

	func division(a int, b int)(float64){
		return float64(a) / float64(b)
	}

func main() {
	var arith int
	var a int
	var b int

	for{

		fmt.Println("👋 Hi there! Ready to do some quick math?")
		fmt.Println("What arithmetic do you want to perform today?")
		fmt.Println("{1 for Add} {2 for Subtract} {3 for Multiply} {4 for Division}?")
		fmt.Scanln(&arith)

		
		if arith >= 1 && arith <= 4 {
			break
		}
		fmt.Println("Please enter a number between 1 and 4.")
	}

	
	fmt.Println("Add the two numbers you want to work with")
	fmt.Println("First Number")
	fmt.Scanln(&a)
	fmt.Println("Second Number")
	fmt.Scanln(&b)
	

	fmt.Println("The Ans is")


	switch arith {
	case 1:
		ans := add(a,b)
		fmt.Println(ans)

	case 2:
		ans := minus(a,b)
		fmt.Println(ans)

	case 3:
		ans := multiply(a,b)
		fmt.Println(ans)

	case 4:
		ans := division(a,b)
		fmt.Println(ans)
	}
}

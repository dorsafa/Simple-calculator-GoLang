package main

import "fmt"

func main() {
	a := 0
	b := 0
	operation := ""

	fmt.Println("Enter number 1 :")
	fmt.Scan(&a)
	fmt.Println("Enter number 2 :")
	fmt.Scan(&b)
	fmt.Println("Enter the operation : add | multiply | divide | substract | fibonacci")
	fmt.Scan(&operation)


	switch operation {
	case "add":
		fmt.Println(a," + ",b , "=", add(a,b))
	case "multiply":
		fmt.Println(a," * ",b , "=", multiply(a,b))
	case "divide":
		fmt.Println(a," / ",b , "=", divide(a,b))
	case "substract":
		fmt.Println(a," - ",b , "=", substract(a,b))
	case "fibonacci" :
		fmt.Println("fibonacci de ", a, " : ", fibonacci(a))
	}
}

func add(a int, b int) int {
	return a+b
}

func multiply(a int, b int) int {
	return a*b
}
func divide (a int, b int) string {
	return "a/b"
}
func substract(a int, b int) int {
	return a - b
}
func fibonacci(a int) int {
	if a<=1 {return 1}
	return fibonacci(a-1) + fibonacci(a-2)
}
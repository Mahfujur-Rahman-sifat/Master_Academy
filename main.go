package main

import "fmt"

func main() {

	fmt.Println("Enter your age")
	var age int
	fmt.Scanf("%d", &age)
	if age < 18 {
		fmt.Println("Children")
	} else if age >= 18 {
		fmt.Println("Adult")

	}
}

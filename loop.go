package main

import "fmt"

func main() {
	var age int

	fmt.Println("Enter your age")
	fmt.Scanf("%d", &age)

	if age < 3 {
		fmt.Println("Infant")
	} else if age > 2 && age < 13 {
		fmt.Println("Children")
	} else if age > 12 && age <= 19 {
		fmt.Println("Teen age")

	}

}

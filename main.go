package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Enter your name =")
	fmt.Scanf("%s", &name)
	fmt.Println("Enter your age =")
	fmt.Scanf("%d", &age)

	fmt.Printf("Your name is %s and age %d", name, age)

}

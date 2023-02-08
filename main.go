package main

import "fmt"

func main() {
	var student [3]string
	student[0] = "sifat"
	student[1] = "muhammad"
	student[2] = "mahfuj"

	fmt.Println(student)
	fmt.Println("string length =", len(student))

}

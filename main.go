package main

import "fmt"

func main() {

	BookInfo := struct {
		title   string
		authore string
		ISBN    int32
		price   float32
	}{
		title:   "The GO",
		authore: "Sifat",
		ISBN:    6895418,
		price:   125,
	}
	fmt.Println(BookInfo)

}

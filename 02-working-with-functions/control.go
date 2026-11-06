package main

import "fmt"

func welcome() {
	fmt.Println("Welcome to the program!")
}

func main() {
	welcome()
	fmt.Println("This is my function.")
	welcome()
}

func createMovie() {
	var title string
	var yearrelease int
	var movie float64

	fmt.Println("Enter the title of the movie:")
	fmt.Scanln(&title)
}

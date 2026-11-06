package main

import "fmt"

func welcome() {
	fmt.Println("Welcome to the program!")
}

func main() {
	welcome()
	fmt.Println("This is my function.")
	createMovie()
}

func createMovie() {
	var title string
	var yearRelease int
	var moviePrice float64

	fmt.Println("Enter the title of the movie:")
	fmt.Scanln(&title)
	fmt.Println("Enter the year of release:")
	fmt.Scanln(&yearRelease)
	fmt.Println("Enter the price of the movie:")
	fmt.Scanln(&moviePrice)

	fmt.Printf("Movie created: %s (%d) - $%.2f\n", title, yearRelease, moviePrice)
}

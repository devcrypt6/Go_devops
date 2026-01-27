package main

import "fmt"

func main() {
    // Ask for: name, age, city, favorite programming language
    // Display in a nice formatted box
    
    var name, city, favLang string
    var age int
    
    // Get user input
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Print("Enter your city: ")
	fmt.Scan(&city)

	fmt.Print("Enter your favorite programming language: ")
	fmt.Scan(&favLang)
	
	// Display formatted info card
	fmt.Println("==============================")
	fmt.Println("|     DEVELOPER INFO CARD      |")
	fmt.Println("==============================")
	fmt.Printf("Name:     %s\n", name) // %s is format specifier for string
	fmt.Printf("Age:      %d\n", age) // %d is format specifier for integer
	fmt.Printf("City:     %s\n", city)
	fmt.Printf("Language: %s\n", favLang)
	fmt.Println("==============================")
    
}

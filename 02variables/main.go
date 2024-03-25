package main

import "fmt"


func main() {
	var userName string = "John Doe"

	fmt.Printf("Type: %T \n", userName)

	var isLogged bool = true
	fmt.Printf("Type: %T \n", isLogged)
	fmt.Println(isLogged)

	var smallValue int = 456
	fmt.Printf("Type: %T \n", smallValue)
	fmt.Println(smallValue)

	var smallFloat float64 = 456.564646467014014010252368242478257878
	fmt.Printf("Type: %T \n", smallFloat)
	fmt.Println(smallFloat)

	//default values and aliases
	var senondValue int
	fmt.Println(senondValue)
	var senondString string
	fmt.Println(senondString)

	//implicit type
	var web = "www.google.com"
	fmt.Printf("Type: %T \n", web)
	fmt.Println(web)

	//no var keyword
	numOfuser := 100
	fmt.Printf("Type: %T \n", numOfuser)
	fmt.Println(numOfuser)
}

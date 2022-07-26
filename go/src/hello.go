package main

import "fmt"

func main() {
	name := "Ivo"
	var version float32 = 1.1
	fmt.Println("Hello mr ", name)
	fmt.Println("This program is in the version ", version)

	fmt.Println("1 - Start the monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("3 - Leave")

	var comand int
	fmt.Scan(&comand)

	if comand == 1 {
		fmt.Println("Monitoring")
	} else if comand == 2 {
		fmt.Println("Showing logs")
	} else if comand == 3 {
		fmt.Println("Leaving")
	} else {
		fmt.Println("Error")
	}
}

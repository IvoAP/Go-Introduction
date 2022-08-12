package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const number_of_monitoring = 5
const number_of_delay = 5 * time.Second

func main() {

	for {
		showIntroduction()
		showMenu()

		comand := readComand()

		// if comand == 1 {
		// 	fmt.Println("Monitoring")
		// } else if comand == 2 {
		// 	fmt.Println("Showing logs")
		// } else if comand == 3 {
		// 	fmt.Println("Leaving")
		// } else {
		// 	fmt.Println("Error")
		// }

		switch comand {
		case 1:
			beginMonitoring()
		case 2:
			fmt.Println("Showing logs")
		case 3:
			fmt.Println("Leaving")
			os.Exit(0)
		default:
			fmt.Println("I don't know this comand")
			os.Exit(-1)

		}
	}

}

func showIntroduction() {
	name := "Ivo"
	var version float32 = 1.1
	fmt.Println("Hello mr.", name)
	fmt.Println("This program is in the version ", version)
}

func readComand() int {
	var comand int
	fmt.Scan(&comand)
	return comand
}

func showMenu() {
	fmt.Println("1 - Start the monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("3 - Leave")
}

func beginMonitoring() {
	fmt.Println("Monitoring...")
	// using slice(type of array dynamic, because array is static)
	sites := []string{"https://random-status-code.herokuapp.com/"}
	sites = append(sites, "https://www.google.com.br/")
	sites = append(sites, "https://www.rocketseat.com.br/")
	for i := 0; i < number_of_monitoring; i++ {
		fmt.Println("Monitoring: ", i+1)
		for i, site := range sites {
			fmt.Println("Testing site ", i)
			testSite((site))
		}
		time.Sleep(number_of_delay)
	}

}

func testSite(site string) {
	response, _ := http.Get(site)
	if response.StatusCode == 200 {
		fmt.Println("Site: ", site, " sucess loading")
	} else {
		fmt.Println("Site: ", site, " error loading")
	}
}

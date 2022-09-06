package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const number_of_monitoring = 5
const number_of_delay = 5 * time.Second

func main() {

	showIntroduction()
	saveLog("site_false", false)

	for {

		showMenu()

		comand := readComand()

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

	sites := readSitesOfFile("sites.txt")
	for i := 0; i < number_of_monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testing site ", i, " : ", site)
			testSite(site)
		}

		time.Sleep(number_of_delay)
		fmt.Println("")
	}
	fmt.Println("")

}

func testSite(site string) {
	response, _ := http.Get(site)
	if response.StatusCode == 200 {
		saveLog(site, true)
	} else {
		saveLog(site, false)
	}
}

func readSitesOfFile(name_file string) []string {

	var sites []string
	file, error := os.Open(name_file)
	// file, error := ioutil.ReadFile("sites.txt")

	if error != nil {
		fmt.Println("There is an Error", error)
	}

	reader := bufio.NewReader(file)

	for {

		line, error := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		// fmt.Println(line)
		if error == io.EOF {
			break
		}

	}

	file.Close()

	return sites
}

func saveLog(site string, status bool) {
	file, error := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if error != nil {
		fmt.Println("There is an error: ", error)
	}
	file.WriteString(site + " online " + strconv.FormatBool(status) + "\n")

	sites := readSitesOfFile("log.txt")
	for i := range sites {
		fmt.Println(sites[i])
	}
	file.Close()
}

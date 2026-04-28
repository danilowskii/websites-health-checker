package main

import (
	"fmt"
	"time"
	"websites-health-checker/websites"
)

const TimeSleep = 2

func main() {
	introduction()
	
	running := true
	for running {
		mainMenu()
		fmt.Print("Selected option: ")
		running = chooseOption()
		timeSleep(TimeSleep)
		}
	
}

// Structure

func introduction() {
	fmt.Println("                ┌──────────────────────────────┐")
	fmt.Println("                │     \033[1;32mHEALTH CHECKER v1.0\033[0m      │")
	fmt.Println("                └──────────────────────────────┘")
	fmt.Println("")
}

func mainMenu() {
	
	fmt.Println("========================= MENU ============================")
	fmt.Println("1 - Check websites")
	fmt.Println("2 - Consult log")
	fmt.Println("3 - Add a new website")
	fmt.Println("4 - List websites")
	fmt.Println("5 - Delete a website")
	fmt.Println("0 - Exit program")
	fmt.Println("===========================================================")
}

func chooseOption() bool {
	var userChoice int
	_, err := fmt.Scan(&userChoice)
	
	if err != nil {
	fmt.Println("ERROR - Invalid option. Wait 3 seconds...")
	var discard string
	fmt.Scanln(&discard)
	timeSleep(TimeSleep)
	return true
}

	switch userChoice {
	case 1:
		websites.ClearLogs()
		fmt.Println("===========================================================")
		fmt.Println("1 - Monitoring your websites...")
		fmt.Println("")
		websites.Check()
		fmt.Println("")
		fmt.Println("")
	case 2:
		fmt.Println("===========================================================")
		fmt.Println("2 - Checking your last log...")
		fmt.Println("")
		response := websites.ReadLogs()
		for i := 0; i < len(response); i++ {
			fmt.Println(response[i])
		}
		fmt.Println("")
	case 3:
		fmt.Println("===========================================================")
		fmt.Println("3 - Adding website to list...")
		fmt.Println("")
		fmt.Printf("Enter a website: ")
		var websiteUrl string
		_, err := fmt.Scan(&websiteUrl)
		if err!= nil {
			fmt.Println("ERROR: ", err)
		}
		websites.AddWebsite(websiteUrl)
	case 4:
		fmt.Println("===========================================================")
		fmt.Println("4 - Listing all wesbites...")
		fmt.Println("")
		websites.ListAllWebsites()
	case 5:
		fmt.Println("===========================================================")
		fmt.Println("3 - Deleting website...")
		fmt.Println("")
		fmt.Printf("Enter a website ID: ")
		var urlId string
		_, err := fmt.Scan(&urlId)
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
		websites.DeleteWebsiteById(urlId)
	case 0:
		fmt.Println("===========================================================")
		fmt.Println("0 - Exiting program...")
		footer()
		fmt.Println("Press ENTER to close...")
		fmt.Scanln()
		return false
	default:
		fmt.Println("===========================================================")
		fmt.Println("ERROR - Invalid option. Wait 3 seconds...")
	}
	return true
}

func footer() {
	fmt.Println("")
	fmt.Println("                ┌──────────────────────────────┐")
	fmt.Println("                │       \033[1;32mCONNECTION CLOSED\033[0m      │")
	fmt.Println("                └──────────────────────────────┘")
	fmt.Println("")
	fmt.Println("                       \033[2mdeveloped by: github.com/danilowskii\033[0m")
	fmt.Println("")
}

// Utilities

func timeSleep(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
}
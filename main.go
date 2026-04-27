package main

import (
	"fmt"
	"health-checker-website/websites"

	"time"
)

const TimeSleep = 3

func main() {
	introduction()
	
	running := true
	for running {
		mainMenu()
		fmt.Print("Selected option: ")
		running = chooseOption()
		timeSleep(TimeSleep)
		}
	footer()
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
	fmt.Println("0 - Exit program")
	fmt.Println("===========================================================")
}

func chooseOption() bool {
	var userChoice int
	_, err := fmt.Scan(&userChoice)
	
	if err != nil {
		fmt.Println("===========================================================")
		fmt.Println("ERROR - Invalid option. Wait 3 seconds...")
		timeSleep(TimeSleep)
		return true
	}

	switch userChoice {
	case 1:
		fmt.Println("===========================================================")
		fmt.Println("1 - Monitoring your websites...")
		fmt.Println("")
		websites.Check()
		fmt.Println("")
		fmt.Println("")
	case 2:
		fmt.Println("===========================================================")
		fmt.Println("2 - Checking your last log...")
	case 0:
		fmt.Println("===========================================================")
		fmt.Println("0 - Exiting program...")
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
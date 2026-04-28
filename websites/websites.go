package websites

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

func Check() []string {

	websitesList, err := readFile()

	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	
	for idx, site := range websitesList {
		result, err := http.Get(site)
		if err != nil {
			fmt.Printf("%d - %-50s STATUS: %d %s\n", idx + 1, site, 404, "\033[31m- UNHEALTHY\033[0m")
			writeLogs(site, 404)
			continue
		}
		converted := strconv.Itoa(result.StatusCode)

		if strings.HasPrefix(converted, "2") {
			fmt.Printf("%d - %-50s STATUS: %d %s\n", idx + 1, site, result.StatusCode, "\033[32m- HEALTHY\033[0m")
			writeLogs(site, result.StatusCode)
		} else if strings.HasPrefix(converted, "9") {
			fmt.Printf("%d - %-50s STATUS: %d %s\n", idx + 1, site, result.StatusCode, "\033[33m- UNHAUTHORIZED\033[0m")
			writeLogs(site, result.StatusCode)
		} else {
			fmt.Printf("%d - %-50s STATUS: %d %s\n", idx + 1, site, result.StatusCode, "\033[33m- UNHEALTHY\033[0m")
			writeLogs(site, result.StatusCode)
		}
	}
	return websitesList
}

func readFile() ([]string, error) {
	var websites []string
	file, err := os.Open("websites/websites.txt ")

	if err != nil {
		return nil, err
	}

	readFile := bufio.NewReader(file)

	for {
		row, err := readFile.ReadString('\n')
		row = strings.TrimSpace(row)

		websites = append(websites, row)

		if err == io.EOF {
			break 
		}
	}

	file.Close()
	return websites, err
}

func writeLogs(site string, statusCode int) {
	file, err := os.OpenFile("websites/logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	color := "\033[31m- UNHEALTHY\033[0m"

	if statusCode >= 200 && statusCode < 300 {
		color = "\033[32m- HEALTHY\033[0m"
	} else if statusCode == 999 {
		color = "\033[33m- UNAUTHORIZED\033[0m"
	}

	file.WriteString(fmt.Sprintf(
	"%s - %s STATUS: %d %s\n",
	time.Now().Format("02/Jan/2006 15:04:05"),
	site,
	statusCode,
	color,
))
	
	
	file.Close()
}

func ReadLogs() []string {
	var websites []string
	file, err := os.Open("websites/logs.txt")

	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	readFile := bufio.NewReader(file)
	for {
		row, err := readFile.ReadString('\n')
		if err == io.EOF {
			break
		}

		row = strings.TrimSpace(row)
		websites = append(websites, row)
	}
	file.Close()
	return websites
}

func DeleteLogs() {
	os.Remove("websites/logs.txt")
}

package websites

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return nil, err
	}
	path := filepath.Join(home, "websites.txt")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, _ := os.Create(path)
		file.WriteString("https://google.com\n")
		file.Close()
	}
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)

		if row != "" {
			websites = append(websites, row)
		}

		if err == io.EOF {
			break 
		}
		if err != nil {
			return nil, err
		}
	}
	return websites, nil
}

func writeLogs(site string, statusCode int) {
	home, err := os.UserHomeDir()
	if err != nil {
	fmt.Println("ERROR:", err)
	return 
	}
	path := filepath.Join(home, "logs.txt")

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	defer file.Close()

	color := "\033[31m- UNHEALTHY\033[0m"

	if statusCode >= 200 && statusCode < 300 {
		color = "\033[32m- HEALTHY\033[0m"
	} else if statusCode == 999 {
		color = "\033[33m- UNAUTHORIZED\033[0m"
	}

	file.WriteString(fmt.Sprintf(
	"%s - %-45s STATUS: %d %s\n",
	time.Now().Format("02/Jan/2006 15:04:05"),
	site,
	statusCode,
	color,
))
}

func ReadLogs() []string {
	var logs []string
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return logs
	}

	path := filepath.Join(home, "logs.txt")
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return logs
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)
		if row != "" {
			logs = append(logs, row)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR reading logs:", err)
			break
		}
	}
	return logs
}

func ClearLogs() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	path := filepath.Join(home, "logs.txt")

	file, err := os.Create(path) 
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()
}

func AddWebsite(websiteUrl string) {
	u, err := url.ParseRequestURI(strings.TrimSpace(websiteUrl))
	if err != nil || u.Scheme == "" || u.Host == "" {
		fmt.Println("Invalid URL. Recommended url structure: 'https://google.com'. ")
		return
	}
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	path := filepath.Join(home, "websites.txt")
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	defer file.Close()

	file.WriteString(websiteUrl + "\n")
	fmt.Println("")
	fmt.Println("Website was successfully added!")
	fmt.Println("")
}

func ListAllWebsites()  {
	websites, err := readFile()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return 
	}
	if len(websites) == 0 {
		fmt.Println("No websites found.")
		return
	}
	for i, site := range websites {
		fmt.Printf("%d - %s\n", i+1, site)
	}
	fmt.Println("")
}

func DeleteWebsiteById(urlId string) {
	converted, err := strconv.Atoi(urlId)
	if err != nil {
		fmt.Printf("ERROR: ID '%s' isn't valid. Please enter a valid ID. \n", urlId)
		return
	}	
	websites, err := readFile()
	if err!= nil {
		fmt.Println("ERROR: ", err)
		return 
	}
	if len(websites) == 0 {
		fmt.Println("No websites found.")
		return
	}
	index := converted - 1
	if index < 0 || index >= len(websites) {
		fmt.Println("ERROR: ID out of range.")
		return
	}
	websites = append(websites[:index], websites[index+1:]...)
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, "websites.txt")

	file, err := os.Create(path) // sobrescreve tudo
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	for _, site := range websites {
		file.WriteString(site + "\n")
	}
	
	fmt.Println("")
	fmt.Printf("Website ID %d was successfully deleted! \n", converted)
	fmt.Println("")
	

}
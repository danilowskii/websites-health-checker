package websites

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Check() []string {
	var WebsiteList = []string{
		"https://www.linkedin.com/in/paivadanilo/",
		"https://github.com/danilowskii/",
		"https://portfolio-danilo-two.vercel.app/pt",
		"https://portfolio-daniloercel.app/pt",
	}

	for idx, value := range WebsiteList {
		result, err := http.Get(value)
		if err != nil {
			fmt.Printf("%d - %-50s STATUS: %d %s\n", idx + 1, value, 404, "\033[31m- UNHEALTHY\033[0m")
			continue
		}
		converted := strconv.Itoa(result.StatusCode)

		

		if strings.HasPrefix(converted, "2") {
			fmt.Printf("%d - %-50s STATUS: %d %s\n", idx + 1, value, result.StatusCode, "\033[32m- HEALTHY\033[0m",
)
		} else if strings.HasPrefix(converted, "9") {
			fmt.Printf("%d - %-50s STATUS: %d %s\n", idx + 1, value, result.StatusCode, "\033[33m- UNHAUTHORIZED\033[0m")
		} else {
			fmt.Printf("%d - %-50s STATUS: %d %s\n", idx + 1, value, result.StatusCode, "\033[33m- UNHEALTHY\033[0m")
		}

		
	}
	return WebsiteList
}

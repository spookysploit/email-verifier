package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Enter email(s) (comma-separated):")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()

	if len(data) == 0 {
		log.Fatal("Error: empty input given")
	}

	emails := strings.Split(data, ",")
	for _, email := range emails {
		email = strings.TrimSpace(email)
		if !checkFormat(email) {
			fmt.Printf("%v IS INVALID (bad format)\n", email)
			continue
		}

		domain := getDomainFromEmail(email)
		hasMX, hasSPF, hasDMARC := checkDomain(domain)
		fmt.Printf("Email: %v\nDomain: %s\n- MX: %v\n- SPF: %v\n- DMARC: %v\n\n",
			email, domain, hasMX, hasSPF, hasDMARC)
	}
}

func getDomainFromEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

func checkFormat(email string) bool {
	re := regexp.MustCompile(`\b[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}\b`)
	return re.MatchString(email)
}

func checkDomain(domain string) (bool, bool, bool) {
	var hasMX, hasSPF, hasDMARC bool

	mxRecords, err := net.LookupMX(domain)
	if err == nil && len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err == nil {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				break
			}
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err == nil {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				break
			}
		}
	}

	return hasMX, hasSPF, hasDMARC
}

package main

import (
	"fmt"
	"regexp"
)

func Is_valid_ip(ip string) bool {

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)

	if re.MatchString(ip) {
		fmt.Printf("%v is valid :: %v\n", ip, true)
		return true
	}
	fmt.Printf("%v is invalid :: %v\n" , ip, false)
	return false
}

func main() {
	Is_valid_ip("172.16.254.1")
	Is_valid_ip("172.216.254.1")
	Is_valid_ip("172.316.254.1") // in a second group number should be 1 - 254 not over to 316 
	Is_valid_ip("0.1.1.256")
	Is_valid_ip("1.1.1.1a")
	Is_valid_ip("1")
	Is_valid_ip("64.233.16.00")
	Is_valid_ip("7283728")
	Is_valid_ip("abc.co")
}

package assignment2

import (
	"regexp"
)

func Is_valid_ip(ip string) bool {

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)

	if re.MatchString(ip) {
		// fmt.Printf("%v is valid :: %v\n", ip, true)
		return true
	}
	// fmt.Printf("%v is invalid :: %v\n" , ip, false)
	return false
}
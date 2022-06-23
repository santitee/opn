package main

import (
	"fmt"
	"opn/assignment2"
	"opn/assignment4"
	"regexp"

	// "regexp"
	"strings"
)

type Student struct {
	name  string
	score int
	grade string
}

func main() {

	// assignment1
	fmt.Println(strings.Repeat("#", 50))
	fmt.Println("Assignment 1 Check First Number in Strings:")
	fmt.Println(strings.Repeat("#", 50))
	strs := []string{
		"Var----____123__int",
		"Q2q-q",
		"eef3243**s",
	}

	re := regexp.MustCompile(`[-]?\d{1}`)

	for _, str := range strs {
		fmt.Printf("String Pattern: %q\n", str) // Print Pattern
		fmt.Printf("First Number contains is: %v\n", re.FindAllString(str, 1))
		submatchall := re.FindAllString(str, 1)
		for _, element := range submatchall {
			fmt.Println(element)
		}
	}

	// assignment2
	fmt.Println(strings.Repeat("*", 50))
	// fmt.Println("++++++++++++++++++++++++++++++++++++++++:")
	fmt.Println("Assignment 2 Validated Check Ip Address:")
	fmt.Println(strings.Repeat("*", 50))
	var ips = []string{"172.16.254.1", "172.316.254.1", "0.1.1.256", "1.1.1.1a", "1", "64.233.16.00", "7283728", "abc.co"}

	for _, ip := range ips {
		fmt.Printf("ip: \"%s\"\nresult: %v\n", ip, assignment2.Is_valid_ip(ip))
	}

	// assignment4
	fmt.Println(strings.Repeat("=", 50))
	// fmt.Println("++++++++++++++++++++++++++++++++++++++++:")
	fmt.Println("Assignment 4 Grading Students scores:")
	fmt.Println(strings.Repeat("=", 50))

	// testing ...
	s := []Student{
		{name: "A", score: 90},
		{name: "B", score: 48},
		{name: "C", score: 63},
		{name: "D", score: 75},
		{name: "E", score: 70},
		{name: "F", score: 107},
		{name: "G", score: 2},
		{name: "H", score: 49},
	}

	for _, score := range s {
		grade := assignment4.CalStudentGrade(score.score)
		fmt.Printf("Student: %v  has score: %v = grade %s\n", score.name, score.score, grade)
	}
}

func SubMatchAll(strs []string) {

}

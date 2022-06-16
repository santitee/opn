package main

import (
	"fmt"
	"regexp"
)

func main() {

	str1 := "Var----____123__int"
	str2 := "Q2q-q"
	str3 := "eef3243**s"

	re := regexp.MustCompile(`[-]?\d{1}`)

	fmt.Printf("String Pattern: %v\n", str1) // Print Pattern
	fmt.Printf("First Number contains is: %v\n", re.FindAllString(str1, 1))
	submatchall1 := re.FindAllString(str1, 1)
	for _, element := range submatchall1 {
		fmt.Println(element)
	}

	fmt.Printf("String Pattern: %v\n", str2) // Print Pattern
	fmt.Printf("First Number contains is: %v\n", re.FindAllString(str2, 1))
	submatchall2 := re.FindAllString(str2, 1)
	for _, element := range submatchall2 {
		fmt.Println(element)
	}

	fmt.Printf("String Pattern: %v\n", str3) // Print Pattern
	fmt.Printf("First Number contains is: %v\n", re.FindAllString(str3, 1))
	submatchall3 := re.FindAllString(str3, 1)
	for _, element := range submatchall3 {
		fmt.Println(element)
	}
}

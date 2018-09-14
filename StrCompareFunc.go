package main

import (
	"strings"
	"fmt"
)

func StrEqual(a string, b string) bool{
	return a == b
}

func StrCompare(a string, b string) bool{
	return strings.Compare(a, b) == 0
}

func StrEqualFold(a string, b string) bool{
	return strings.EqualFold(a, b)
}

func StrContains(a string, b string) bool{
	return strings.Contains(a, b)
}

func StrCount(a string, b string) bool{
	return strings.Count(a, b) == 1
}

func StrHasPrefix(a string, b string) bool{
	return strings.HasPrefix(a, b)
}

func StrIndex(a string, b string) bool{
	return strings.Index(a, b) == 0
}

func main(){
	fmt.Printf("StrIndex : %t\n" , StrEqual(	"AC01", "AC01"))
	fmt.Printf("StrIndex : %t\n" , StrCompare("AC01", "AC01"))
	fmt.Printf("StrIndex : %t\n" , StrEqualFold("AC01", "AC01"))
	fmt.Printf("StrIndex : %t\n" , StrContains("AC01", "AC01"))
	fmt.Printf("StrIndex : %t\n" , StrCount("AC01", "AC01"))
	fmt.Printf("StrIndex : %t\n" , StrHasPrefix("AC01", "AC01"))
	fmt.Printf("StrIndex : %t\n" , StrIndex("AC01", "AC01"))
}
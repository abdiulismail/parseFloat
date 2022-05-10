package main

import (
	"fmt"
	"strconv"
)

func main() {

	str1 := "29.3"

	result, error1 := strconv.ParseFloat(str1, 64) //parses to float64

	if error1 == nil {
		fmt.Println("parse value is:", result)
	} else {
		fmt.Println("The string could not be converted to floating point number")
	}
}

package convert

import "fmt"

func DoSomethingWrong() bool {
	fmt.Println("Add code to convert numbers in words!")
	return true
}

func ConvertNumberToWord(number int) string {
	b := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	suffix := " dollars"

	var value string = ""
	if number < 10 {
		value = b[number]
	}

	return value + suffix
}

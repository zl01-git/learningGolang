// На вход подается целое число. Необходимо возвести в квадрат 
// каждую цифру числа и вывести получившееся число. 

// Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81.
// Дальше 1. Единица в квадрате - 1. В итоге получаем 811181

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func checkDigit(number string) bool {
	for _, digit := range number {
		if !unicode.IsDigit(digit) {
			return false
		}
	}
	return true
}

func getSquareDigit(number string) string {
	var (
		result string
		digit int
	)

	for _, elem := range number {
		digit, _ = strconv.Atoi(string(elem))
		result += strconv.Itoa(digit*digit)
	}
	return result
}

func main() {
	var numberString string

	// _, errorScan := fmt.Scan(&numberString)   // Получается не имеет смыла обрабатывать 
	// if errorScan != nil {					  // ошибки Scan если ждем стоку?
	// 	fmt.Println("Enter the string")		  // в любом же случае строка через stdin приходит?
	// } else {
	// 	fmt.Println(getSquareDigit(numberString))
	// }
	
											//  Только дополнительной функией проверять что прилетело?
											
	fmt.Scan(&numberString)        		 
	if !checkDigit(numberString) {
		fmt.Println("Not a string of digit passed to the input")
	} else {
		fmt.Println(getSquareDigit(numberString))
	}
}
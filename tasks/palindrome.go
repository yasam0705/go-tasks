package tasks

import "fmt"

func Palindrome() {
	var str string
	fmt.Scan(&str)

	var count int
	str_len := len(str)

	for i := 0; i < str_len/2; i++ {
		if str[i] == str[str_len-1-i] {
			count++
		} else {
			break
		}
	}

	if count == str_len/2 {
		fmt.Println("this word is polindrom")
	} else {
		fmt.Println("No polindrome")
	}
}

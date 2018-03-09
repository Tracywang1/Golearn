package main

import (
	"fmt"
)

func findString(inputStr string, findStr string) (sum int, err string) {
	num := 0
	fmt.Println(inputStr, findStr, len(inputStr), len(findStr))
	if len(inputStr) < len(findStr) {
		errstr := "输入字符串长度应大于查询字符串长度"
		return 0, errstr
	}
	for i := 0; i <= len(inputStr)-len(findStr); i++ {
		// fmt.Println(inputStr[i : i+len(findStr)])
		if findStr == string(inputStr[i:i+len(findStr)]) {
			num = num + 1
		}
	}
	return num, ""
}

func reverseString(inputStr string) string {
	finalstr := ""
	for i := len(inputStr) - 1; i >= 0; i-- {
		// fmt.Println(string(inputStr[i]))
		finalstr += string(inputStr[i])
	}
	return finalstr
}

func main() {
	fmt.Println(findString("aaaaddddaadd", "aadd"))
	fmt.Println(reverseString("abcdefghigklmn"))
}

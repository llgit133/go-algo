package main

import (
	"fmt"
	"log"
)

func reverseString(s []byte) []byte {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}

func main() {

	bytes := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Print(string(bytes))
	fmt.Println()
	reverseBytes := reverseString(bytes)
	fmt.Println(string(reverseBytes))

	log.Println("log 测试")

}

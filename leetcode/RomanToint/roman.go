package main

import (
	"fmt"
)

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	num := 0
	j := len(s) - 2

	for i := len(s) - 1; i >= 0; i-- {
		if i == 0 {
			num += m[string(s[i])]
			break
		}
		if j != -1 {
			switch string(s[i]) {
			case "I":
				num += m[string(s[i])]
				j--
			case "V":
				num += m[string(s[i])]
				if string(s[j]) == "I" {
					num -= m[string(s[j])]
					i--
					j--
				}
				j--
			case "X":
				num += m[string(s[i])]
				if string(s[j]) == "I" {
					num -= m[string(s[j])]
					i--
					j--
				}
				j--
			case "L":
				num += m[string(s[i])]
				if string(s[j]) == "X" {
					num -= m[string(s[j])]
					i--
					j--
				}
				j--
			case "C":
				num += m[string(s[i])]
				if string(s[j]) == "X" {
					num -= m[string(s[j])]
					i--
					j--
				}
				j--
			case "D":
				num += m[string(s[i])]
				if string(s[j]) == "C" {
					num -= m[string(s[j])]
					i--
					j--
				}
				j--
			case "M":
				num += m[string(s[i])]
				if string(s[j]) == "C" {
					num -= m[string(s[j])]
					i--
					j--
				}
				j--
			default:
				fmt.Println("No key")
			}
		}
	}
	return num
}

func main() {
	fmt.Println(romanToInt("IV"))
}

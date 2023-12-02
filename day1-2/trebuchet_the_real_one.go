package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func extract_ints(re *regexp.Regexp, text string) int64 {
	var new_text string
	var spelled_out = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var hold string
	for i := 0; i < len(text); i++  {
		v := text[i]
		hold += string(v)
		if unicode.IsDigit(rune(v)) {
			new_text += string(v)
			hold = ""
		}
		for j, value := range spelled_out {
			if strings.Contains(hold, value) {
				new_text += fmt.Sprint(j + 1)
				hold = ""
				i--
				break
			}
		}
	}
	fmt.Println(new_text)
	numbers := strings.Split(new_text,"")
	conversion, err := strconv.ParseInt((numbers[0] + numbers[len(numbers)-1]), 10, 64)
		if err != nil {
			fmt.Println(err)
		}
	return conversion
}

func main() {
	re := regexp.MustCompile(`[a-z]`)
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	var numeri []int64
	for scanner.Scan() {
		text := scanner.Text()
		conversione := extract_ints(re, text)
		numeri = append(numeri, conversione)
	}
	defer file.Close()
	var sum int64
	for i := 0; i < len(numeri); i++ {
		sum += numeri[i]
	}
	fmt.Println(sum)
}

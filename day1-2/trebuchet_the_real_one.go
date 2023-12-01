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
	spelled_out := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, number := range spelled_out {
		if strings.Contains(text, number) {
			text = strings.ReplaceAll(text, spelled_out[i], fmt.Sprint(i+1))
		}
	}
	fmt.Println(text)
	for _, char := range text {
		if unicode.IsDigit(char) {
			new_text = new_text + " " + string(char) + " "
		} else {
			new_text += string(char)
		}
	}
	new_text = strings.ReplaceAll(new_text, " ", "")
	splitted := strings.Split(new_text, " ")
	//var finally []string
	fmt.Println(splitted)
	/*for _, v := range splitted {
		if re.Match([]byte(v)) {
			for i, number := range spelled_out {
				if strings.Contains(string(v), number) {
					finally = append(finally, strconv.Itoa(i+1))
				}
			}
		} else {
			finally = append(finally, string(v))
		}
	}*/
	conversion, err := strconv.ParseInt((splitted[0] + splitted[len(splitted)-1]), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return conversion
}

func main() {
	re := regexp.MustCompile(`[a-z]`)
	file, err := os.Open("./input1.txt")
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
		bytes := re.ReplaceAll([]byte(text), []byte(""))
		myString := string(bytes)
		numbers := strings.Split(myString, "")
		conversion, err := strconv.ParseInt((numbers[0] + numbers[len(numbers)-1]), 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		numeri = append(numeri, conversion)
	}
	defer file.Close()
	var sum int64 = 0
	for i := 0; i < len(numeri); i++ {
		sum += numeri[i]
	}
	fmt.Println(sum)
}

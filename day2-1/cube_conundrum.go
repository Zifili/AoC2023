package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func check_game(text string, condition map[string]int) bool {
	extractions := strings.SplitAfter(text, ":")
	extractions = strings.Split(extractions[1], ";")
	n_pulls := len(extractions)
	for i := 0; i < n_pulls; i++ {
		pulls := strings.Split(extractions[i], ",")
		n_cubes := len(pulls)
		for j := 0; j < n_cubes; j++ {
			check := strings.Split(strings.Trim(pulls[j]," ")," ")
			number,_ := strconv.Atoi(check[0])
			if condition[check[1]] < number{
				fmt.Printf("impossibile perchè: %s > %d\n",check[0],condition[check[1]])
				return false
			}
		}
	}
	return true
}

func main() {
	var checker = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	game_id := 0
	result := 0
	for scanner.Scan() {
		game_id++
		text := scanner.Text()
		if check_game(text, checker) {
			result += game_id
		}
	}
	defer file.Close()
	fmt.Println(result)
}

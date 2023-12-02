package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate_fewer(text string) int {
	extractions := strings.SplitAfter(text, ":")
	extractions = strings.Split(extractions[1], ";")
	n_pulls := len(extractions)
	var power int
	var maxs = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for i := 0; i < n_pulls; i++ {
		pulls := strings.Split(extractions[i], ",")
		n_cubes := len(pulls)
		for j := 0; j < n_cubes; j++ {
			check := strings.Split(strings.Trim(pulls[j], " "), " ")
			curr, err := strconv.Atoi(check[0])
			if err!=nil{
				print(err)
			}
			if curr > maxs[check[1]]{
				//fmt.Printf("%d è maggiore di %d\n",curr,maxs[check[1]])
				maxs[check[1]] = curr
			} 
		}
	}
	for _,v := range maxs{
		fmt.Println(v)
	}
	power = maxs["red"] * maxs["green"] * maxs["blue"]
	fmt.Println(power)
	fmt.Println("---------------")
	return power
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		text := scanner.Text()
		result+=calculate_fewer(text)
	}
	defer file.Close()
	fmt.Println(result)
}

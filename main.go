package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dwprz/komerce-coding-test/task"
)

/*
example input:

 1. SortCharacters:
    - Sample Case
    - Next Case
 2. CalculateMinBuses:
    - number of families: 2 | number of members family: 14 4
	- number of families: 3 | number of members family: 3 1 3
	- number of families: 5 | number of members family: 3 3 3 3 2
	- number of families: 5 | number of members family: 1 2 4 3 3
	- number of families: 8 | number of members family: 2 3 4 4 2 1 3 1
	- number of families: 2 | number of members family: 3 9
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("select 1 for (sort characters) or 2 for (calculate min buses)!")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		task.SortCharacters(reader)
	case "2":
		task.CalculateMinBuses(reader)
	default:
		fmt.Println("invalid input")
	}
}

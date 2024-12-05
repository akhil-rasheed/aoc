package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	lists := strings.Split(string(content), "\n")

	const maxLength = 3000
	list1 := make([]int, 0, maxLength)
	list2 := make([]int, 0, maxLength)

	for _, element := range lists {
		if strings.TrimSpace(element) == "" {
			continue
		}

		indices := strings.Split(strings.TrimSpace(element), "   ")

		num1, _ := strconv.Atoi(indices[0])
		num2, _ := strconv.Atoi(indices[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	totalDistance := 0
	for i, _ := range list1 {
		totalDistance += int(math.Abs(float64(list2[i] - list1[i])))
	}

	fmt.Println(totalDistance)
}

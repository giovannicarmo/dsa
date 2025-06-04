package main

import (
	"fmt"
	"strings"
)

func recursiveSum(numbers []int) int {

	arrLen := len(numbers)

	if arrLen == 0 {
		return 0
	}

	if arrLen == 1 {
		return numbers[0]
	}

	return numbers[0] + recursiveSum(numbers[1:])
}

func arrayLength[T any](arr []T) int {

	if len(arr) == 0 {
		return 0
	}

	return 1 + arrayLength(arr[1:])
}

func greatestValueIn(numbers []int) int {

	arrLen := len(numbers)

	if arrLen == 0 {
		return 0
	}

	if arrLen == 1 {
		return numbers[0]
	}

	if arrLen == 2 {
		return max(numbers[0], numbers[1])
	}

	pivot := arrLen / 2

	return max(greatestValueIn(numbers[:pivot]), greatestValueIn(numbers[pivot:]))
}

func max(a int, b int) int {
	if a >= b {
		return a

	} else {
		return b
	}
}

// "Grooking algorithms solution"
func binaryIntSearch(numbers []int, queriedNumber int) int {

	lowestValuePosition := 0
	highesValuePosition := len(numbers) - 1

	for lowestValuePosition <= highesValuePosition {

		middle := (lowestValuePosition + highesValuePosition) / 2

		attempt := numbers[middle]

		if attempt == queriedNumber {
			return middle
		}

		if attempt > queriedNumber {
			highesValuePosition = middle - 1

		} else {
			lowestValuePosition = middle + 1
		}
	}

	return -1
}

// Recursive solution
func recursiveBinaryIntSearch(queriedNumber int, numbers []int) int {

	lowestValuePosition := 0
	highesValuePosition := len(numbers) - 1

	var search func(int, []int, int, int) int

	search = func(queriedNumber int, numbers []int, lowestValuePosition int, highesValuePosition int) int {

		if lowestValuePosition > highesValuePosition {
			return -1
		}

		middle := (lowestValuePosition + highesValuePosition) / 2

		attempt := numbers[middle]

		if attempt == queriedNumber {
			return middle
		}

		if attempt > queriedNumber {
			return search(queriedNumber, numbers, lowestValuePosition, (middle - 1))

		} else {
			return search(queriedNumber, numbers, (middle + 1), highesValuePosition)
		}
	}

	return search(queriedNumber, numbers, lowestValuePosition, highesValuePosition)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less []int
	var greater []int

	for _, item := range arr[1:] {
		if item <= pivot {
			less = append(less, item)
		} else {
			greater = append(greater, item)
		}
	}

	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func breadthFirstSearch() bool {

	graph := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}

	// var personIsSeller func(string) bool

	personIsSeller := func(person string) bool {
		return strings.HasPrefix(person, "t")
	}

	searchQueue := []string{"you"}
	verifiedPeople := map[string]bool{}

	for len(searchQueue) > 0 {

		person := searchQueue[0]
		searchQueue = searchQueue[1:]

		if !verifiedPeople[person] {

			if personIsSeller(person) {
				fmt.Println(person, "is a mango seller!")
				return true
			} else {

				for _, friend := range graph[person] {
					searchQueue = append(searchQueue, friend)
				}
				verifiedPeople[person] = true
			}
		}
	}

	return false
}

func main() {

	// arr := []int{15, 2, 99, 36, 47, 14, 59, 8, 12, 33, 36, 14}

	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// fruits := []string{"banana", "apple", "pasisonfruit", "strawberry", "pineapple"}

	// fmt.Print(recursiveBinaryIntSearch(4, arr))

	// fmt.Print(quickSort(arr))

	breathFirstSearch()
}

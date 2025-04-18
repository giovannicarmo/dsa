// package main

// import (
// 	"fmt"
// 	"math/rand/v2"
// )

// /* Utilities */
// func generate(length int) []int {
// 	return rand.Perm(length)
// }

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }

// /* END */

// /* [SOME] sort */
// /* END */

// /* Buble sort */
// func bubbleSort(arr []int) {
// 	fmt.Printf("Starting bubble sort algorithm...\n")
// 	fmt.Printf("Unsorted array: %v\n\n", arr)

// 	arrLen := len(arr)

// 	for i := range arrLen - 1 {
// 		swapped := false

// 		for currentPosition := range arrLen - i - 1 {

// 			nextArrPosition := currentPosition + 1

// 			if arr[currentPosition] > arr[nextArrPosition] {
// 				arr[currentPosition], arr[nextArrPosition] = arr[nextArrPosition], arr[currentPosition]

// 				swapped = true
// 			}
// 		}

// 		if !swapped {
// 			break
// 		}

// 		fmt.Printf("Sorted array after run nº %v: %v\n", i+1, arr)
// 	}

// 	fmt.Printf("\nSorted array: %v\n", arr)
// 	fmt.Printf("Bubble sort algorithm successfully finished!\n")
// }

// /*END*/

// /* Selection sort */

// func selectionSort(arr []int) {
// 	fmt.Printf("Starting selection sort algorithm...\n")
// 	fmt.Printf("Unsorted array: %v\n\n", arr)

// 	arrLen := len(arr)

// 	for i := range arrLen - 1 {
// 		lowestValuePosition := i

// 		for j := i + 1; j < arrLen; j++ {

// 			if arr[lowestValuePosition] > arr[j] {
// 				lowestValuePosition = j
// 			}
// 		}

// 		/* Swapping variables. (PERFORMATIC solution) */
// 		arr[i], arr[lowestValuePosition] = arr[lowestValuePosition], arr[i]
// 		/*END*/

// 		/* Moving all elements on the array and insert the current lowest level on its position. (NON-PERFORMATIC solution) */
// 		// lowestValue := arr[lowestValuePosition]

// 		// // Remove lowest value from its orginal position
// 		// for k := i; k < len(arr); k++ {
// 		// 	if lowestValue == arr[k] {
// 		// 		if lowestValuePosition+1 == arrLen {
// 		// 			arr = arr[:lowestValuePosition]
// 		// 		} else {
// 		// 			arr = append(arr[:lowestValuePosition], arr[lowestValuePosition+1:]...)
// 		// 		}
// 		// 	}
// 		// }

// 		// // Inserting lowest value into its new position
// 		// arr = append(arr[:i], append([]int{lowestValue}, arr[i:]...)...)
// 		/* END */

// 		fmt.Printf("Sorted array after run nº %v: %v\n", i+1, arr)
// 	}

// 	fmt.Printf("\nSorted array: %v\n", arr)
// 	fmt.Printf("Selection sort algorithm successfully finished!\n")
// }

// /* END */

// func insertionSort(arr []int) {
// 	fmt.Printf("Starting insertion sort algorithm...\n")
// 	fmt.Printf("Unsorted array: %v\n\n", arr)

// 	arrLen := len(arr)

// 	/* Solution moving only necessary elements (W3 schools solution) */

// 	for i := range arrLen {

// 		insertIndex := i
// 		currentVal := arr[i]

// 		for j := i - 1; j >= 0; j-- {

// 			if currentVal < arr[j] {

// 				arr[j+1] = arr[j]
// 				insertIndex = j

// 			} else {
// 				break
// 			}
// 		}

// 		arr[insertIndex] = currentVal
// 	}

// 	fmt.Printf("\nSorted array: %v\n", arr)

// 	/*END*/

// 	/* Solution using a copy (MY solution) */

// 	// newArr := []int{}

// 	// for i := range arrLen {

// 	// 	newArr = append(newArr, arr[i])
// 	// 	swaped := false

// 	// 	for j := i; j > 0; j-- {

// 	// 		if newArr[j] < newArr[j-1] {

// 	// 			newArr[j], newArr[j-1] = newArr[j-1], newArr[j]
// 	// 			swaped = true
// 	// 		}

// 	// 		if !swaped {
// 	// 			break
// 	// 		}

// 	// 	}

// 	// 	fmt.Printf("Sorted parital array after run nº %v: %v\n", i+1, newArr)
// 	// }

// 	// fmt.Printf("\nSorted array: %v\n", newArr)

// 	/* END */

// 	fmt.Printf("Insertion sort algorithm successfully finished!\n")
// }

// func quickSort(arr []int) []int {

// 	arrLen := len(arr)

// 	//base case
// 	if arrLen < 2 {
// 		return arr
// 	}

// 	var leftSubArray []int
// 	var rightSubArray []int

// 	pivot := arr[arrLen/2]

// 	for i := range arrLen {

// 		if arr[i] < pivot {

// 			leftSubArray = append(leftSubArray, arr[i])

// 		} else if arr[i] > pivot {
// 			rightSubArray = append(rightSubArray, arr[i])
// 		}
// 	}

// 	return append(append(quickSort(leftSubArray), pivot), quickSort(rightSubArray)...)
// }

// func quickSortCall(arr []int) {

// 	fmt.Printf("Starting quick sort algorithm...\n")
// 	fmt.Printf("Unsorted array: %v\n\n", arr)

// 	newArr := quickSort(arr)

// 	fmt.Printf("\nSorted array: %v\n", newArr)
// 	fmt.Printf("Selection sort algorithm successfully finished!\n")
// }

// func solution(numbers []int) []int {

// 	var zigZag []int
// 	arrLen := len(numbers)

// 	// if arrLen - 2 < 0
// 	//     return zigZag

// 	// if numbers[arrLen-3] < numbers[arrLen-2] > numbers[arrLen-1]

// 	for j := 0; j < arrLen-2; j++ {

// 		if (numbers[j] < numbers[j+1] && numbers[j+1] > numbers[j+2]) || (numbers[j] > numbers[j+1] && numbers[j+1] < numbers[j+2]) {

// 			zigZag = append(zigZag, 1)

// 		} else {
// 			zigZag = append(zigZag, 0)
// 		}
// 	}

// 	return zigZag
// }

// func main() {
// 	// bubbleSort(generate(30))
// 	// selectionSort(generate(30))
// 	// insertionSort(generate(20))
// 	// quickSortCall(generate(100))

// 	arr := []int{1, 2, 1, 3, 4}
// 	fmt.Print(solution(arr))
// }

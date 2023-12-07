package main

import "fmt"

const size = 10

// написать функцию для сортировки массива вставками
/*
func main() {
	data := [size]int{10, 5, 90, 3, 6, 33, 37, 81, 500, 1000}
	fmt.Println("BEFORE", data)
	sortedData := insertionSort(data)
	fmt.Println("AFTER", sortedData)
}

func insertionSort(input [size]int) [size]int {
	for i := 1; i < size; i++ {
		j := i
		for j > 0 {
			if input[j-1] > input[j] {
				input[j-1], input[j] = input[j], input[j-1]
			}
			j = j - 1
		}
	}
	return input
}
*/

// написать функцию, принимающую на вход неограниченное число переменных типа int
// сортирующая этот массив пузырьком в обратном порядке
func main() {
	data := [size]int{10, 5, 90, 3, 6, 33, 37, 81, 500, 1000}
	fmt.Println("BEFORE", data)
	newData := general(data)
	fmt.Println("AFTER", newData)
}

func bubbleSort(input [size]int) [size]int {
	for i := 0; i < size-1; i++ {
		if input[i] > input[i+1] {
			input[i], input[i+1] = input[i+1], input[i]
		}
	}
	return input
}

func reverse(input [size]int) [size]int {
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func general(input [size]int) [size]int {
	input = bubbleSort(input)
	input = reverse(input)
	return input
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nums [5]int
	var num string
	fmt.Println("Требуется ввести 5 чисел.")
	for i := 0; i <= 4; i++ {
		fmt.Printf("Введите %d число:\n", i+1)
		_, err := fmt.Scan(&num)
		if err != nil {
			panic(err)
		}
		num, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	fmt.Println(reverseArray(nums))
}

func reverseArray(arr [5]int) [5]int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

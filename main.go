package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	// Проверка прямо в условии
	if number%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
}

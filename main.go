package main

import (
	"fmt"
	"log"
)

func main() {
	var number int
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatalf("ERROR %s\n", err)
	}

	// Проверка прямо в условии
	if number%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
}

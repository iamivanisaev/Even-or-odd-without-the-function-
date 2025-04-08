package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Время начала:", start)

	var number int
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatalf("ERROR %s\n", err)
	}

	if number%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}

	end := time.Now()
	fmt.Println("Время конца:", end)
}

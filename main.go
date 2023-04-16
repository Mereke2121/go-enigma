package main

import (
	"fmt"
	"log"
)

func main() {
	rotor := NewRotor()
	fmt.Println("====================== Rotor is ready ======================")

	// сдвигаем ротор на начальное положение
	fmt.Println("Input start position for rotor:")
	var initPosition int
	_, err := fmt.Scan(&initPosition)
	if err != nil {
		log.Fatal(err)
	}
	rotor.Move(initPosition)

	// слушаем ввод
	for {
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal("input in wrong format")
		}

		hashedMessage := rotor.Hash(input)
		fmt.Println(hashedMessage)
	}
}

func printLetters(letters map[rune]rune) {
	for key, value := range letters {
		fmt.Printf("key: %s, value: %s\n", string(key), string(value))
	}
	fmt.Println("===================================================")
}

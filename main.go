package main

import (
	"fmt"
	"log"
)

func main() {
	firstRotor := NewRotor()
	secondRotor := NewRotor()
	thirdRotor := NewRotor()

	enigma := NewEnigma(firstRotor, secondRotor, thirdRotor)
	fmt.Println("====================== Enigma is ready ======================")

	// сдвигаем роторы на начальное положение
	moveRotorsOnBasePositions(enigma)

	// слушаем ввод
	for {
		fmt.Println("Input word:")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal("input in wrong format")
		}

		fmt.Println(enigma.Hash(input))
	}
}

func moveRotorsOnBasePositions(enigma *Enigma) {
	fmt.Println("Input start position for first rotor:")
	var initPosition int
	_, err := fmt.Scan(&initPosition)
	if err != nil {
		log.Fatal(err)
	}
	enigma.firstRotor.Move(initPosition)

	fmt.Println("Input start position for second rotor:")
	_, err = fmt.Scan(&initPosition)
	if err != nil {
		log.Fatal(err)
	}
	enigma.secondRotor.Move(initPosition)

	fmt.Println("Input start position for third rotor:")
	_, err = fmt.Scan(&initPosition)
	if err != nil {
		log.Fatal(err)
	}
	enigma.thirdRotor.Move(initPosition)
}

func printLetters(letters map[rune]rune) {
	for key, value := range letters {
		fmt.Printf("key: %s, value: %s\n", string(key), string(value))
	}
	fmt.Println("===================================================")
}

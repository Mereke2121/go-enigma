package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-enigma/enigma"
	"github.com/go-enigma/rotor"
)

func main() {
	firstRotor := rotor.NewRotor()
	secondRotor := rotor.NewRotor()
	thirdRotor := rotor.NewRotor()

	enigma := enigma.NewEnigma(firstRotor, secondRotor, thirdRotor)
	fmt.Println("====================== Enigma is ready ======================")

	// сдвигаем роторы на начальное положение
	moveRotorsOnBasePositions(enigma)

	// слушаем ввод
	for {
		fmt.Println("Input word:")
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal("input in wrong format")
		}

		fmt.Println(enigma.Hash(input))
	}
}

func moveRotorsOnBasePositions(enigma *enigma.Enigma) {
	fmt.Println("Do you want to encrypt or decrypt message; print 'e' or 'd':")
	var isEncrypt bool
	var encryptInput string
	_, err := fmt.Scan(&encryptInput)
	if err != nil {
		log.Fatal(err)
	}
	switch encryptInput {
	case "e":
		isEncrypt = true
	case "d":
		isEncrypt = false
	default:
		err = errors.New("invalid encrypt input")
		log.Fatal(err)
	}

	enigma.IsEncrypt = isEncrypt

	fmt.Println("Input start position for first rotor:")
	var initPosition int
	_, err = fmt.Scan(&initPosition)
	if err != nil {
		log.Fatal(err)
	}
	enigma.FirstRotor.Move(initPosition, isEncrypt)

	fmt.Println("Input start position for second rotor:")
	_, err = fmt.Scan(&initPosition)
	if err != nil {
		log.Fatal(err)
	}
	enigma.SecondRotor.Move(initPosition, isEncrypt)

	fmt.Println("Input start position for third rotor:")
	_, err = fmt.Scan(&initPosition)
	if err != nil {
		log.Fatal(err)
	}
	enigma.ThirdRotor.Move(initPosition, isEncrypt)
}

func printLetters(letters map[rune]rune) {
	for key, value := range letters {
		// if string(key) == "a" {
		fmt.Printf("key: %s, value: %s\n", string(key), string(value))
		// }
	}
	fmt.Println("===================================================")
}

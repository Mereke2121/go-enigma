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

	fmt.Println("Do you want to encrypt or decrypt message; print 'e' or 'd':")
	isEncrypt := isEncrypt()

	enigma := enigma.NewEnigma(firstRotor, secondRotor, thirdRotor, isEncrypt)
	fmt.Println("====================== Enigma is ready ======================")

	// сдвигаем роторы на начальное положение
	enigma.MoveRotorsOnBasePositions()

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

func isEncrypt() bool {
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
	return isEncrypt
}

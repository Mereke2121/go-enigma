package main

import "fmt"

func main() {
	rotor := NewRotor()
	printLetters(rotor.letters)
	rotor.Move(2)
	fmt.Println("==============================================================")
	printLetters(rotor.letters)
}

func printLetters(letters map[rune]rune) {
	for key, value := range letters {
		fmt.Printf("key: %s, value: %s\n", string(key), string(value))
	}
}

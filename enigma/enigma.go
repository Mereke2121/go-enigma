package enigma

import (
	"fmt"
	"log"

	"github.com/go-enigma/rotor"
)

type Enigma struct {
	FirstRotor  *rotor.Rotor
	SecondRotor *rotor.Rotor
	ThirdRotor  *rotor.Rotor
	IsEncrypt   bool
}

func NewEnigma(firstRotor, secondRotor, thirdRotor *rotor.Rotor, isEncrypt bool) *Enigma {
	return &Enigma{
		FirstRotor:  firstRotor,
		SecondRotor: secondRotor,
		ThirdRotor:  thirdRotor,
		IsEncrypt:   isEncrypt,
	}
}

func (e *Enigma) Hash(input string) string {
	var hashedMessage string

	// основная логика алгоритма enigma
	firstCounter, secondCounter := 0, 0
	for _, letter := range input {
		letterOfFirstRotor := e.FirstRotor.Letters[letter]
		e.FirstRotor.Move(1, e.IsEncrypt)
		firstCounter++

		if firstCounter%26 == 0 && firstCounter > 26 {
			e.SecondRotor.Move(1, e.IsEncrypt)
			secondCounter++
		}
		letterOfSecondRotor := e.SecondRotor.Letters[letterOfFirstRotor]

		if secondCounter%26 == 0 && secondCounter > 26 {
			e.ThirdRotor.Move(1, e.IsEncrypt)
		}
		letterOfThirdRotor := e.ThirdRotor.Letters[letterOfSecondRotor]

		hashedMessage += string(letterOfThirdRotor)
	}

	return hashedMessage
}

func (e *Enigma) MoveRotorsOnBasePositions() {
	fmt.Println("Input start position for first rotor:")
	var initPosition int
	_, err := fmt.Scan(&initPosition)
	if err != nil {
		log.Fatal(err)
	}
	e.FirstRotor.Move(initPosition, e.IsEncrypt)

	fmt.Println("Input start position for second rotor:")
	_, err = fmt.Scan(&initPosition)
	if err != nil {
		log.Fatal(err)
	}
	e.SecondRotor.Move(initPosition, e.IsEncrypt)

	fmt.Println("Input start position for third rotor:")
	_, err = fmt.Scan(&initPosition)
	if err != nil {
		log.Fatal(err)
	}
	e.ThirdRotor.Move(initPosition, e.IsEncrypt)
}

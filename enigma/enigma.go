package enigma

import "github.com/go-enigma/rotor"

type Enigma struct {
	FirstRotor  *rotor.Rotor
	SecondRotor *rotor.Rotor
	ThirdRotor  *rotor.Rotor
	IsEncrypt   bool
}

func NewEnigma(firstRotor, secondRotor, thirdRotor *rotor.Rotor) *Enigma {
	return &Enigma{
		FirstRotor:  firstRotor,
		SecondRotor: secondRotor,
		ThirdRotor:  thirdRotor,
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

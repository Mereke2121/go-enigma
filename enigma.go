package main

type Enigma struct {
	firstRotor  *Rotor
	secondRotor *Rotor
	thirdRotor  *Rotor
}

func NewEnigma(firstRotor, secondRotor, thirdRotor *Rotor) *Enigma {
	return &Enigma{
		firstRotor:  firstRotor,
		secondRotor: secondRotor,
		thirdRotor:  thirdRotor,
	}
}

func (e *Enigma) Hash(input string) string {
	var hashedMessage string

	// основная логика алгоритма enigma
	firstCounter, secondCounter := 0, 0
	for _, letter := range input {
		letterOfFirstRotor := e.firstRotor.letters[letter]
		e.firstRotor.Move(1)
		firstCounter++

		if firstCounter%26 == 0 && firstCounter > 26 {
			e.secondRotor.Move(1)
			secondCounter++
		}
		letterOfSecondRotor := e.secondRotor.letters[letterOfFirstRotor]

		if secondCounter%26 == 0 && secondCounter > 26 {
			e.thirdRotor.Move(1)
		}
		letterOfThirdRotor := e.thirdRotor.letters[letterOfSecondRotor]

		hashedMessage += string(letterOfThirdRotor)
	}

	return hashedMessage
}

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

	for _, letter := range input {
		letterOfFirstRotor := e.firstRotor.letters[letter]
		letterOfSecondRotor := e.secondRotor.letters[letterOfFirstRotor]
		letterOfThirdRotor := e.thirdRotor.letters[letterOfSecondRotor]

		hashedMessage += string(letterOfThirdRotor)
	}

	return hashedMessage
}

package main

type Rotor struct {
	letters map[rune]rune
}

func NewRotor() *Rotor {
	// full letters
	letters := make(map[rune]rune)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, char := range alphabet {
		letters[char] = char
	}

	return &Rotor{letters: letters}
}

func (r *Rotor) Move(shift int) {
	shiftedLetters := make(map[rune]rune)
	for key, char := range r.letters {
		// Вычисляем новое значение символа, сдвинутое на shift позиций
		shiftedChar := (char-'a'+rune(shift))%26 + 'a'

		// Добавляем новое значение в новую мапу
		shiftedLetters[key] = shiftedChar
	}
	r.letters = shiftedLetters
}

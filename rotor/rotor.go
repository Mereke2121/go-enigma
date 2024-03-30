package rotor

type Rotor struct {
	Letters map[rune]rune
}

func NewRotor() *Rotor {
	// full letters
	letters := make(map[rune]rune)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, char := range alphabet {
		letters[char] = char
	}

	return &Rotor{Letters: letters}
}

func (r *Rotor) Move(shift int, isEncrypt bool) {
	shiftedLetters := make(map[rune]rune)
	for key, char := range r.Letters {
		var shiftedChar rune
		// сдвигаем значение на n(при дешифровке сдвигаем в обратную сторону)
		if isEncrypt {
			// вычисляем новое значение символа, сдвинутое на shift позиций
			shiftedChar = (char-'a'+rune(shift))%26 + 'a'
		} else {
			shiftedChar = (char-'a'+26-rune(shift))%26 + 'a'
		}
		// добавляем новое значение в новую мапу
		shiftedLetters[key] = shiftedChar
	}
	r.Letters = shiftedLetters
}

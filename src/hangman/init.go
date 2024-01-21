package hangman

func (h *HangmanData) HangmanInit() {
	h.Win = false
	h.Lost = false
	h.Try = 7
	h.RandomWord()
	h.SecretWord()
	
}
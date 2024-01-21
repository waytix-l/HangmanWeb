package hangman

import "fmt"

func (h *HangmanData) Hangman() {
	h.LetterVerification()
	fmt.Println(h.Try)
	if h.WordVerification() {
		h.Win = true
		fmt.Println("vous avez gagné")
	}
	h.LostCondition()
}

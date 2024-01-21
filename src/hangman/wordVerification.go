package hangman

import (
	"fmt"
)

func (h *HangmanData) LetterVerification() {
	tryVerification := false
	for i := 0; i < len(h.Mot_secret); i++ {
		if h.Word_Table[i] == h.Gessed_letters {
			h.ListMotSecret[i] = h.Gessed_letters
			tryVerification = true
		}
	}

	if !tryVerification {
		h.Try -= 1
	}
	
	fmt.Println(h.ListMotSecret)

	h.Mot_secret = ""
	for j := 0; j < len(h.ListMotSecret); j++ {
		h.Mot_secret += h.ListMotSecret[j]
	}
	fmt.Println(h.Mot_secret)
}

func (h *HangmanData) WordVerification() bool {
	for i := 0; i < len(h.ListMotSecret); i++ {
		if (h.Word_Table[i] != h.ListMotSecret[i]) {
			return false
		}
	}
	return true
}

func (h *HangmanData) LostCondition() {
	if (h.Try <= 0) {
		h.Lost = true
		fmt.Println(h.Lost)
	}
}

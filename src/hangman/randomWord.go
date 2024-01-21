package hangman

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func (h *HangmanData) RandomWord() {
	h.Content, h.Err = os.ReadFile("hangman/mot.txt")
	if h.Err != nil {
		fmt.Print(h.Err)
	}
	lines := strings.Split(string(h.Content), "\n")

	rand.Seed(time.Now().UnixNano())
	h.Random_number = rand.Intn(len(lines))

	h.Word = lines[h.Random_number]
	h.Word_Table = strings.Split(h.Word, "")

	// for i := 0; i < len(h.Word); i++ {
	// 	h.Mot_secret = append(h.Mot_secret, "_")
	// }
}

func (h *HangmanData) SecretWord() {
	h.ListMotSecret = nil
	h.Mot_secret = ""
	for i := 0; i < len(h.Word); i++ {
		h.ListMotSecret = append(h.ListMotSecret, "_")
	}
	for i := 0; i < len(h.ListMotSecret); i++ {
		h.Mot_secret += "_"
	}
}

package hangman

import (
	"os"
	"fmt"
	"strings"
	"math/rand"
	"time"

)

type HangmanData struct {
	Word string
	Content []byte
	Err error
	Random_number int
	Tableau_rune []rune
	Mot_secret []string
}

func (h *HangmanData) RandomWord() {
	h.Content, h.Err = os.ReadFile("hangman/mot.txt")
	if h.Err != nil {
		fmt.Print(h.Err)
	}
	lines := strings.Split(string(h.Content), "\n")

	rand.Seed(time.Now().UnixNano())
	h.Random_number = rand.Intn(len(lines))

	h.Word = lines[h.Random_number]
	h.Tableau_rune = []rune(h.Word)

	for i := 0; i < len(h.Word); i++ {
		h.Mot_secret = append(h.Mot_secret, "_")
	}
}









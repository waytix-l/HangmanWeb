package hangman

type HangmanData struct {
	Content        []byte
	Err            error

	Random_number  int

	Word           string
	Word_Table []string

	Mot_secret     string
	ListMotSecret  []string
	
	Gessed_letters string

	Win bool
	Lost bool
	Try int
}

package shortener

import (
	"math/rand"
	"time"
)

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
var capitalCharacterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")

func CutUrl() string {
	return randomLink()
}

func randomLink() string {
	res := make([]rune, 10)
	for i := 0; i < 3; i++ {
		res[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	for i := 3; i < 6; i++ {
		res[i] = capitalCharacterRunes[rand.Intn(len(capitalCharacterRunes))]
	}
	for i := 6; i < 9; i++ {
		res[i] = numberRunes[rand.Intn(len(numberRunes))]
	}
	res[9] = rune('_')

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(res), func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})

	return string(res)
}

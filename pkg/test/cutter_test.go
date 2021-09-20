package test

import (
	"strings"
	"testing"

	"github.com/yurioganesyan/LinkShortener/pkg/shortener"
)

func TestCutter_randomline(t *testing.T) {
	underscore := '_'
	var characterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	var capitalCharacterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numberRunes = []rune("0123456789")

	result := shortener.CutUrl()

	if !strings.Contains(result, string(underscore)) {
		t.Error("incorrect result. There is no underscore")
	}
	if !strings.ContainsAny(result, string(characterRunes)) {
		t.Error("incorrect result. There is no lowercase letters")
	}
	if !strings.ContainsAny(result, string(capitalCharacterRunes)) {
		t.Error("incorrect result. There is no capital letters")
	}
	if !strings.ContainsAny(result, string(numberRunes)) {
		t.Error("incorrect result. There is no numbers")
	}
}

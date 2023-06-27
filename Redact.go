package redact

import (
	"fmt"
	"strings"
)

// Word returns the first character of the word, followed by "*"s for each additional character
func Word(word string) string {
	return WordOptions(word, 1, 0, 0)
}

// WordOptions lets you choose how many characters to reveal at the beginning and max a
func WordOptions(word string, prefixChars, suffixChars uint, maxAsterisk uint) string {
	runes := []rune(word)
	rl := uint(len(runes))

	if rl <= 1 {
		return word
	}
	if rl <= prefixChars+suffixChars {
		if prefixChars >= rl {
			prefixChars = rl - 1
			if suffixChars > 0 && prefixChars+1 > 2 {
				prefixChars--
			}
		}
		if suffixChars > 0 {
			suffixChars = rl - prefixChars - 1
		}
	}
	l := rl - prefixChars - suffixChars
	if maxAsterisk > 0 {
		if l > uint(maxAsterisk) {
			l = uint(maxAsterisk)
		}
	}

	prefix := string(runes[:prefixChars])
	suffix := string(runes[rl-suffixChars:])
	return fmt.Sprintf("%s%s%s", prefix, strings.Repeat("*", int(l)), suffix)
}

// Words redacts each word in a string.
func Words(words string) string {
	ws := strings.Split(words, " ")
	for i := range ws {
		ws[i] = Word(ws[i])
	}
	return strings.Join(ws, " ")
}

// Email redacts the part of an email before the @.
func Email(email string) string {
	parts := strings.Split(email, "@")
	parts[0] = Words(parts[0])
	return strings.Join(parts, "@")
}

// Phone redacts the last digits of a phone number.
func Phone(phone string) string {
	parts := strings.Split(phone, "-")
	lastRunes := []rune(parts[len(parts)-1])
	lastPartlen := len(lastRunes)
	repeatChars := 4
	if repeatChars > lastPartlen {
		repeatChars = lastPartlen
	}

	parts[len(parts)-1] = string(append(lastRunes[0:lastPartlen-repeatChars], []rune(strings.Repeat("*", repeatChars))...))
	return strings.Join(parts, "-")
}

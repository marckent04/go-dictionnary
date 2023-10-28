package domain

import "time"

type WordEntry struct {
	Word, Definition string
	CreatedAt        time.Time
}

func NewWordEntry(word, definition string) WordEntry {
	return WordEntry{
		Word:       word,
		Definition: definition,
		CreatedAt:  time.Now(),
	}
}

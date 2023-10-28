package domain

import (
	"fmt"
	"log"
	"marckent/dictionary/shared"
)

type Dictionary struct {
	repository WordRepository
}

func (d *Dictionary) AddWord(word, definition string) {
	err := d.repository.Save(NewWordEntry(word, definition))
	shared.HandleError(err, fmt.Sprintf("\"%s\" has been saved", word))
}

func (d *Dictionary) RemoveWord(word string) {
	err := d.repository.Remove(word)
	shared.HandleError(err, fmt.Sprintf("\"%s\" has been removed", word))
}

func (d *Dictionary) GetWord(word string) WordEntry {
	wordEntry, err := d.repository.GetOne(word)
	if err != nil {
		log.Fatal(err)
	}

	return wordEntry

}

func (d *Dictionary) ListAllWords() (words []WordEntry, err error) {
	words, err = d.repository.GetAll()
	return
}

func NewDictionary(repository WordRepository) Dictionary {
	return Dictionary{repository: repository}
}

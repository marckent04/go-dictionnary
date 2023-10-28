package cli

import (
	"errors"
	"fmt"
	"marckent/dictionary/domain"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type App struct {
	label, definition, action string
	dictionary                domain.Dictionary
}

func (l *App) Init(dictionary domain.Dictionary) {
	l.dictionary = dictionary
}

func (l *App) Start() (err error) {
	err = l.getActionAndArgs()
	if err != nil {
		return
	}

	err = l.actionValidator()
	if err != nil {
		return
	}

	err = l.execute()

	return
}

func (l *App) getActionAndArgs() (err error) {

	args := os.Args
	if args[1] != "-action" {
		err = errors.New("unknown command")
	}

	for key, v := range args[2:] {
		switch key {
		case 0:
			l.action = v
		case 1:
			l.label = v
		case 2:
			l.definition = v
		}
	}
	return
}

func (l *App) actionValidator() (err error) {
	isLabelEmpty := l.label == ""
	isDefEmpty := l.definition == ""

	if l.action == "list" {
		return
	}

	if isLabelEmpty {
		err = errors.New("word is missed")
		return
	}

	if l.action == "add" && isDefEmpty {
		err = errors.New("definition is missed")
	}

	return
}

func (l *App) execute() (err error) {
	switch l.action {
	case "list":
		words, getError := l.dictionary.ListAllWords()

		if getError != nil {
			err = getError
			return
		}

		l.displayWords(words)

	case "add":
		l.dictionary.AddWord(l.label, l.definition)
	case "show":
		entry := l.dictionary.GetWord(l.label)
		l.displayWords([]domain.WordEntry{entry})
	case "remove":
		l.dictionary.RemoveWord(l.label)
	default:
		err = errors.New("unknown action")
	}

	return

}

func (l *App) displayWords(words []domain.WordEntry) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Word", "Definition", "Created At"})

	for _, word := range words {
		tw.AppendRow(table.Row{word.Word, word.Definition, word.CreatedAt.Format(time.Stamp)})
	}

	fmt.Println(tw.Render())

}

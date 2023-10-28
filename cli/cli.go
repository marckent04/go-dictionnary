package cli

import (
	"errors"
	"flag"
	"fmt"
	"marckent/dictionary/domain"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type App struct {
	word, definition, action string
	dictionary               domain.Dictionary
}

func (l *App) Init(dictionary domain.Dictionary) {
	l.dictionary = dictionary
}

func (l *App) Start() (err error) {
	l.getActionAndArgs()

	err = l.actionValidator()
	if err != nil {
		return
	}

	err = l.execute()

	return
}

func (l *App) getActionAndArgs() {

	action := flag.String("action", "list", "dictionary action")
	flag.Parse()
	args := flag.Args()

	l.action = *action

	if len(args) >= 1 {
		l.word = args[0]
	}

	if len(args) >= 2 {
		l.definition = args[1]
	}
}

func (l *App) actionValidator() (err error) {
	isWordEmpty := l.word == ""
	isDefEmpty := l.definition == ""

	if l.action == "list" {
		return
	}

	if isWordEmpty {
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
		l.dictionary.AddWord(l.word, l.definition)
	case "show":
		entry := l.dictionary.GetWord(l.word)
		l.displayWords([]domain.WordEntry{entry})
	case "remove":
		l.dictionary.RemoveWord(l.word)
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

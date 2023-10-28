package main

import (
	"fmt"
	"log"
	"marckent/dictionary/cli"
	"marckent/dictionary/database"
	"marckent/dictionary/domain"

	"github.com/dgraph-io/badger/v4"
)

func main() {
	db, err := badger.Open(badger.DefaultOptions("./db/words").WithLoggingLevel(badger.ERROR))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := setup(db)

	appError := app.Start()

	if appError != nil {
		fmt.Println(appError)
	}
}

func setup(db *badger.DB) cli.App {
	var wordRepository domain.WordRepository
	wordDb := database.NewWordRepository(db)
	wordRepository = wordDb

	dictionary := domain.NewDictionary(wordRepository)

	program := cli.App{}
	program.Init(dictionary)

	return program
}

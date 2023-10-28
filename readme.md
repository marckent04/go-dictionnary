# GO Dictionary
### an persistent cli dictionary built in Golang


## Features
- [X] Add word
- [X] List all dictionary words
- [X] Remove a word from the dictionary
- [X] Get specific word definition
- [ ] Unit tests
- [ ] e2e tests


## Requirements
- Go 2.21.0 minimum
- command line interface

## How to install
Run `go get` in order to install project dependencies


## How to use
Run `go run main -action` with these arguments:
-   `add {word} {definiton}` for add a new word
-   `list` for list all dictionary words
-   `remove {word}` for remove a word
-   `show {word}` for see a specific word definition

> Example: go run main -action add Go "Cool Language"
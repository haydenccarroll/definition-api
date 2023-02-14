# definition_api

A simple http API in Go to lookup the definition of a word, and simultaneously add it to an Anki deck.

This is used for myself when I am reading. I can simply lookup the word as a URL path arg, and add it to a deck at the same time for studying.

## Contributing

There is a pre-commit githook for this repository. Please use it. I don't want to setup GH Actions, so make sure you have your githooks setup by running `make githooks`. You may need to install golangci-lint with `brew install golangci-lint`.
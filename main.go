package main

import (
	"github.com/joho/godotenv"
	"github.com/launchboxio/launchbox/command"
)

func main() {
	godotenv.Load()

	command.Execute()
}

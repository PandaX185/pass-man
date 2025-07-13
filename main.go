package main

import (
	"github.com/PandaX185/pass-man/cmd"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("/usr/local/bin/pass-man.env"); err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	cmd.Execute()
}

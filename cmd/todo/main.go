package main

import (
	"log"
	"os"

	"graphql-server/di"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed to load config from environment")
	}
	cmd := di.ResolveCommand()()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

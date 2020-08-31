package main

import (
	"log"

	"github.com/RanchoCooper/go-by-demos/wire/example/db"
	"github.com/RanchoCooper/go-by-demos/wire/example/di"
	"github.com/RanchoCooper/go-by-demos/wire/example/service"
)

func main() {
	// using wire to auto inject dependency
	dbConfig := &db.Config{}
	mailConfig := &service.MailConfig{}

	s, err := di.NewService(dbConfig, mailConfig)
	if err != nil {
		log.Fatal(err)
	}
	s.UserSignUp()
}

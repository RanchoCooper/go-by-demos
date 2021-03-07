package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/RanchoCooper/go-by-demos/articles/singleton-database/pkg/config"
	"github.com/RanchoCooper/go-by-demos/articles/singleton-database/pkg/db"
	"github.com/RanchoCooper/go-by-demos/articles/singleton-database/pkg/reader"
	"github.com/RanchoCooper/go-by-demos/articles/singleton-database/pkg/writer"
)

func init() {
	err := config.Init()
	if err != nil {
		panic(err)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var quit = make(chan struct{})

	// do some init from db
	_ = db.DB()

	go func() {
		writer.Run(quit)
		wg.Done()
	}()
	go func() {
		reader.Run(quit)
		wg.Done()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	_ = <-c
	close(quit)
	log.Printf("recv exit signal...")
	wg.Wait()
	log.Printf("program exit ok")
}
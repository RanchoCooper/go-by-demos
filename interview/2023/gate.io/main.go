package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type APIKey struct {
	ID        string
	UpdatedAt time.Time
}

func getAPIKey(ctx context.Context) (*APIKey, error) {
	return &APIKey{ID: uuid.New().String(), UpdatedAt: time.Now()}, nil
}

func saveToDB(ctx context.Context, records []*APIKey) error {
	fmt.Println("save to db")
	return nil
}

// 1. 100 records
// 2. every 30s
func main() {
	ctx := context.TODO()
	tmp := make([]*APIKey, 0)

	go func() {
		for {
			apiKey, err := getAPIKey(ctx)
			if err != nil {
				// log
			}
			//fmt.Println("len tmp: ", len(tmp))
			tmp = append(tmp, apiKey)
			time.Sleep(1 * time.Millisecond)
		}
	}()

	success := make(chan bool)

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Println("5s ticker")
				success <- true
			}
		}
	}()

	go func() {
		for {
			if len(tmp) >= 100 {
				fmt.Println("100 record")
				success <- true
			}
		}
	}()

	for {
		select {
		case <-success:
			err := saveToDB(ctx, tmp)
			if err != nil {
				// log
			}
			tmp = make([]*APIKey, 0)
		}
	}
}

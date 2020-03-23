package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	uri := os.Getenv("DB_URI")
	client, err := connect(uri)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("glacierapi")

	themes, err := getAllThemes(db)
	if err != nil {
		log.Fatal(err)
	}

	votes, err := getAllVotes(db)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	start := time.Now().UnixNano()

	var allTransition []transition
	now := int(time.Now().UnixNano() / int64(time.Millisecond))

	for _, theme := range themes {
		transition := calcTransition(votes, theme, now)
		allTransition = append(allTransition, transition)
	}

	// fmt.Println(allTransition)
	end := time.Now().UnixNano()
	fmt.Println(start, end, (end-start)/1000000)
}

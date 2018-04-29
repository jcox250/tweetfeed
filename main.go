package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jcox250/tweetfeed/usecases"
)

var accessToken = os.Getenv("ACCESS_TOKEN")
var accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")

func checkEnvVars() {
	if accessToken == "" {
		log.Fatal("ACCESS_TOKEN envvar is not set")
	}
	log.Printf("ACCESS_TOKEN: %s", accessToken)

	if accessTokenSecret == "" {
		log.Fatal("ACCESS_TOKEN_SECRET envvar is not set")
	}
	log.Printf("ACCESS_TOKEN_SECRET: %s", accessTokenSecret)
}

func main() {
	checkEnvVars()
	i := usecases.NewTwitterInteractor(accessToken, accessTokenSecret)
	results, err := i.GetSearch("golang", nil)
	if err != nil {
		log.Fatalf("couldn't get tweets: %s", err)
	}

	fmt.Println(results)
}

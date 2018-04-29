package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jcox250/tweetfeed/adapters/service"
	"github.com/jcox250/tweetfeed/usecases"
)

var accessToken = os.Getenv("ACCESS_TOKEN")
var accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
var consumerKey = os.Getenv("CONSUMER_KEY")
var consumerKeySecret = os.Getenv("CONSUMER_KEY_SECRET")

func checkEnvVars() {
	if accessToken == "" {
		log.Fatal("ACCESS_TOKEN envvar is not set")
	}

	if accessTokenSecret == "" {
		log.Fatal("ACCESS_TOKEN_SECRET envvar is not set")
	}

	if consumerKey == "" {
		log.Fatal("CONSUMER_KEY envvar is not set")
	}
	if consumerKeySecret == "" {
		log.Fatal("CONSUMER_KEY_SECRET envvar is not set")
	}
}

func main() {
	checkEnvVars()
	twitterInteractor := usecases.NewTwitterInteractor(accessToken, accessTokenSecret, consumerKey, consumerKeySecret)
	twitterService := service.NewTwitterService(twitterInteractor)

	results := twitterService.Search("golang", nil)

	fmt.Println(results)
}

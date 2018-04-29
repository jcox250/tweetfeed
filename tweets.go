package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/jcox250/tweets/adapters/service"
	"github.com/jcox250/tweets/usecases"
)

var accessToken = os.Getenv("ACCESS_TOKEN")
var accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
var consumerKey = os.Getenv("CONSUMER_KEY")
var consumerKeySecret = os.Getenv("CONSUMER_KEY_SECRET")

// Flag vars
var user string
var count int

// Default twitter values
var values = url.Values{
	"count": []string{strconv.Itoa(count)},
}

func main() {
	checkEnvVars()
	flag.Parse()

	// if user does not supply flags, print usage
	// we can clean this up later by putting this into its own function
	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	twitterInteractor := usecases.NewTwitterInteractor(accessToken, accessTokenSecret, consumerKey, consumerKeySecret)
	twitterService := service.NewTwitterService(twitterInteractor)

	tweets := twitterService.GetUserTimeline(user, values)

	for _, v := range tweets {
		fmt.Println("Tweet From:", v.User)
		fmt.Println(v.Text)
		fmt.Println("On:", v.CreatedAt)
		fmt.Println("")
	}

}

func init() {
	flag.StringVar(&user, "u", "", "Search for users tweets")
	flag.IntVar(&count, "c", 10, "Number of tweets to return")
}

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

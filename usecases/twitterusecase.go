package usecases

import (
	"log"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/jcox250/tweetfeed/domain"
)

type TwitterInteractor struct {
	api *anaconda.TwitterApi
}

func NewTwitterInteractor(accessToken, accessTokenSecret, consumerKey, consumerKeySecret string) TwitterInteractor {
	return TwitterInteractor{
		api: anaconda.NewTwitterApiWithCredentials(
			accessToken,
			accessTokenSecret,
			consumerKey,
			consumerKeySecret,
		),
	}
}

// GetSearch
func (t TwitterInteractor) GetSearch(searchTerm string, values url.Values) (anaconda.SearchResponse, error) {
	results, err := t.api.GetSearch(searchTerm, values)
	return results, err
}

func (t TwitterInteractor) GetUserTimeline(user string, values url.Values) []domain.Tweet {
	values.Set("screen_name", user)
	tweets, err := t.api.GetUserTimeline(values)
	if err != nil {
		log.Fatalf("error getting tweets for user: %s: %s", user, err)
	}
	return anacondaToTweet(tweets)
}

func anacondaToTweet(tweets []anaconda.Tweet) []domain.Tweet {
	var ts []domain.Tweet
	for _, v := range tweets {
		t := domain.Tweet{
			User: v.User.Name,
			Text: v.Text,
		}
		ts = append(ts, t)
	}
	return ts
}

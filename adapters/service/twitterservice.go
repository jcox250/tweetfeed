package service

import (
	"log"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/jcox250/tweets/domain"
)

type twitterInteractor interface {
	GetSearch(string, url.Values) (anaconda.SearchResponse, error)
	GetUserTimeline(string, url.Values) []domain.Tweet
}

type TwitterService struct {
	// this needs to be an interface
	interactor twitterInteractor
}

func NewTwitterService(interactor twitterInteractor) *TwitterService {
	return &TwitterService{
		interactor: interactor,
	}
}

func (t *TwitterService) Search(searchTerm string, values url.Values) anaconda.SearchResponse {
	results, err := t.interactor.GetSearch(searchTerm, values)
	if err != nil {
		log.Fatal("error searching for %s with values %v: %s",
			searchTerm,
			values,
			err,
		)
	}
	return results
}

func (t *TwitterService) GetUserTimeline(user string, values url.Values) []domain.Tweet {
	tweets := t.interactor.GetUserTimeline(user, values)
	return tweets
}

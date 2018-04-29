package usecases

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

type TwitterInteractor struct {
	api *anaconda.TwitterApi
}

func NewTwitterInteractor(accessToken, accessTokenSecret string) TwitterInteractor {
	return TwitterInteractor{
		api: anaconda.NewTwitterApi(accessToken, accessTokenSecret),
	}
}

// GetSearch
func (t TwitterInteractor) GetSearch(searchTerm string, values url.Values) (anaconda.SearchResponse, error) {
	results, err := t.api.GetSearch(searchTerm, values)
	return results, err
}

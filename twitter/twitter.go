package twitter

import (
	"../anaconda"
	"fmt"
)

const (
	ConsumerKey       = "aVPbYklTENfmoXI6V7GWhUqUm"
	ConsumerSecret    = "Nc2ibYyC4fjlrXuRrTQCRtCjcSm5HUz3TnW3O6bTOiLMMSWIb8"
	AccessToken       = "420103191-ESIZxnlnxeumcfJYNqbXFzgGObwwlulBtoEGgNf5"
	AccessTokenSecret = "RHfbeaPNEGdnUUSThYDjWxLlGxsyWq8CSvftN4EgfCQeS"
)

func Search(search string) string {

	anaconda.SetConsumerKey(ConsumerKey)
	anaconda.SetConsumerSecret(ConsumerSecret)
	api := anaconda.NewTwitterApi(AccessToken, AccessTokenSecret)

	search_result, err := api.GetSearch(search, nil)

	if err != nil {
		panic(err)
		return "false"
	}

	var result string

	for _, tweet := range search_result {
		result += tweet.Text
	}

	return result
}

func Tweet(status string) bool {

	anaconda.SetConsumerKey(ConsumerKey)
	anaconda.SetConsumerSecret(ConsumerSecret)
	api := anaconda.NewTwitterApi(AccessToken, AccessTokenSecret)

	result, err := api.PostTweet(status, nil)

	if err != nil {
		panic(err)
		return false
	}

	fmt.Println(result)

	return true
}

package twitter

import (
	"../anaconda"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JsonObj struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func parseJson(path string) *JsonObj {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	obj := &JsonObj{}
	if err := json.Unmarshal([]byte(data), obj); err != nil {
		fmt.Println(err)
		return nil
	}

	return obj
}

func Search(search string, path string) (result string, err error) {

	config := parseJson(path)

	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.AccessToken, config.AccessTokenSecret)

	search_result, err := api.GetSearch(search, nil)

	if err != nil {
		panic(err)
		return result, err
	}

	for _, tweet := range search_result {
		result += tweet.Text
	}

	return result, nil
}

func Tweet(status string, path string) (err error) {

	config := parseJson(path)

	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.AccessToken, config.AccessTokenSecret)

	result, err := api.PostTweet(status, nil)

	if err != nil {
		panic(err)
		return err
	}

	fmt.Println(result)

	return nil
}

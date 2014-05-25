package main

import (
	"./twitter"
	"fmt"
)

func main() {

	/* search tweet*/
	search := twitter.Search("fisproject")
	fmt.Println(search)

	/* post tweet*/
	tweet := twitter.Tweet("post-test")
	fmt.Println(tweet)

}

package main

import (
	"./twitter"
	"fmt"
)

func main() {

	/* search tweet*/
	search := twitter.Search("fisproject", "./roomba2d2.json")
	fmt.Println(search)

	/* post tweet*/
	tweet := twitter.Tweet("post-test", "./roomba2d2.json")
	fmt.Println(tweet)

}

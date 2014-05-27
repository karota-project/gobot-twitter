package main

import (
	"./twitter"
	"fmt"
)

func main() {

	/* search tweet*/
	search, err := twitter.Search("fisproject", "./roomba2d2.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(search)

	/* post tweet*/
	err = twitter.Tweet("post-test", "./roomba2d2.json")
	if err != nil {
		fmt.Println(err)
	}

}

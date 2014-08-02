package main

import (
	"fmt"
	"net/url"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/karota-project/gobot-twitter"
)

func main() {
	master := gobot.NewGobot()
	api.NewAPI(master).Start()

	twitterDriver := twitter.NewTwitterDriver("twitter-d01", "your-consumer-key", "your-consumer-secret")
	twitterDriver.SetAccessToken("your-access-token", "your-access-token-secret")

	master.AddRobot(
		gobot.NewRobot(
			"twitter",
			[]gobot.Connection{},
			[]gobot.Device{twitterDriver},
			func() {
				fmt.Println("work")

				value := url.Values{}
				tweets, err := twitterDriver.TwitterApi().GetSearch("roomba", value)

				if err == nil {
					for _, tweet := range tweets {
						fmt.Println(tweet.User.Name + "    @" + tweet.User.ScreenName)
						fmt.Println(tweet.Text)
						fmt.Println()
					}
				}
			}))

	master.Start()
}

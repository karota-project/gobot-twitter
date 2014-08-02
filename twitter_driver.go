package twitter

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/hybridgroup/gobot"
)

type TwitterDriver struct {
	gobot.Driver
	twitterApi *anaconda.TwitterApi
}

type TwitterInterface interface {
}

func NewTwitterDriver(name string, consumerKey string, consumerSecret string) *TwitterDriver {
	t := &TwitterDriver{
		Driver: *gobot.NewDriver(
			name,
			"twitter.TwitterDriver",
		),
	}

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	t.AddCommand("DeleteTweet", func(params map[string]interface{}) interface{} {
		id := params["id"].(float64)
		trimUser := params["trimUser"].(bool)

		//t.twitterApi.DeleteTweet(id int64, trimUser bool) (tweet Tweet, err error)
		return resultApi(t.twitterApi.DeleteTweet(int64(id), trimUser))
	})

	t.AddCommand("DisableThrottling", func(params map[string]interface{}) interface{} {
		t.twitterApi.DisableThrottling()
		return nil
	})

	t.AddCommand("EnableThrottling", func(params map[string]interface{}) interface{} {
		rate := params["rate"].(float64) // ms
		bufferSize := params["bufferSize"].(float64)

		//t.twitterApi.EnableThrottling(rate time.Duration, bufferSize int64)
		t.twitterApi.EnableThrottling(time.Duration(rate)*time.Millisecond, int64(bufferSize))
		return nil
	})

	t.AddCommand("Favorite", func(params map[string]interface{}) interface{} {
		id := params["id"].(float64)

		//t.twitterApi.Favorite(id int64) (rt Tweet, err error)
		return resultApi(t.twitterApi.Favorite(int64(id)))
	})

	t.AddCommand("GetDelay", func(params map[string]interface{}) interface{} {
		//t.twitterApi.GetDelay() time.Duration
		delay := t.twitterApi.GetDelay()

		return struct {
			Result int64 `json:"result"`
		}{int64(delay / time.Millisecond)}
	})

	t.AddCommand("GetDirectMessages", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetDirectMessages(v url.Values) (messages []DirectMessage, err error)
		return resultApi(t.twitterApi.GetDirectMessages(v))
	})

	t.AddCommand("GetDirectMessagesSent", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetDirectMessagesSent(v url.Values) (messages []DirectMessage, err error)
		return resultApi(t.twitterApi.GetDirectMessagesSent(v))
	})

	t.AddCommand("GetDirectMessagesShow", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetDirectMessagesShow(v url.Values) (messages []DirectMessage, err error)
		return resultApi(t.twitterApi.GetDirectMessagesShow(v))
	})

	t.AddCommand("GetFavorites", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetFavorites(v url.Values) (favorites []Tweet, err error)
		return resultApi(t.twitterApi.GetFavorites(v))
	})

	t.AddCommand("GetFollowersIds", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetFollowersIds(v url.Values) (c Cursor, err error)
		return resultApi(t.twitterApi.GetFollowersIds(v))
	})

	t.AddCommand("GetFollowersList", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetFollowersList(v url.Values) (c UserCursor, err error)
		return resultApi(t.twitterApi.GetFollowersList(v))
	})

	/* same as "GetFollowersList"
	t.AddCommand("GetFollowersListAll", func(params map[string]interface{}) interface{} {
	  v := url.Values{}

	  //t.twitterApi.GetFollowersListAll(v url.Values) (result chan FollowersPage)
	  result := t.twitterApi.GetFollowersListAll(v)
	  return nil
	})
	*/

	t.AddCommand("GetFriendsIds", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetFriendsIds(v url.Values) (c Cursor, err error)
		return resultApi(t.twitterApi.GetFriendsIds(v))
	})

	t.AddCommand("GetFriendsIdsAll", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetFriendsIdsAll(v url.Values) (c Cursor, err error)
		return resultApi(t.twitterApi.GetFriendsIdsAll(v))
	})

	t.AddCommand("GetFriendsList", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetFriendsList(v url.Values) (c UserCursor, err error)
		return resultApi(t.twitterApi.GetFriendsList(v))
	})

	t.AddCommand("GetFriendshipsIncoming", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetFriendshipsIncoming(v url.Values) (c Cursor, err error)
		return resultApi(t.twitterApi.GetFriendshipsIncoming(v))
	})

	t.AddCommand("GetFriendshipsLookup", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetFriendshipsLookup(v url.Values) (friendships []Friendship, err error)
		return resultApi(t.twitterApi.GetFriendshipsLookup(v))
	})

	t.AddCommand("GetFriendshipsNoRetweets", func(params map[string]interface{}) interface{} {
		//t.twitterApi.GetFriendshipsNoRetweets() (ids []int64, err error)
		return resultApi(t.twitterApi.GetFriendshipsNoRetweets())
	})

	t.AddCommand("GetFriendshipsOutgoing", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetFriendshipsOutgoing(v url.Values) (c Cursor, err error)
		return resultApi(t.twitterApi.GetFriendshipsOutgoing(v))
	})

	t.AddCommand("GetHomeTimeline", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetHomeTimeline(v url.Values) (timeline []Tweet, err error)
		return resultApi(t.twitterApi.GetHomeTimeline(v))
	})

	t.AddCommand("GetMentionsTimeline", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetMentionsTimeline(v url.Values) (timeline []Tweet, err error)
		return resultApi(t.twitterApi.GetMentionsTimeline(v))
	})

	t.AddCommand("GetOEmbed", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetOEmbed(v url.Values) (o OEmbed, err error)
		return resultApi(t.twitterApi.GetOEmbed(v))
	})

	t.AddCommand("GetOEmbedId", func(params map[string]interface{}) interface{} {
		id := params["id"].(float64)
		v := url.Values{}

		//t.twitterApi.GetOEmbedId(id int64, v url.Values) (o OEmbed, err error)
		return resultApi(t.twitterApi.GetOEmbedId(int64(id), v))
	})

	t.AddCommand("GetRetweets", func(params map[string]interface{}) interface{} {
		id := params["id"].(float64)
		v := url.Values{}

		//t.twitterApi.GetRetweets(id int64, v url.Values) (tweets []Tweet, err error)
		return resultApi(t.twitterApi.GetRetweets(int64(id), v))
	})

	t.AddCommand("GetRetweetsOfMe", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetRetweetsOfMe(v url.Values) (tweets []Tweet, err error)
		return resultApi(t.twitterApi.GetRetweetsOfMe(v))
	})

	t.AddCommand("GetSearch", func(params map[string]interface{}) interface{} {
		queryString := params["queryString"].(string)
		v := url.Values{}

		//t.twitterApi.GetSearch(queryString string, v url.Values) (timeline []Tweet, err error)
		return resultApi(t.twitterApi.GetSearch(queryString, v))
	})

	t.AddCommand("GetSelf", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetSelf(v url.Values) (u User, err error)
		return resultApi(t.twitterApi.GetSelf(v))
	})

	t.AddCommand("GetTweet", func(params map[string]interface{}) interface{} {
		id := params["id"].(float64)
		v := url.Values{}

		//t.twitterApi.GetTweet(id int64, v url.Values) (tweet Tweet, err error)
		return resultApi(t.twitterApi.GetTweet(int64(id), v))
	})

	t.AddCommand("GetUserSearch", func(params map[string]interface{}) interface{} {
		searchTerm := params["searchTerm"].(string)
		v := url.Values{}

		//t.twitterApi.GetUserSearch(searchTerm string, v url.Values) (u []User, err error)
		return resultApi(t.twitterApi.GetUserSearch(searchTerm, v))
	})

	t.AddCommand("GetUserTimeline", func(params map[string]interface{}) interface{} {
		v := url.Values{}

		//t.twitterApi.GetUserTimeline(v url.Values) (timeline []Tweet, err error)
		return resultApi(t.twitterApi.GetUserTimeline(v))
	})

	t.AddCommand("GetUsersLookup", func(params map[string]interface{}) interface{} {
		usernames := params["usernames"].(string)
		v := url.Values{}

		//t.twitterApi.GetUsersLookup(usernames string, v url.Values) (u []User, err error)
		return resultApi(t.twitterApi.GetUsersLookup(usernames, v))
	})

	t.AddCommand("GetUsersLookupByIds", func(params map[string]interface{}) interface{} {
		s := params["ids"].(string)
		idsStr := strings.Split(strings.Replace(s, " ", "", -1), ",")
		v := url.Values{}

		ids := []int64{}

		for _, idStr := range idsStr {
			id, err := strconv.Atoi(idStr)
			if err == nil {
				ids = append(ids, int64(id))
			}
		}

		//t.twitterApi.GetUsersLookupByIds(ids []int64, v url.Values) (u []User, err error)
		return resultApi(t.twitterApi.GetUsersLookupByIds(ids, v))
	})

	t.AddCommand("GetUsersShow", func(params map[string]interface{}) interface{} {
		username := params["username"].(string)
		v := url.Values{}

		//t.twitterApi.GetUsersShow(username string, v url.Values) (u User, err error)
		return resultApi(t.twitterApi.GetUsersShow(username, v))
	})

	t.AddCommand("GetUsersShowById", func(params map[string]interface{}) interface{} {
		id := params["id"].(float64)
		v := url.Values{}

		//t.twitterApi.GetUsersShowById(id int64, v url.Values) (u User, err error)
		return resultApi(t.twitterApi.GetUsersShowById(int64(id), v))
	})

	t.AddCommand("PostTweet", func(params map[string]interface{}) interface{} {
		status := params["status"].(string)
		v := url.Values{}

		//t.twitterApi.PostTweet(status string, v url.Values) (tweet Tweet, err error)
		return resultApi(t.twitterApi.PostTweet(status, v))
	})

	t.AddCommand("ReturnRateLimitError", func(params map[string]interface{}) interface{} {
		enable := params["enable"].(bool)

		//t.twitterApi.ReturnRateLimitError(b bool)
		t.twitterApi.ReturnRateLimitError(enable)
		return nil
	})

	t.AddCommand("Retweet", func(params map[string]interface{}) interface{} {
		id := params["id"].(float64)
		trimUser := params["trimUser"].(bool)

		//t.twitterApi.Retweet(id int64, trimUser bool) (rt Tweet, err error)
		return resultApi(t.twitterApi.Retweet(int64(id), trimUser))
	})

	t.AddCommand("SetDelay", func(params map[string]interface{}) interface{} {
		delay := params["delay"].(float64) // ms

		//t.twitterApi.SetDelay(t time.Duration)
		t.twitterApi.SetDelay(time.Duration(delay) * time.Millisecond)
		return nil
	})

	t.AddCommand("VerifyCredentials", func(params map[string]interface{}) interface{} {
		//t.twitterApi.VerifyCredentials() (ok bool, err error)
		return resultApi(t.twitterApi.VerifyCredentials())
	})

	return t
}

func (t *TwitterDriver) Start() bool {
	return true
}

func (t *TwitterDriver) Halt() bool {
	t.twitterApi.Close()
	return true
}

func (t *TwitterDriver) TwitterApi() *anaconda.TwitterApi {
	return t.twitterApi
}

func (t *TwitterDriver) SetAccessToken(accessToken string, accessTokenSecret string) {
	t.twitterApi = anaconda.NewTwitterApi(accessToken, accessTokenSecret)
}

func resultApi(v interface{}, err error) interface{} {
	if err == nil {
		return struct {
			Result interface{} `json:"result"`
		}{v}
	} else {
		return struct {
			Result error `json:"result"`
		}{err}
	}
}

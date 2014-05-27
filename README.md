tweet
============

## Feature
- Roomba-Hack
- Docomo雑談APIから取得したtextをTweetする
- Twitter API 1.1
- 非公開設定 https://support.twitter.com/articles/20169930-

## submodule
<pre>
[submodule "tokenbucket"]
	path = tokenbucket
	url = https://github.com/ChimeraCoder/tokenbucket
[submodule "go-oauth"]
	path = go-oauth
	url = https://github.com/garyburd/go-oauth
[submodule "anaconda"]
	path = anaconda
	url = https://github.com/ChimeraCoder/anaconda
</pre>

## config
<pre>
{
    "ConsumerKey": "xxxxxxxx",
    "ConsumerSecret": "xxxxxxxx",
    "AccessToken": "xxxxxxxx",
    "AccessTokenSecret": "xxxxxxxx"
}
</pre>

## Usage
<pre>
/* search tweet*/
search,err := twitter.Search("fisproject", "./roomba2d2.json")
if err != nil {
  fmt.Println(err)
}
fmt.Println(search)

/* post tweet*/
err = twitter.Tweet("post-test", "./roomba2d2.json")
if err != nil {
  fmt.Println(err)
}
</pre>

## Result
<pre>
./tweet_example
【FQDN】DNSの基礎を整理した【ICANN】 - http://t.co/UOAdPmHg5o【機械学習】PerceptronをGo言語で書いてみた【線形分類器】 - http://t.co/eOP8bcpXzV
{[] <nil> Sun May 25 09:17:10 +0000 2014 {[] [] [] []} 0 false <nil> 470493769389731841 470493769389731841  0  0  {map[] {[] } []    {[] }    [] } false 0 false <nil> <a href="http://127.0.0.1/" rel="nofollow">roomba2d2</a> post-test false {false Thu Nov 24 05:51:10 +0000 2011 true true  0 false 0 false 0 false 420103191 420103191 false ja 0  whitecat false C0DEED http://abs.twimg.com/images/themes/theme1/bg.png https://abs.twimg.com/images/themes/theme1/bg.png false http://abs.twimg.com/sticky/default_profile_images/default_profile_4_normal.png https://abs.twimg.com/sticky/default_profile_images/default_profile_4_normal.png 0084B4 C0DEED DDEEF6 333333 true true roomba2d2 false <nil> 3 Tokyo  32400 false}}
true
</pre>

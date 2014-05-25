tweet
============

## Feature
- Roomba-Hack
- Docomo雑談APIから取得したtextをTweetする
- Twitter API 1.1
- Using package Anaconda https://godoc.org/github.com/ChimeraCoder/anaconda

## Account
- テスト用アカウントを使用 (ID : roomba2d2)
- 非公開設定 https://support.twitter.com/articles/20169930-

## Usage
<pre>
/* search tweet*/
search := twitter.Search("fisproject")
fmt.Println(search)

/* post tweet*/
tweet := twitter.Tweet("post-test")
fmt.Println(tweet) 
</pre>

## Result
<pre>
./tweet_example
【FQDN】DNSの基礎を整理した【ICANN】 - http://t.co/UOAdPmHg5o【機械学習】PerceptronをGo言語で書いてみた【線形分類器】 - http://t.co/eOP8bcpXzV
{[] <nil> Sun May 25 09:17:10 +0000 2014 {[] [] [] []} 0 false <nil> 470493769389731841 470493769389731841  0  0  {map[] {[] } []    {[] }    [] } false 0 false <nil> <a href="http://127.0.0.1/" rel="nofollow">roomba2d2</a> post-test false {false Thu Nov 24 05:51:10 +0000 2011 true true  0 false 0 false 0 false 420103191 420103191 false ja 0  whitecat false C0DEED http://abs.twimg.com/images/themes/theme1/bg.png https://abs.twimg.com/images/themes/theme1/bg.png false http://abs.twimg.com/sticky/default_profile_images/default_profile_4_normal.png https://abs.twimg.com/sticky/default_profile_images/default_profile_4_normal.png 0084B4 C0DEED DDEEF6 333333 true true roomba2d2 false <nil> 3 Tokyo  32400 false}}
true
</pre>
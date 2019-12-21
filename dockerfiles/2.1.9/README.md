# twint-docker

<!-- ToC start -->
## Table of Contents
1. [How to use](#how-to-use)
  1. [First use](#first-use)
     1. [Create an alias](#create-an-alias)
1. [Docker-Compose](#docker-compose)
     1. [Elasticsearch and Kibana](#elasticsearch-and-kibana)
     1. [Tor proxy](#tor-proxy)
     1. [Execute Twint command](#execute-twint-command)
     1. [Examples of command](#examples-of-command)
  1. [Datas](#datas)
1. [Authors](#authors)
1. [🤝 Contributing](#-contributing)
1. [Show your support](#show-your-support)
<!-- ToC end -->

## How to use

### First use

For first usage, you need to build image docker.

```shell
docker pull x0rzkov/twint:2.1.9
```

or 

```shell
docker-compose build
```

#### Create an alias
```shell
alias twint="docker run -ti --rm -v $(pwd)/data:/opt/app/data x0rzkov/twint:2.1.9"
```               

### Elasticsearch and Kibana

Start to up elaticsearch and kibana

```shell
docker-compose up -d elasticsearch kibana
```

### Tor proxy

Start to up tor

```shell
docker-compose up -d tor
```

### Twint Search interface

Start to up elasticsearch and twint-search

```shell
docker-compose up -d twint-search elasticsearch
```

Then open in your browser [http://localhost:3000](http://localhost:3000)

### Execute Twint command

```shell
docker-compose run -v $PWD/twint:/opt/app/data twint -h
```

### Examples of command

A few simple examples to help you understand the basics:

```shell
docker-compose run -v $PWD/twint:/opt/app/data twint -u username - Scrape all the Tweets from user's timeline.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username -s pineapple - Scrape all Tweets from the user's timeline containing pineapple.
docker-compose run -v $PWD/twint:/opt/app/data twint -s pineapple - Collect every Tweet containing pineapple from everyone's Tweets.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username --year 2014 - Collect Tweets that were tweeted before 2014.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username --since 2015-12-20 - Collect Tweets that were tweeted since 2015-12-20.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username -o file.txt - Scrape Tweets and save to file.txt.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username -o file.csv --csv - Scrape Tweets and save as a csv file.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username --email --phone - Show Tweets that might have phone numbers or email addresses.
docker-compose run -v $PWD/twint:/opt/app/data twint -s "Donald Trump" --verified - Display Tweets by verified users that Tweeted about Donald Trump.
docker-compose run -v $PWD/twint:/opt/app/data twint -g="48.880048,2.385939,1km" -o file.csv --csv - Scrape Tweets from a radius of 1km around a place in Paris and export them docker-compose run -v $PWD/twint:/opt/app/data twint to a csv file.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username -es localhost:9200 - Output Tweets to Elasticsearch
docker-compose run -v $PWD/twint:/opt/app/data twint -u username -o file.json --json - Scrape Tweets and save as a json file.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username --database tweets.db - Save Tweets to a SQLite database.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username --followers - Scrape a Twitter user's followers.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username --following - Scrape who a Twitter user follows.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username --favorites - Collect all the Tweets a user has favorited.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username --following --user-full - Collect full user information a person follows
docker-compose run -v $PWD/twint:/opt/app/data twint -u username --profile-full - Use a slow, but effective method to gather Tweets from a user's profile (Gathers ~3200 Tweets, docker-compose run -v $PWD/twint:/opt/app/data twint Including Retweets).
docker-compose run -v $PWD/twint:/opt/app/data twint -u username --retweets - Use a quick method to gather the last 900 Tweets (that includes retweets) from a user's profile.
docker-compose run -v $PWD/twint:/opt/app/data twint -u username --resume 10940389583058 - Resume a search starting from the specified Tweet ID.
```

### Datas

For datas generate by twint, you can found result on folder twint

Let's play now :)

## Authors

👤 **x0rzkov**
* Github: [@x0rzkov](https://github.com/x0rzkov)

👤 **sebastienhouzet**
* Github: [@sebastienhouzet](https://github.com/sebastienhouzet)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/x0rzkov/twint-docker/issues).
See [`./docs/CONTRIBUTING.md`](https://github.com/x0rzkov/twint-dockers/blob/master/docs/CONTRIBUTING.md) for details.

## Show your support

Give a ⭐️ if this project helped you!
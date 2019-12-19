# twint-docker [![Build Status](https://travis-ci.com/x0rzkov/twint-docker.svg?branch=alpine)](https://travis-ci.com/x0rzkov/twint-docker) [![Build Status](https://cloud.drone.io/api/badges/x0rzkov/twint-docker/status.svg?ref=refs/heads/alpine)](https://cloud.drone.io/x0rzkov/twint-docker)

[![github all releases](https://img.shields.io/github/downloads/x0rzkov/twint-docker/total.svg)](https://github.com/x0rzkov/twint-docker) [![github latest release](https://img.shields.io/github/downloads/x0rzkov/twint-docker/latest/total.svg)](https://github.com/x0rzkov/twint-docker) [![github tag](https://img.shields.io/github/tag/x0rzkov/twint-docker.svg)](https://github.com/x0rzkov/twint-docker) [![github release](https://img.shields.io/github/release/x0rzkov/twint-docker.svg)](https://github.com/x0rzkov/twint-docker) [![github pre release](https://img.shields.io/github/release/x0rzkov/twint-docker/all.svg)](https://github.com/x0rzkov/twint-docker) [![github fork](https://img.shields.io/github/forks/x0rzkov/twint-docker.svg?style=social&label=Fork)](https://github.com/x0rzkov/twint-docker) [![github stars](https://img.shields.io/github/stars/x0rzkov/twint-docker.svg?style=social&label=Star)](https://github.com/x0rzkov/twint-docker) [![github watchers](https://img.shields.io/github/watchers/x0rzkov/twint-docker.svg?style=social&label=Watch)](https://github.com/x0rzkov/twint-docker) [![github open issues](https://img.shields.io/github/issues/x0rzkov/twint-docker.svg)](https://github.com/x0rzkov/twint-docker) [![github closed issues](https://img.shields.io/github/issues-closed/x0rzkov/twint-docker.svg)](https://github.com/x0rzkov/twint-docker) [![github open pr](https://img.shields.io/github/issues-pr/x0rzkov/twint-docker.svg)](https://github.com/x0rzkov/twint-docker) [![github closed pr](https://img.shields.io/github/issues-pr-closed/x0rzkov/twint-docker.svg)](https://github.com/x0rzkov/twint-docker) [![github contributors](https://img.shields.io/github/contributors/x0rzkov/twint-docker.svg)](https://github.com/x0rzkov/twint-docker) [![github license](https://img.shields.io/github/license/x0rzkov/twint-docker.svg)](https://github.com/x0rzkov/twint-docker) [![gitter chat room](https://badges.gitter.im/x0rzkov/twint-docker.svg)](https://gitter.im/x0rzkov/twint-docker) [![travis badge](https://img.shields.io/travis/x0rzkov/twint-docker.svg)](https://travis-ci.com/x0rzkov/twint-docker) [![Codecov](https://img.shields.io/codecov/c/github/x0rzkov/twint-docker.svg)](https://codecov.io/gh/x0rzkov/twint-docker) [![Coveralls](https://img.shields.io/coveralls/x0rzkov/twint-docker.svg)](https://coveralls.io/github/x0rzkov/twint-docker) [![Code Climate](https://img.shields.io/codeclimate/github/x0rzkov/twint-docker.svg)](https://codeclimate.com/github/x0rzkov/twint-docker) [![Code Climate](https://img.shields.io/codeclimate/coverage/github/x0rzkov/twint-docker.svg)](https://codeclimate.com/github/x0rzkov/twint-docker/coverage) [![Code Climate](https://img.shields.io/codeclimate/issues/github/x0rzkov/twint-docker.svg)](https://codeclimate.com/github/x0rzkov/twint-docker/issues)

## How to use

### First use

For fisrt usage, you need to build image docker.

``` bash
docker-compose build
```

#### Create an alias
``` bash
alias twint="docker run -ti --rm -v $(pwd)/data:/opt/app/data x0rzkov/twint:latest-alpine"
```               

### Elasticsearch and Kibana

Start to up elaticsearch and kibana

``` bash
docker-compose up -d elasticsearch kibana
```

### Execute Twint command

``` bash
docker-compose run -v $PWD/twint:/srv/twint twint {{CMD TWINT}}
```

### Examples of command

A few simple examples to help you understand the basics:

``` bash
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username - Scrape all the Tweets from user's timeline.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username -s pineapple - Scrape all Tweets from the user's timeline containing pineapple.
docker-compose run -v $PWD/twint:/srv/twint twint twint -s pineapple - Collect every Tweet containing pineapple from everyone's Tweets.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username --year 2014 - Collect Tweets that were tweeted before 2014.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username --since 2015-12-20 - Collect Tweets that were tweeted since 2015-12-20.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username -o file.txt - Scrape Tweets and save to file.txt.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username -o file.csv --csv - Scrape Tweets and save as a csv file.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username --email --phone - Show Tweets that might have phone numbers or email addresses.
docker-compose run -v $PWD/twint:/srv/twint twint twint -s "Donald Trump" --verified - Display Tweets by verified users that Tweeted about Donald Trump.
docker-compose run -v $PWD/twint:/srv/twint twint twint -g="48.880048,2.385939,1km" -o file.csv --csv - Scrape Tweets from a radius of 1km around a place in Paris and export them docker-compose run -v $PWD/twint:/srv/twint twint to a csv file.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username -es localhost:9200 - Output Tweets to Elasticsearch
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username -o file.json --json - Scrape Tweets and save as a json file.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username --database tweets.db - Save Tweets to a SQLite database.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username --followers - Scrape a Twitter user's followers.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username --following - Scrape who a Twitter user follows.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username --favorites - Collect all the Tweets a user has favorited.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username --following --user-full - Collect full user information a person follows
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username --profile-full - Use a slow, but effective method to gather Tweets from a user's profile (Gathers ~3200 Tweets, docker-compose run -v $PWD/twint:/srv/twint twint Including Retweets).
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username --retweets - Use a quick method to gather the last 900 Tweets (that includes retweets) from a user's profile.
docker-compose run -v $PWD/twint:/srv/twint twint twint -u username --resume 10940389583058 - Resume a search starting from the specified Tweet ID.
```

### Datas

For datas generate by twint, you can found result on folder twint

Let's play now :)
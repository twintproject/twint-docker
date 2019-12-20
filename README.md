# How to use

## First use

For fisrt usage, you need to build image docker.

``` bash
docker-compose build
```

## Quick start
```
docker network create nw_twint
docker-compose up -d elasticsearch
sleep 10
docker-compose up -d searchapp
docker-compose run twint -u noneprivacy -es elasticsearch:9200 --json -o /opt/twint/noneprivacy.json
open http://localhost:3000
```

## Twint, Elasticsearch & Searchapp

Start to up elaticsearch and searchapp

``` bash
docker network create nw_twint
docker-compose up -d elasticsearch searchapp twint

```

## Execute Twint command

``` bash
docker network create nw_twint
docker-compose run -v $PWD/output:/srv/twint twint {{CMD TWINT}}
```

## Examples of command

A few simple examples to help you understand the basics:

``` bash
docker-compose run twint -u username -es elasticsearch:9200
docker-compose run twint -s "#osint" -es elasticsearch:9200
docker-compose run twint -u username -es elasticsearch:9200 --json -o /opt/twint/username.json
USERNAME=username docker-compose run twint -u ${USERNAME} -es elasticsearch:9200 --json -o /opt/twint/${USERNAME}.json
```

if local install of twint
``` bash
twint -u username -es localhost:9200
twint -u username -es localhost:9200 --json -o /opt/twint/username.json
```

## Search engine

- Allows to do faceted search
- Current build is developpment so you can do change and re-compile in real-time.

```
open http://localhost:3000
```

## Debug Elasticsearch

```
open http://localhost:9000
open http://localhost:9200/twinttweets/_search?pretty=true&q=*:*
open http://localhost:9200/twinttweets/_count?pretty
```

### Screenshots
![alt text](https://github.com/lucmski/twint-search/raw/master/docs/screenshot1.png "Screenshot #1")

![alt text](https://github.com/lucmski/twint-search/raw/master/docs/screenshot2.png "Screenshot #2")

## Known Issues :warning:

I have noticed when running the new **5.0+** version on a linux host you need to increase the memory map areas with the following command

``` bash
sudo sysctl -w vm.max_map_count=262144
```

More at https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html

## Issues

Find a bug? Want more features? Find something missing in the documentation? Let me know! Please don't hesitate to [file an issue](https://github.com/blacktop/docker-elasticsearch-alpine/issues/new)

## To do
- Embed video in tweets
- Tweet with images
- Tweet with location
- Filter re-tweets
- All the most crazy things possible

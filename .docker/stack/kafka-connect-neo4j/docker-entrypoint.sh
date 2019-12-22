#!/bin/sh

set -x
set -e

python twint-graph.py

curl -i -X POST -H "Accept:application/json" \
    -H  "Content-Type:application/json" http://localhost:8083/connectors/ \
    -d '{
      "name": "connect.sink.neo4j.tweets",
      "config": {
        "topics": "tweets10",
        "connector.class": "streams.kafka.connect.sink.Neo4jSinkConnector",
        "neo4j.server.uri": "bolt://neo4j:7687",
        "neo4j.authentication.basic.username": "neo4j",
        "neo4j.authentication.basic.password": "neo",
        "neo4j.topic.cypher.tweets10": "WITH event AS data MERGE (t:Tweet {id: data.id}) SET t.text = data.tweet, t.createdAt = datetime({epochmillis:data.datetime}) MERGE (u:User {username: 
data.username}) SET u.id = data.user_id   MERGE (u)-[:POSTED]->(t) FOREACH (ht IN data.hashtags | MERGE (hashtag:HashTag {value: ht}) MERGE (t)-[:HAS_HASHTAG]->(hashtag))"
      }
    }'


FROM x0rzkov/twint:v2.1.10-alpine3.10

MAINTAINER x0rzkov@protonmail.com

RUN apk add --no-cache librdkafka-dev gcc nano curl

COPY requirements.txt /opt/app/requirements.txt
WORKDIR /opt/app/

RUN pip3 install -r requirements.txt

COPY twint-graph.py /opt/app/twint-graph.py
COPY docker-entrypoint.sh /opt/app/entrypoint.sh

ENTRYPOINT ["/opt/app/entrypoint.sh"]

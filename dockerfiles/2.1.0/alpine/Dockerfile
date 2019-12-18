FROM alpine:3.10

MAINTAINER x0rxkov@protonmail.com

ARG TWINT_GID=997
ARG TWINT_UID=997

RUN addgroup -g 997 twint && \
    adduser -u 997 -D -h /opt/twint -s /bin/sh -G twint twint

# This hack is widely applied to avoid python printing issues in docker containers.
# See: https://github.com/Docker-Hub-frolvlad/docker-alpine-python3/pull/13
ENV PYTHONUNBUFFERED=1

RUN echo "**** install Python ****" && \
    apk add --no-cache python3 sqlite sqlite-dev git ca-certificates cython openblas-dev musl-dev python3-dev libffi-dev gcc g++ && \
    if [ ! -e /usr/bin/python ]; then ln -sf python3 /usr/bin/python ; fi && \
    \
    echo "**** install pip ****" && \
    python3 -m ensurepip && \
    rm -r /usr/lib/python*/ensurepip && \
    pip3 install --no-cache --upgrade pip setuptools wheel && \
    if [ ! -e /usr/bin/pip ]; then ln -s pip3 /usr/bin/pip ; fi

WORKDIR /opt/twint

RUN git clone --depth=1 -b v2.1.0 https://github.com/twintproject/twint /opt/twint \
	&& cd /opt/twint \
	&& pip install -e .

WORKDIR /opt/twint

ENTRYPOINT ["twint"]

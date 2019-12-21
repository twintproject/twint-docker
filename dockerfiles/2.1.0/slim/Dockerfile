FROM python:3.7-slim-stretch

MAINTAINER x0rxkov <x0rxkov@protonmail.com>

ARG TWINT_VERSION=v2.1.0

COPY docker-entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

RUN \
apt-get update && \
apt-get install -y \
git

RUN \
pip3 install --upgrade -e git+https://github.com/twintproject/twint.git@v2.1.0#egg=twint

RUN \
apt-get clean autoclean && \
rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENTRYPOINT ["/entrypoint.sh"]
VOLUME /twint
WORKDIR /opt/twint/data

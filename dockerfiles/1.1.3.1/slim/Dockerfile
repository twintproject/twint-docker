FROM debian:stretch-slim

MAINTAINER x0rxkov <x0rxkov@protonmail.com>

ARG TWINT_VERSION=1.1.3.1

COPY docker-entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

RUN \
apt-get update && \
apt-get install -y \
git \
python3-pip

RUN \
pip3 install --upgrade -e git+https://github.com/twintproject/twint.git@1.1.3.1#egg=twint

RUN \
apt-get clean autoclean && \
rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENTRYPOINT ["/entrypoint.sh"]
VOLUME /twint
WORKDIR /srv/twint
FROM python:3.6-alpine3.9
LABEL maintainer="Farshad Nematdoust <farshad@nematdoust.com>"
RUN apk add --no-cache --no-progress git subversion gcc python-dev build-base libffi-dev
RUN pip3 install --upgrade -e git+https://github.com/twintproject/twint.git@origin/master#egg=twint
RUN rm -rf /root/.cache
ENTRYPOINT [ "twint" ]


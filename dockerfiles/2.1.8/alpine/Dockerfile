FROM alpine:3.10 AS build

WORKDIR /opt/app

# Install Python and external dependencies, including headers and GCC
RUN apk add --no-cache python3 python3-dev py3-pip libffi libffi-dev musl-dev gcc git ca-certificates openblas-dev musl-dev g++

# Install Pipenv
RUN pip3 install pipenv

# Create a virtual environment and activate it
RUN python3 -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH" \
	VIRTUAL_ENV="/opt/venv"

# Install dependencies into the virtual environment with Pipenv
RUN git clone --depth=1 -b v2.1.8 https://github.com/twintproject/twint /opt/app \
	&& cd /opt/app \
	&& pip3 install --upgrade pip \
	&& pip3 install cython \
	&& pip3 install numpy \
	&& pip3 install .

FROM alpine:3.10
MAINTAINER x0rxkov@protonmail.com

WORKDIR /opt/app

# Install Python and external runtime dependencies only
RUN apk add --no-cache python3 libffi openblas libstdc++

# Copy the virtual environment from the previous image
COPY --from=build /opt/venv /opt/venv

# Activate the virtual environment
ENV PATH="/opt/venv/bin:$PATH" \
	VIRTUAL_ENV="/opt/venv"

# Copy your application
# COPY . /opt/app/
WORKDIR /opt/app

ENTRYPOINT ["twint"]
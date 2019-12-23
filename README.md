# twint-docker based on Alpine, Ubuntu and Debian Slim

<p align="center">
    <a href="https://travis-ci.com/twintproject/twint-docker"><img src="https://img.shields.io/travis/twintproject/twint-docker.svg" /></a>
    <a href="https://cloud.drone.io/twintproject/twint-docker"><img src="https://cloud.drone.io/api/badges/twintproject/twint-docker/status.svg?ref=refs/heads/master" /></a>
</p>

<p align="center">
    <a href="https://github.com/twintproject/twint-docker" alt="github all releases"><img src="https://img.shields.io/github/downloads/twintproject/twint-docker/total.svg" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github latest release"><img src="https://img.shields.io/github/downloads/twintproject/twint-docker/latest/total.svg" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github tag"><img src="https://img.shields.io/github/tag/twintproject/twint-docker.svg" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github release"><img src="https://img.shields.io/github/release/twintproject/twint-docker.svg" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github pre release"><img src="https://img.shields.io/github/release/twintproject/twint-docker/all.svg" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github fork"><img src="https://img.shields.io/github/forks/twintproject/twint-docker.svg?style=social&label=Fork" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github stars"><img src="https://img.shields.io/github/stars/twintproject/twint-docker.svg?style=social&label=Star" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github watchers"><img src="https://img.shields.io/github/watchers/twintproject/twint-docker.svg?style=social&label=Watch" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github open issues"><img src="https://img.shields.io/github/issues/twintproject/twint-docker.svg" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github closed issues"><img src="https://img.shields.io/github/issues-closed/twintproject/twint-docker.svg" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github open pr"><img src="https://img.shields.io/github/issues-pr/twintproject/twint-docker.svg" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github closed pr"><img src="https://img.shields.io/github/issues-pr-closed/twintproject/twint-docker.svg" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github contributors"><img src="https://img.shields.io/github/contributors/twintproject/twint-docker.svg" /></a>
    <a href="https://github.com/twintproject/twint-docker" alt="github license"><img src="https://img.shields.io/github/license/twintproject/twint-docker.svg" /></a>
    <a href="https://travis-ci.com/twintproject/twint-docker" alt="travis badge"><img src="https://img.shields.io/travis/twintproject/twint-docker.svg" /></a>
</p>

## Requirements
If you don't have Docker/Docker-Compose check **Setup Docker** section

<details>
<summary><b>Setup Docker</b></summary>
<p>

## Docker
macOS: <a href="https://docs.docker.com/docker-for-mac/install/"> https://docs.docker.com/docker-for-mac/install/ </a>

linux: <a href="https://docs.docker.com/install/linux/docker-ce/ubuntu/"> https://docs.docker.com/install/linux/docker-ce/ubuntu/ </a>

## Docker Compose

linux: <a href="https://docs.docker.com/compose/install/"> https://docs.docker.com/compose/install/ </a>
</p>
</details>

## How to use

For first usage, you need to build image docker.

```shell
git clone --depth=1 https://github.com//twintproject/twint-docker
cd ./dockerfiles/[[VERSION]]/[[OS]]
docker-compose up -d
docker-compose run twint -h
```

or 

```
docker pull x0rzkov/twint:[[TAG]]
docker run -ti --rm x0rzkov/twint:[[TAG]] -h
```

Then check the README.md for each versions.

### Available images
| Image   |      Size      |  Arch |  Os |  Link |
|----------|:-------------:|------|------|------|
| docker pull x0rzkov/twint:latest|**248 MB**|amd64|linux|[`./dockerfiles/latest`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/latest/)|
| docker pull x0rzkov/twint:latest-slim|**168 MB**|amd64|linux|[`./dockerfiles/latest/slim`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/latest/slim/)|
| docker pull x0rzkov/twint:latest-alpine|**81 MB**|amd64|linux|[`./dockerfiles/latest/alpine`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/latest/alpine/)|
| docker pull x0rzkov/twint:2.0.0|**247 MB**|amd64|linux|[`./dockerfiles/2.0.0`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.0.0/)|
| docker pull x0rzkov/twint:2.0.0-alpine|**81 MB**|amd64|linux|[`./dockerfiles/2.0.0/alpine`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.0.0/alpine/)|
| docker pull x0rzkov/twint:2.0.0-slim|**168 MB**|amd64|linux|[`./dockerfiles/2.0.0/slim`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.0.0/slim/)|
| docker pull x0rzkov/twint:2.1.0|**247 MB**|amd64|linux|[`./dockerfiles/2.1.0`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.0/)|
| docker pull x0rzkov/twint:2.1.0-alpine|**81 MB**|amd64|linux|[`./dockerfiles/2.1.0/alpine`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.0/alpine/)|
| docker pull x0rzkov/twint:2.1.0-slim|**168 MB**|amd64|linux|[`./dockerfiles/2.1.0/slim`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.0/slim/)|
| docker pull x0rzkov/twint:2.1.10|**248 MB**|amd64|linux|[`./dockerfiles/2.1.10`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.10/)|
| docker pull x0rzkov/twint:2.1.10-alpine|**81 MB**|amd64|linux|[`./dockerfiles/2.1.10/alpine`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.10/alpine/)|
| docker pull x0rzkov/twint:2.1.10-slim|**168 MB**|amd64|linux|[`./dockerfiles/2.1.10/slim`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.10/slim/)|
| docker pull x0rzkov/twint:2.1.11|**248 MB**|amd64|linux|[`./dockerfiles/2.1.11`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.11/)|
| docker pull x0rzkov/twint:2.1.11-alpine|**81 MB**|amd64|linux|[`./dockerfiles/2.1.11/alpine`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.11/alpine/)|
| docker pull x0rzkov/twint:2.1.11-slim|**168 MB**|amd64|linux|[`./dockerfiles/2.1.11/slim`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.11/slim/)|
| docker pull x0rzkov/twint:2.1.4|**247 MB**|amd64|linux|[`./dockerfiles/2.1.4`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.4/)|
| docker pull x0rzkov/twint:2.1.4-alpine|**81 MB**|amd64|linux|[`./dockerfiles/2.1.4/alpine`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.4/alpine/)|
| docker pull x0rzkov/twint:2.1.4-slim|**168 MB**|amd64|linux|[`./dockerfiles/2.1.4/slim`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.4/slim/)|
| docker pull x0rzkov/twint:2.1.6|**247 MB**|amd64|linux|[`./dockerfiles/2.1.6`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.6/)|
| docker pull x0rzkov/twint:2.1.6-alpine|**81 MB**|amd64|linux|[`./dockerfiles/2.1.6/alpine`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.6/alpine/)|
| docker pull x0rzkov/twint:2.1.6-slim|**168 MB**|amd64|linux|[`./dockerfiles/2.1.6/slim`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.6/slim/)|
| docker pull x0rzkov/twint:2.1.8|**247 MB**|amd64|linux|[`./dockerfiles/2.1.8`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.8/)|
| docker pull x0rzkov/twint:2.1.8-alpine|**81 MB**|amd64|linux|[`./dockerfiles/2.1.8/alpine`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.8/alpine/)|
| docker pull x0rzkov/twint:2.1.8-slim|**168 MB**|amd64|linux|[`./dockerfiles/2.1.8/slim`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.8/slim/)|
| docker pull x0rzkov/twint:2.1.9|**248 MB**|amd64|linux|[`./dockerfiles/2.1.9`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.9/)|
| docker pull x0rzkov/twint:2.1.9-alpine|**81 MB**|amd64|linux|[`./dockerfiles/2.1.9/alpine`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.9/alpine/)|
| docker pull x0rzkov/twint:2.1.9-slim|**168 MB**|amd64|linux|[`./dockerfiles/2.1.9/slim`](https://github.com//twintproject/twint-docker/tree/master/dockerfiles/2.1.9/slim/)|


## Authors

👤 **pielco11**
* Twitter: [@noneprivacy](https://twitter.com/noneprivacy) ![Twitter Follow](https://img.shields.io/twitter/follow/noneprivacy?label=Follow&style=social)
* Github: [@pielco11](https://github.com/pielco11)


👤 **x0rzkov**
* Twitter: [@x0rzkov](https://twitter.com/x0rzkov) ![Twitter Follow](https://img.shields.io/twitter/follow/x0rzkov?label=Follow&style=social)
* Github: [@x0rzkov](https://github.com/x0rzkov)
* Email: x0rzkov@protonmail.com

👤 **sebastienhouzet**
* Twitter: [@sebastienhouzet](https://twitter.com/sebastienhouzet) ![Twitter Follow](https://img.shields.io/twitter/follow/sebastienhouzet?label=Follow&style=social)
* Github: [@sebastienhouzet](https://github.com/sebastienhouzet)



## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com//twintproject/twint-docker/issues).
See [`./docs/CONTRIBUTING.md`](https://github.com//twintproject/twint-docker/tree/master/docs/CONTRIBUTING.md) for details.

## Show your support

Give a ⭐️ if this project helped you!


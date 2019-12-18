#!/bin/sh

BASE_DIR="$(dirname -- "`readlink -f -- "$0"`")"

cd -- "$BASE_DIR"
set -e
set -x

# subshell
PYTHONPATH="$BASE_DIR"
# TWINT_DIR="$BASE_DIR/twint"
ACTION="$1"

docker_build() {
    # Check if it is a git repository
    if [ ! -d .git ]; then
	echo "This is not Git repository"
	exit 1
    fi

    if [ ! -x "$(which git)" ]; then
	echo "git is not installed"
	exit 1
    fi

    if [ ! git remote get-url origin 2> /dev/null ]; then
	echo "there is no remote origin"
	exit 1
    fi

    # This is a git repository

    # "git describe" to get the Docker version (for example : v0.15.0-89-g0585788e)
    # awk to remove the "v" and the "g"
    TWINT_GIT_VERSION=$(git describe --match "v[0-9]*\.[0-9]*\.[0-9]*" HEAD 2>/dev/null | awk -F'-' '{OFS="-"; $1=substr($1, 2); $3=substr($3, 2); print}')

    # add the suffix "-dirty" if the repository has uncommited change
    git update-index -q --refresh
    if [ ! -z "$(git diff-index --name-only HEAD --)" ]; then
	TWINT_GIT_VERSION="${TWINT_GIT_VERSION}-dirty"
    fi

    # Get the last git commit id, will be added to the Searx version (see Dockerfile)
    VERSION_GITCOMMIT=$(echo $TWINT_GIT_VERSION | cut -d- -f2-4)
    echo "Last commit : $VERSION_GITCOMMIT"

    # Check consistency between the git tag and the searx/version.py file
    # /!\ HACK : parse Python file with bash /!\
    # otherwise it is not possible build the docker image without all Python dependencies ( version.py loads __init__.py )
    # TWINT_PYTHON_VERSION=$(python -c "import six; import searx.version; six.print_(searx.version.VERSION_STRING)")
    TWINT_PYTHON_VERSION=$(cat searx/version.py | grep "\(VERSION_MAJOR\|VERSION_MINOR\|VERSION_BUILD\) =" | cut -d\= -f2 | sed -e 's/^[[:space:]]*//' | paste -sd "." -)
    if [ $(echo "$TWINT_GIT_VERSION" | cut -d- -f1) != "$TWINT_PYTHON_VERSION" ]; then
	echo "Inconsistency between the last git tag and the searx/version.py file"
	echo "git tag:          $TWINT_GIT_VERSION"
	echo "searx/version.py: $TWINT_PYTHON_VERSION"
	exit 1
    fi

    # define the docker image name
    # /!\ HACK to get the user name /!\
    GITHUB_USER=$(git remote get-url origin | sed 's/.*github\.com\/\([^\/]*\).*/\1/')
    TWINT_IMAGE_NAME="${GITHUB_USER:-searx}/searx"

    # build Docker image
    echo "Building image ${TWINT_IMAGE_NAME}:${TWINT_GIT_VERSION}"
    docker build \
        --build-arg TWINT_GIT_VERSION="${TWINT_GIT_VERSION}" \
        --build-arg VERSION_GITCOMMIT="${VERSION_GITCOMMIT}" \
        --build-arg LABEL_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
        --build-arg LABEL_VCS_REF=$(git rev-parse HEAD) \
        --build-arg LABEL_VCS_URL=$(git remote get-url origin) \
	    --build-arg TIMESTAMP_SETTINGS=$(git log -1 --format="%cd" --date=unix -- searx/settings.yml) \
	    --build-arg TIMESTAMP_UWSGI=$(git log -1 --format="%cd" --date=unix -- dockerfiles/uwsgi.ini) \
        -t ${TWINT_IMAGE_NAME}:latest -t ${TWINT_IMAGE_NAME}:${TWINT_GIT_VERSION} .

    if [ "$1" = "push" ]; then
	   docker push ${TWINT_IMAGE_NAME}:latest
	   docker push ${TWINT_IMAGE_NAME}:${TWINT_GIT_VERSION}
    fi
}

docker_build

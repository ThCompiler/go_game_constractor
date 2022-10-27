if [ -n "$1" ]; then
  VERSION=$1
else
  VERSION="v0.0.0"
fi

LATEST_TAG_REV=$(git rev-list --tags --max-count=1)
LATEST_COMMIT_REV=$(git rev-list HEAD --max-count=1)

if [ -n "$LATEST_TAG_REV" ]; then
    LATEST_TAG=$(git describe --tags "$(git rev-list --tags --max-count=1)")
else
    LATEST_TAG="v0.0.0"
fi

if [  "${VERSION}" = "${LATEST_TAG}" ]; then
    echo "$VERSION-$(git rev-list HEAD --max-count=1 --abbrev-commit)"
else
    echo "$VERSION"
fi
#!/bin/sh
set -e

test() {
  go test ./... -v --cover
}

lint() {
  GOIMPORTS_DIFF="$(goimports -d .)"

  if [ -n "$GOIMPORTS_DIFF" ] ; then
    echo -e "$GOIMPORTS_DIFF"
    exit 1
  fi
}

case "$1" in
  test) test ;;
  lint) lint ;;
  *)    exec "$@" ;;
esac

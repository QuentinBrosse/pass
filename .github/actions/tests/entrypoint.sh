#!/bin/sh
set -e

echo ">>> entrypoint"
ls -al

if [ $# -eq 0 ]
then
  go test ./... -v --cover
fi

exec "$@"

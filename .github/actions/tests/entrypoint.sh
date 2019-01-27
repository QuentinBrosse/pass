#!/bin/sh
set -e

if [ $# -eq 0 ]
then
  go test ./... -v --cover
fi

exec "$@"

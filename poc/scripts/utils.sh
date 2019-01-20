#!/usr/bin/env bash

##
# Echo on stderr
##
function echoerr ()
{
  echo "$@" >&2
}

##
# Colorize output
##
function color() {
  case $1 in
    yellow) echoerr -e -n "\033[33m"   ;;
    green)  echoerr -e -n "\033[32m"   ;;
    red)    echoerr -e -n "\033[0;31m" ;;
  esac
  echoerr "$2"
  echoerr -e -n "\033[0m"
}

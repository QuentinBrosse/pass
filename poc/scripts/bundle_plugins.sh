#!/usr/bin/env bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
PLUGINS_DIR="$ROOT_DIR/plugins"
PLUGINS_FILE="$PLUGINS_DIR/plugins.go"

source ${SCRIPT_DIR}/utils.sh

##
# Log message only if verbose mode is enabled
##
function log() {
  [ ! -z ${OPT_VERBOSE} ] && echo $@
}

##
# Print Usage
##
function usage() {
  color yellow "Usage:"
  echoerr "  ${BASH_SOURCE[0]} [OPTIONS]"
  echoerr ""

  color yellow "Options:"

  color green "  -c, --clear"
  echoerr -e "\tRemove bundle"

  echoerr ""
  exit $1;
}

OPT_BUNDLE=1
OPT_CLEAR=""

##
# Parse arguments
##
while [[ $# > 0 ]]
do
  case "$1" in
    -c|--clear) OPT_CLEAR=1 OPT_BUNDLE="" ;;
	  -h|--help) usage ;;
  esac
  shift
done

##
# Create Bundle
##
if [ ! -z ${OPT_BUNDLE} ] ; then
  echoerr "Create plugin bundle in $PLUGINS_FILE"
  pushd ${PLUGINS_DIR} > /dev/null
    printf "// Code generated by bundle_plugins. DO NOT EDIT.\n\n" > "$PLUGINS_FILE"
    printf "package plugins\n\n" >> "$PLUGINS_FILE"
    printf "var PluginsBundle = map[string]string{\n" >> "$PLUGINS_FILE"

    for file in $(find ${PLUGINS_DIR} -name "*.yml");
    do
      filename=$(basename -- "$file")
      filename="${filename%.*}"
      content=$(cat $file)
      printf "    \"$filename\": \`$content\`,\n\n" >> "$PLUGINS_FILE"
    done

    printf "}\n" >> "$PLUGINS_FILE"
  popd > /dev/null
fi

##
# Remove Bundle
##
if [ ! -z ${OPT_CLEAR} ] ; then
  echoerr "Remove plugin bundle in $PLUGINS_FILE"
  pushd ${PLUGINS_DIR} > /dev/null
    rm "$PLUGINS_FILE"
  popd > /dev/null
fi

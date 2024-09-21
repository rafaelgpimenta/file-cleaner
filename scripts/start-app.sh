#!/bin/bash

# Default variable values
keep_deps=false

# Function to display script usage
usage() {
 echo "Usage: $0 [OPTIONS]"
 echo "Options:"
 echo " -h, --help      Display this help message"
 echo " -k, --keep      Enable verbose mode"
}

has_argument() {
    [[ ("$1" == *=* && -n ${1#*=}) || ( ! -z "$2" && "$2" != -*)  ]];
}

extract_argument() {
  echo "${2:-${1#*=}}"
}

# Function to handle options and arguments
handle_options() {
  while [ $# -gt 0 ]; do
    case $1 in
      -h | --help)
        usage
        exit 0
        ;;
      -k | --keep)
        keep_deps=true
        ;;
      *)
        echo "Invalid option: $1" >&2
        usage
        exit 1
        ;;
    esac
    shift
  done
}

# Main script execution
handle_options "$@"

FOLDER="$( cd $(dirname "${BASH_SOURCE[0]}"); pwd )"
cd $FOLDER/..

docker compose -f docker-compose/docker-compose.yml up -d

export AWS_ACCESS_KEY_ID=foobar
export AWS_SECRET_ACCESS_KEY=foobar
go run cmd/cleaner/main.go | jq .

# Perform the desired actions based on the provided flags and arguments
if [ "$keep_deps" = false ]; then
 docker compose -f docker-compose/docker-compose.yml down
fi

# Ref.: https://medium.com/@wujido20/handling-flags-in-bash-scripts-4b06b4d0ed04

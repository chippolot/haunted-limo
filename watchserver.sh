#!/bin/bash
case "$OSTYPE" in
  darwin*|linux*) nodemon --delay 1.5 -x "go run ./cmd/server || exit 1" --signal SIGTERM -e go,html,gohtml,sh,env,json ;;
  msys*|cygwin*) nodemon --delay 1.5 -x "go run ./cmd/server || exit 1" --signal SIGKILL -e go,html,gohtml,sh,env,json ;;
  *)        echo "Cannot start script with unknown OSTYPE: $OSTYPE" ;;
esac
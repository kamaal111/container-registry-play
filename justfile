set export

PORT := "8080"
APPLICATION_NAME := "container-registery-play"

default: build run

run-dev:
    #!/bin/sh

    export SERVER_ADDRESS="127.0.0.1:$PORT"

    ~/go/bin/reflex -r '\.go' -s -- sh -c "go run ."

build:
    docker build -t $APPLICATION_NAME .

run:
    #!/bin/sh

    docker rm $APPLICATION_NAME
    docker run -dp $PORT:$PORT --name $APPLICATION_NAME -e PORT=$PORT $APPLICATION_NAME

tag-image:
    docker tag $APPLICATION_NAME "gcr.io/$GCP_PROJECT_ID/$APPLICATION_NAME:$COMMIT_SHA"

push-image:
    docker push "gcr.io/$GCP_PROJECT_ID/$APPLICATION_NAME:$COMMIT_SHA"

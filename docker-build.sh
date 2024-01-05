#!/bin/bash

IMAGE_NAME="y4-lang"
CONTAINER_NAME="y4-lang-builder"

EXISTING_IMAGE=$(docker images -q $IMAGE_NAME:latest)
if [ ! -z "$EXISTING_IMAGE" ]; then
    echo "image $IMAGE_NAME:latest already exists. deleting..."
    docker rmi $EXISTING_IMAGE
fi

docker build -t $IMAGE_NAME .
CONTAINER_ID=$(docker run -d --name $CONTAINER_NAME $IMAGE_NAME)
echo "container is running with ID: $CONTAINER_ID"

echo "waiting for container to complete the task"
docker wait $CONTAINER_ID

docker cp $CONTAINER_ID:/app/build.zip .
echo "file copied to build.zip"

docker stop $CONTAINER_ID
docker rm $CONTAINER_ID

echo "container stopped and removed"

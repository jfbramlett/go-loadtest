#!/bin/sh

REPOSITORY=nwp/nwp-load-test
TAG=latest

if [ -z "$JENKINS_URL" ]; then
    echo "Not on jenkins, skipping cleanup"
else
    echo "Delete local image"
    docker rmi -f $REPOSITORY:$TAG
    docker rmi -f $AWS_URL/$REPOSITORY:$TAG
fi
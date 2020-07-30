#!/bin/sh

aws_login() {
  aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin $AWS_URL
}

AWS_URL=871451932206.dkr.ecr.us-east-1.amazonaws.com
REPOSITORY=nwp/nwp-load-test
TAG=latest

IMAGE=$AWS_URL/$REPOSITORY
LOAD_IMAGE=$AWS_URL/$LOAD_REPOSITORY

docker build --pull --no-cache -t $REPOSITORY:$TAG .

if [ -z "$JENKINS_URL" ]; then
    echo "Not on jenkins, skipping publish"
else
    aws_login

    docker tag $REPOSITORY:$TAG $AWS_URL/$REPOSITORY:$TAG

    echo "Create REPOSITORY $REPOSITORY if it doesn't exist"
    aws ecr describe-repositories --repository-names $REPOSITORY || aws ecr create-repository --repository-name $REPOSITORY

    echo "Push the image"
    docker push $AWS_URL/$REPOSITORY:$TAG
fi
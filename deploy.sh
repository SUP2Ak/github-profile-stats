#!/bin/bash

APP_DIR="/home/git/github-profile-stats"
cd $APP_DIR || exit
export $(cat .env | xargs)

echo "Pulling latest changes from the repository..."
git pull origin main
echo "Building the Docker image..."
docker build -t github-stats-api ./backend

if [ $(docker ps -q -f name=github-stats-api) ]; then
    echo "Stopping existing container..."
    docker stop github-stats-api
fi

if [ $(docker ps -aq -f name=github-stats-api) ]; then
    echo "Removing existing container..."
    docker rm github-stats-api
fi

echo "Starting the new container..."
docker run -d --network=my_network_github_api -p 10000:10000 --name github-stats-api -e GITHUB_TOKEN=$GITHUB_TOKEN github-stats-api

echo "Deployment completed!"

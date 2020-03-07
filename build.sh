#!/bin/sh

rm -f ./source.tar.gz
tar -z -c --exclude ".git" --exclude "docker*" --exclude ".idea" --exclude "data" --exclude "source.tar.gz" -f ./source.tar.gz ./

docker build -t stock-issuers-server . -f ./docker/server/Dockerfile

rm ./source.tar.gz

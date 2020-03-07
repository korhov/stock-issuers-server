#!/bin/sh

tar -z -c --exclude ".git" --exclude "docker*" --exclude ".idea" -f ./docker/server/source.tar.gz ./

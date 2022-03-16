#!/bin/bash

VER=0.0.1

docker build -t gm-center:$VER .
docker-compose up -d
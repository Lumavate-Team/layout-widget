#!/usr/bin/env bash

DATE=`date +%m-%-d-%Y-%H-%M-%S`
echo $DATE
docker build --no-cache --rm -t layoutwidget:dev$DATE .
docker service update --image layoutwidget:dev$DATE layoutwidget-$1

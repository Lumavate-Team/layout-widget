#!/usr/bin/env bash

DATE=`date +%m-%-d-%Y-%H-%M-%S`
echo $DATE
docker build --no-cache --rm -t layoutwidget:prod$DATE .
#docker service update --image layoutwidget:prod$DATE layoutwidget-$1

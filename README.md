## Build the container containers

Clone the repo, and from the root directory of the repo, run the following command.

```
docker build --no-cache --rm -t gobasewidget:1.0 .
```

This command will use the Dockerfile to build the image.  We are specifying that we
do not want to use cache when building containers and that we want to remove intermediate containers after a successful builld.

[Docker Build Options](https://docs.docker.com/engine/reference/commandline/build/)


## Run the container

After the container has been built, you can now run the container.  Run the following command.

```
docker run -d -p 5000:8080 --volume "$(pwd)"/widget:/go/src/widget gobasewidget:1.0
```

This command will run the container in detached mode.  It will map port 5000 on your machine to port 8080 on the container.
It will map the widget directory to the /go/src/widget directory inside the container.  This will will allow you to modify files in your local widget directory, and it will reload the process when the files change.  You should be able to go to http://localhost:5000

[Docker Run Options](https://docs.docker.com/engine/reference/commandline/run/)

## Check the logs

To stream the logs to your terminal, you first need to run

```
docker ps
```

which will result in something like

```
CONTAINER ID        IMAGE                  COMMAND                  CREATED             STATUS              PORTS                    NAMES
676f62d88565        gobasemac4:dev021418   "/bin/sh -c 'bee run'"   15 minutes ago      Up 16 minutes       0.0.0.0:5000->8080/tcp   dreamy_albattani
```

you can then use the Container ID to stream the logs

```
docker logs -f 676
```

## Run inside Thor

```
DOCKER_IP=`ifconfig | grep "inet " | grep -Fv 127.0.0.1 | awk '{print $2}'`
docker run --rm -d \
-e "PUBLIC_KEY=mIhuoMJh0jbA5W4pUUNK" \
-e "PRIVATE_KEY=LXycaMpw5BzgfhsS4ydNxGzJ36qMnPrQHI8u2x3wQCZCZyGtZ4sOQbkEWnHmVchZEa79a0Y3xK7IKCymSLkugyabbJUGuXfyuoKL" \
-e "HOST_IP=$DOCKER_IP" \
-e "WIDGET_PORT=8091" \
-e "HOST_PORT=80" \
--name=thor \
-p 80:4201 \
quay.io/lumavate/thor:latest

docker run -d --rm \
  --volume "$(pwd)"/widget:/go/src/widget:rw \
  -e "PUBLIC_KEY=mIhuoMJh0jbA5W4pUUNK" \
  -e "PRIVATE_KEY=LXycaMpw5BzgfhsS4ydNxGzJ36qMnPrQHI8u2x3wQCZCZyGtZ4sOQbkEWnHmVchZEa79a0Y3xK7IKCymSLkugyabbJUGuXfyuoKL" \
  -e "BASE_URL=http://$DOCKER_IP" \
  -e "WIDGET_URL_PREFIX=/ic/widget/" \
  -e "PROTO=http://" \
  --name=widget-base-go \
  -p 8091:8080 \
  quay.io/lumavate/widget-base-go:latest
```

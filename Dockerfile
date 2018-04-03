FROM node:8.9-alpine as builder

RUN apk update && \
                apk add --no-cache \
                git \
                curl \
                openssh \
                libc-dev \
                go \
#&& git rev-parse HEAD > /revision \
#&& rm -rf .git \
  && mkdir -p /go/src/widget

ADD git.sh /git.sh
ENV GOPATH=/go
ENV PATH="/go/bin:${PATH}"

WORKDIR /go/src/widget
COPY ./widget /go/src/widget

RUN mkdir /root/.ssh/ && \
  touch /root/.ssh/known_hosts && \
  ssh-keyscan github.com >> /root/.ssh/known_hosts

ADD ims-go-components-rsa /root/.ssh/ims-go-components-rsa
ADD lumavate-components-rsa /root/.ssh/lumavate-components-rsa
ADD ims-lumavate-components-rsa /root/.ssh/ims-lumavate-components-rsa
RUN chmod 400 /root/.ssh/*-rsa
RUN go get github.com/astaxie/beego && \
  go get github.com/beego/bee && \
  go get github.com/Lumavate-Team/lumavate-go-common && \
  go get github.com/bitly/go-simplejson && \
  cd /go/src/github.com/Lumavate-Team && \
  sh /git.sh -i /root/.ssh/ims-go-components-rsa clone git@github.com:Lumavate-Team/ims-go-components.git && \
  cd ims-go-components && \
  mv static/views/ /go/src/widget/views/layout/ && \
  rm -rf .git && \
  cd / && \
  sh /git.sh -i /root/.ssh/lumavate-components-rsa clone git@github.com:Lumavate-Team/lumavate-components.git && \
  cd lumavate-components && \
  npm install && \
  npm run build && \
  cd / && \
  sh /git.sh -i /root/.ssh/ims-lumavate-components-rsa clone git@github.com:Lumavate-Team/ims-lumavate-components.git && \
  cd ims-lumavate-components && \
  npm install && \
  npm run build && \
  rm /root/.ssh/*

CMD bee run

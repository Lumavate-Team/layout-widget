FROM node:8.9-alpine as builder

RUN apk update && \
                apk add --no-cache \
                git \
                curl \
                libc-dev \
                go \
  && mkdir -p /go/src/widget

ENV GOPATH=/go
ENV PATH="/go/bin:${PATH}"

WORKDIR /go/src/widget
COPY ./widget /go/src/widget

RUN go get github.com/astaxie/beego && \
  go get github.com/beego/bee && \
  go get github.com/Lumavate-Team/lumavate-go-common && \
  go get github.com/bitly/go-simplejson && \
  cd / && \
  npm install @lumavate/components --save

CMD ["bee", "run"]

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
  go get github.com/bitly/go-simplejson && \
  mkdir /go/src/github.com/Lumavate-Team && \
  cd /go/src/github.com/Lumavate-Team && \
  git clone https://github.com/Lumavate-Team/lumavate-go-common.git && \
  cd lumavate-go-common && \
  git checkout performance && \
	mkdir /lumavate-core-components-install && \
	cd /lumavate-core-components-install && \
	npm install @lumavate/lumavate-core-components

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/src/widget /app/
COPY --from=builder /lumavate-core-components-install/node_modules/@lumavate/lumavate-core-components/dist /lumavate-core-components

WORKDIR /app

CMD ["./main"]

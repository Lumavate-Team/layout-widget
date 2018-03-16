from golang:alpine

RUN apk add --no-cache \
                git \
                curl \
                openssh \
                nodejs-current nodejs-npm \
#&& git rev-parse HEAD > /revision \
#&& rm -rf .git \
  && mkdir -p /go/src/widget

ADD git.sh /git.sh

WORKDIR /go/src/widget
COPY ./widget /go/src/widget

RUN mkdir /root/.ssh/ && \
  touch /root/.ssh/known_hosts && \
  ssh-keyscan github.com >> /root/.ssh/known_hosts

ADD ims-components-rsa /root/.ssh/ims-components-ims
ADD lumavate-components-rsa /root/.ssh/lumavate-components-ims
RUN chmod 400 /root/.ssh/*-ims
RUN go get -u github.com/astaxie/beego && \
  go get -u github.com/beego/bee && \
  go get -u github.com/Lumavate-Team/go-signer && \
  go get -u github.com/Lumavate-Team/go-properties && \
  cd /go/src/github.com/Lumavate-Team && \
  sh /git.sh -i /root/.ssh/ims-components-ims clone git@github.com:Lumavate-Team/ims-go-components.git && \
  cd ims-go-components && \
  rm -rf .git && \
	cd / && \
	sh /git.sh -i /root/.ssh/lumavate-components-ims clone git@github.com:Lumavate-Team/lumavate-components.git && \
	cd lumavate-components && \
	npm install && \
	npm run build && \
  rm /root/.ssh/* && \
	cp -r /lumavate-components/dist/* /go/src/widget/static/js

CMD bee run

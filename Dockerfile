from golang:alpine

RUN apk add --no-cache \
                git \
                curl \
                openssh \
#&& git rev-parse HEAD > /revision \
#&& rm -rf .git \
  && mkdir -p /go/src/widget

RUN mkdir /root/.ssh/
RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
ADD ims-components-rsa /root/.ssh/id_rsa

WORKDIR /go/src/widget
COPY ./widget /go/src/widget

RUN go get -u github.com/astaxie/beego
RUN go get -u github.com/beego/bee
RUN go get -u github.com/Lumavate-Team/go-signer
RUN go get -u github.com/Lumavate-Team/go-properties
RUN cd /go/src/github.com/Lumavate-Team && \
  git clone git@github.com:Lumavate-Team/ims-go-components.git && \
  cd ims-go-components && \
  rm -rf .git && \
  rm /root/.ssh/id_rsa

CMD bee run

FROM edit:base as editor

COPY layout_supervisord.conf /etc/supervisor/conf.d

FROM node:8.9-alpine

COPY --from=editor /editor /editor
COPY --from=editor /logs /logs
COPY --from=editor /etc/supervisor /etc/supervisor/
COPY --from=editor /edit_requirements.txt /edit_requirements.txt

RUN apk update && apk add --no-cache -t .build_deps \
    gcc \
    libc-dev \
    libgcc \
    linux-headers \
    libffi-dev \
    libressl-dev \
    musl-dev \
  && apk add --no-cache py3-greenlet \
  && apk add --no-cache python3 \
  && python3 -m ensurepip \
  && pip3 install --upgrade pip setuptools \
  && mkdir -p /editor \
  && pip3 install -r edit_requirements.txt \
  && apk del .build_deps

ENV EDITOR_SETTINGS=config/go_app.cfg

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
  git checkout master

CMD ["supervisord", "-c", "/etc/supervisor/supervisord.conf"]

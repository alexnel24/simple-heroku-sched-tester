FROM golang:1.16-alpine

ENV PORT=8081
ENV CRON_STR_FOR_SCHED="CRON_TZ=US/Mountain 24 12 * * 1-5"

COPY .build/certs/*.crt /usr/local/share/ca-certificates/

COPY . /src

RUN apk add --allow-untrusted --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/latest-stable/main/ ca-certificates \
  && update-ca-certificates \
  && apk add --no-cache bash curl \
  && apk --no-cache --available upgrade

WORKDIR /src

RUN go mod download

EXPOSE $PORT

CMD go run cmd/main.go
FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/../../pers/educonn-monorepo/services/account
COPY . /usr/local/go/src/../../pers/educonn-monorepo/services/account
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/account ./account


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/../../pers/educonn-monorepo/services/account/build/account /bin/account
CMD ["account", "up"]

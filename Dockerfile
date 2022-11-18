FROM golang:alpine AS build

RUN apk add git curl

RUN mkdir /prevail
ADD . /prevail
WORKDIR /prevail

RUN go build -o /tmp/http-server ./main.go

FROM alpine:edge

COPY --from=build /tmp/http-server /sbin/http-server

CMD /sbin/http-server

FROM golang:latest as goimage
ENV SRC=/go/src/
RUN mkdir -p /go/src/
WORKDIR /go/src/shrike
RUN git clone -b master --single-branch https://github.com/SteveCastle/shrike.git /go/src/shrike/
RUN go get github.com/labstack/echo && go get github.com/sirupsen/logrus && go get github.com/dgrijalva/jwt-go && go get github.com/gorilla/websocket
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -v -o bin/shrike
FROM alpine:latest
WORKDIR /docker/bin
COPY --from=goimage /go/src/shrike/bin/shrike /docker/bin/
ENTRYPOINT ["/docker/bin/shrike"]
EXPOSE 8080

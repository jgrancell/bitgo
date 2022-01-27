FROM docker.io/library/golang:1.17-alpine as build

WORKDIR /go/src/app

COPY . .
RUN go get -d -v ./... \
  && go build -o api \
  && chmod +x api

FROM docker.io/library/alpine:latest

COPY --from=build /go/src/app/api /usr/bin/go-api

USER nobody

EXPOSE 8080

CMD ["/usr/bin/go-api"]
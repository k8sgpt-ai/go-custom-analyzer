FROM golang:1.22-alpine3.20
RUN mkdir /src
ADD . /src/
WORKDIR /src
RUN go build -ldflags "-s -w -X main.version=$(cat VERSION)" -o analyzer
EXPOSE 8080
ENTRYPOINT ["/src/analyzer"]

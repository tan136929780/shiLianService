FROM golang:1.17-stretch as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN mkdir -p /go/src/vfile

WORKDIR /go/src/vfile
COPY . /go/src/vfile
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o main

FROM alpine
RUN mkdir -p /app
COPY    --from=builder /go/src/vfile/main  /app
COPY    --from=builder /go/src/vfile/config/config.yaml  /app/config/config.yaml
WORKDIR /app
RUN chmod 777 /app/main
EXPOSE 8200
CMD ["/app/main"]
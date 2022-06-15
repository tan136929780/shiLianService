FROM golang:1.17-stretch as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN mkdir -p /go/src/vfile

WORKDIR /go/src/vfile
COPY . /go/src/vfile
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o main .

FROM alpine

COPY    --from=builder /go/src/vfile/main  /
COPY    --from=builder /go/src/vfile/config/config.yaml  /config/config.yaml
WORKDIR /
RUN chmod 777 /main
EXPOSE 8200
CMD ["/main"]
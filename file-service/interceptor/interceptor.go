/*
@Time : 2022/6/7 下午3:36
@Author : tan
@File : AccessInterceptor
@Software: GoLand
*/
package interceptor

import (
	"context"
	"file-service/utils/logging"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

// 拦截器 - 请求记录
func AccessInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	//level, _ := logrus.ParseLevel("info")
	//logging.AccessLogger.Log(level, info.FullMethod, req)
	resp, err := handler(ctx, req)
	//logging.AccessLogger.Log(level, info.FullMethod, resp)
	return resp, err
}

type wrappedStream struct {
	grpc.ServerStream
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	fmt.Printf("Receive a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	fmt.Printf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {
	level, _ := logrus.ParseLevel("info")
	logging.AccessLogger.Log(level, info.FullMethod)
	err := handler(srv, newWrappedStream(ss))
	logging.AccessLogger.Log(level, info.FullMethod)
	return err
}

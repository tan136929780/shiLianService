package grpcClient

import (
	"context"
	"file-service/vfile/proto/vcms"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

func GetGrpcClient(host string, port int) *grpc.ClientConn {
	var serviceHost = host + ":" + strconv.Itoa(port)
	conn, err := grpc.Dial(serviceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}
	return conn
}

func AddFileInstence(host string, port int) (*vcms.CreateResult, error) {
	conn := GetGrpcClient(host, port)
	defer conn.Close()
	grpcClient := vcms.NewInstanceServiceClient(conn)
	rsp, err := grpcClient.CreateInstance(context.TODO(), &vcms.InstanceCreateRequest{
		Identifier: "",
		Data:       nil,
	})
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func GetFileInstence(host string, port int) (*vcms.InstanceFindByIdResponse, error) {
	conn := GetGrpcClient(host, port)
	defer conn.Close()
	grpcClient := vcms.NewInstanceServiceClient(conn)
	rsp, err := grpcClient.FindInstanceById(context.TODO(), &vcms.InstanceFindByIdRequest{
		Uid:         "",
		Identifiers: nil,
	})
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

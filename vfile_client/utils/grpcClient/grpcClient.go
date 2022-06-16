package grpcClient

import (
	"context"
	"file-service/vfile_client/proto/newvcms"
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

func AddFileInstence(host string, port int, request *newvcms.InstanceCreateRequest) (*newvcms.CreateResult, error) {
	conn := GetGrpcClient(host, port)
	defer conn.Close()
	grpcClient := newvcms.NewInstanceServiceClient(conn)
	rsp, err := grpcClient.CreateInstance(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func DelFileInstence(host string, port int, request *newvcms.InstanceDeleteRequest) (*newvcms.DeleteResult, error) {
	conn := GetGrpcClient(host, port)
	defer conn.Close()
	grpcClient := newvcms.NewInstanceServiceClient(conn)
	rsp, err := grpcClient.DeleteInstance(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func FindInstenceById(host string, port int) (*newvcms.InstanceFindByIdResponse, error) {
	conn := GetGrpcClient(host, port)
	defer conn.Close()
	grpcClient := newvcms.NewInstanceServiceClient(conn)
	rsp, err := grpcClient.FindInstanceById(context.TODO(), &newvcms.InstanceFindByIdRequest{
		Uid:         "",
		Identifiers: nil,
	})
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func FindInstance(host string, port int, request *newvcms.InstanceFindRequest) (*newvcms.InstanceFindResponse, error) {
	conn := GetGrpcClient(host, port)
	defer conn.Close()
	grpcClient := newvcms.NewInstanceServiceClient(conn)
	rsp, err := grpcClient.FindInstance(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

/*
@Time : 2022/6/1 上午11:05
@Author : tan
@File : filehandler
@Software: GoLand
*/
package handler

import (
	"context"
	"file-service/vfile/proto/vfile"
	"file-service/vfile/utils/constant"
	"file-service/vfile/utils/errorUtil"
	"strings"
)

type FileHandler struct {
	vfile.UnimplementedFileServiceServer
}

// 返回serviceInfo，该方法不需要复写
func (f FileHandler) ServiceInfo(ctx context.Context, request *vfile.ServiceInfoRequest) (*vfile.ServiceInfoResponse, error) {
	serviceInfoResponse := &vfile.ServiceInfoResponse{
		Version:        constant.Version,
		Os:             *constant.Os,
		Hostname:       *constant.HostName,
		AcceptProtocol: *constant.AcceptProtocol,
	}
	return serviceInfoResponse, nil
}

func (f FileHandler) Upload(ctx context.Context, request *vfile.UploadRequest) (*vfile.UploadResponse, error) {
	handler, err := converter(request.Metadata.FileName)
	if err != nil {
		return nil, err
	}
	res, err := handler.Upload(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f FileHandler) UploadStream(server vfile.FileService_UploadStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (f FileHandler) UploadWithOption(server vfile.FileService_UploadWithOptionServer) error {
	//TODO implement me
	panic("implement me")
}

func (f FileHandler) Download(ctx context.Context, request *vfile.DownloadRequest) (*vfile.DownloadResponse, error) {
	handler, err := converter(request.Uri)
	if err != nil {
		return nil, err
	}
	res, err := handler.Download(request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f FileHandler) DownloadStream(request *vfile.DownloadRequest, server vfile.FileService_DownloadStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (f FileHandler) DownloadWithOption(request *vfile.DownloadRequest, server vfile.FileService_DownloadWithOptionServer) error {
	//TODO implement me
	panic("implement me")
}

// 转化器
func converter(protocol string) (handler BaseHandler, err error) {
	support := checkIfSupportProtocol(protocol)
	if !support {
		checkIfSupportProtocol(protocol)
	}
	switch {
	case strings.HasPrefix(protocol, "vfile://"):
		handler = VFile{}
	default:
		handler = VFile{}
	}
	if handler == nil {
		return nil, errorUtil.NewStringError("File Handler Not Found!")
	}
	return handler, nil
}

// 协议支持校验
func checkIfSupportProtocol(protocol string) bool {
	acceptProtocol := strings.Split(*constant.AcceptProtocol, "|")
	for _, p := range acceptProtocol {
		if strings.HasPrefix(protocol, p+"://") {
			return true
		}
	}
	return false
}

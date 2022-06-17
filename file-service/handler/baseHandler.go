package handler

import (
	"file-service/proto/vfile"
)

type BaseHandler interface {
	Upload(request *vfile.UploadRequest) (*vfile.UploadResponse, error)
	UploadStream(metaData *vfile.MetaData, file *vfile.File) error
	UploadWithOption(server vfile.FileService_UploadWithOptionServer) error
	Download(request *vfile.DownloadRequest) (*vfile.DownloadResponse, error)
	DownloadStream(request *vfile.DownloadRequest, server vfile.FileService_DownloadStreamServer) error
	DownloadWithOption(request *vfile.DownloadRequest, server vfile.FileService_DownloadWithOptionServer) error
}

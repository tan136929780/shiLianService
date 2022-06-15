package handler

import (
	"file-service/vfile/proto/vfile"
	"file-service/vfile/utils/fileUtil"
	"path"
)

type VFile struct {
}

func (f VFile) Upload(uploadRequest *vfile.UploadRequest) (*vfile.UploadResponse, error) {
	fileName, err := fileUtil.FileWrite(uploadRequest.File.Content, uploadRequest.Metadata.Hash+uploadRequest.Metadata.Type)
	code := 0
	message := "上传成功"
	uri := path.Base(fileName)
	if err != nil {
		code = 1
		message = err.Error()
		uri = ""
	}
	res := &vfile.UploadResponse{
		Code:    int32(code),
		Message: message,
		Uri:     uri,
	}
	return res, nil
}

func (f VFile) UploadStream(metaData *vfile.MetaData, file *vfile.File) error {
	//TODO implement me
	panic("implement me")
}

func (f VFile) UploadWithOption(server vfile.FileService_UploadWithOptionServer) error {
	//TODO implement me
	panic("implement me")
}

func (f VFile) Download(DownloadRequest *vfile.DownloadRequest) (*vfile.DownloadResponse, error) {
	fileBytes, err := fileUtil.FileRead(DownloadRequest.Uri)
	code := 0
	message := "下载成功"
	metaData := &vfile.MetaData{
		FileName: "test.png",
		Uri:      "bf91a92aa29f34911540557b41ce79ac.png",
		Type:     ".png",
		Hash:     "bf91a92aa29f34911540557b41ce79ac",
		FileSize: int64(len(fileBytes)),
	}
	file := &vfile.File{Content: fileBytes}
	if err != nil {
		code = 1
		message = err.Error()
		metaData = nil
		file = nil
	}
	res := &vfile.DownloadResponse{
		Code:     int32(code),
		Message:  message,
		Metadata: metaData,
		File:     file,
	}
	return res, nil
}

func (f VFile) DownloadStream(request *vfile.DownloadRequest, server vfile.FileService_DownloadStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (f VFile) DownloadWithOption(request *vfile.DownloadRequest, server vfile.FileService_DownloadWithOptionServer) error {
	//TODO implement me
	panic("implement me")
}

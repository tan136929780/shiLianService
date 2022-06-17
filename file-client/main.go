package main

/*
 struct ServiceInfo {
	int  version;
	char *os;
	char *hostname;
	char *acceptProtocol;
};
struct HostInfo {
	char *host;
	int port;
};
struct ServiceInfoRequest {
	struct HostInfo *hostInfo;
};
struct MetaData {
     char *fileName;
     char *uri;
     char *fileType;
     char *hash;
     long long int fileSize;
};
struct File {
	void* content;
};
struct UploadRequest {
	struct MetaData *metaData;
	struct File *file;
};
struct UploadInfoRequest {
	struct HostInfo *hostInfo;
	struct UploadRequest *uploadRequest;
};
struct UploadResponse {
    long int code;
    char *message;
    char *uri;
};
struct DownloadOption {
     long long int sliceSize;
     long long int timeout;
     long long int rangeStart;
     long long int rangeEnd;
};
struct DownloadRequest {
	char *uri;
	struct DownloadOption *downloadOption;
};
struct DownloadInfoRequest {
	struct HostInfo *hostInfo;
	struct DownloadRequest *downloadRequest;
};
struct DownloadResponse {
	long int code;
    char *message;
    struct MetaData *metadata;
    struct File *file;
};
*/
import "C"
import (
	"context"
	"encoding/json"
	"file-client/proto/newvcms"
	"file-client/proto/vfile"
	"file-client/utils/config"
	"file-client/utils/fileUtil"
	"file-client/utils/grpcClient"
	"fmt"
)

//export ServiceInfo
func ServiceInfo(in C.struct_ServiceInfoRequest) C.struct_ServiceInfo {
	host := C.GoString(in.hostInfo.host)
	port := int(in.hostInfo.port)
	conn := grpcClient.GetGrpcClient(host, port)
	if conn == nil {
		fmt.Println("client error")
	}
	defer conn.Close()
	rsp, err := vfile.NewFileServiceClient(conn).ServiceInfo(context.TODO(), &vfile.ServiceInfoRequest{})
	if err != nil {
		fmt.Println(err.Error())
	}
	response := C.struct_ServiceInfo{}
	response.version = C.int(rsp.Version)
	response.os = C.CString(rsp.Os)
	response.hostname = C.CString(rsp.Hostname)
	response.acceptProtocol = C.CString(rsp.AcceptProtocol)
	return response
}

//export Upload
func Upload(in C.struct_UploadInfoRequest) C.struct_UploadResponse {
	host := C.GoString(in.hostInfo.host)
	port := int(in.hostInfo.port)
	conn := grpcClient.GetGrpcClient(host, port)
	if conn == nil {
		panic("grpc client error")
	}
	defer conn.Close()
	uploadRequest := &vfile.UploadRequest{
		Metadata: &vfile.MetaData{
			FileName: C.GoString(in.uploadRequest.metaData.fileName),
			Uri:      C.GoString(in.uploadRequest.metaData.uri),
			Type:     C.GoString(in.uploadRequest.metaData.fileType),
			Hash:     C.GoString(in.uploadRequest.metaData.hash),
			FileSize: int64(in.uploadRequest.metaData.fileSize),
		},
		//TODO
		File: &vfile.File{Content: nil},
	}
	fmt.Printf("content: %#v\n", uploadRequest)
	rsp, err := vfile.NewFileServiceClient(conn).Upload(context.TODO(), uploadRequest)
	response := C.struct_UploadResponse{}
	if err != nil {
		response.code = C.long(rsp.Code)
		response.message = C.CString(err.Error())
		return response
	}
	response.code = C.long(rsp.Code)
	response.message = C.CString(rsp.Message)
	response.uri = C.CString(rsp.Uri)
	return response
}

//export Download
func Download(in C.struct_DownloadInfoRequest) C.struct_DownloadResponse {
	host := C.GoString(in.hostInfo.host)
	port := int(in.hostInfo.port)
	conn := grpcClient.GetGrpcClient(host, port)
	if conn == nil {
		panic("grpc client error")
	}
	defer conn.Close()
	downloadRequest := &vfile.DownloadRequest{
		Uri: C.GoString(in.downloadRequest.uri),
		Option: &vfile.DownloadOption{
			SliceSize:  int64(in.downloadRequest.downloadOption.sliceSize),
			Timeout:    int64(in.downloadRequest.downloadOption.timeout),
			RangeStart: int64(in.downloadRequest.downloadOption.rangeStart),
			RangeEnd:   int64(in.downloadRequest.downloadOption.rangeEnd),
		},
	}
	rsp, err := vfile.NewFileServiceClient(conn).Download(context.TODO(), downloadRequest)
	response := C.struct_DownloadResponse{}
	if err != nil {
		response.code = C.long(1)
		response.message = C.CString(err.Error())
		return response
	}
	response.code = C.long(rsp.Code)
	response.message = C.CString(rsp.Message)
	metaData := C.struct_MetaData{}
	metaData.fileName = C.CString(rsp.Metadata.FileName)
	metaData.uri = C.CString(rsp.Metadata.Uri)
	metaData.fileType = C.CString(rsp.Metadata.Type)
	metaData.hash = C.CString(rsp.Metadata.Hash)
	metaData.fileSize = C.longlong(rsp.Metadata.FileSize)
	response.metadata = &metaData
	file := C.struct_File{}
	file.content = C.CBytes(rsp.File.Content)
	response.file = &file
	return response
}

func main() {
	//upload()
	download()
	//instanceCreate()
	//deleteInstance()
	//createFileMetaData()
	//FindInstance()
}

func FindInstance() {
	//0x78ec
	//identifier := "bf91a92aa29f34911540557b41ce79ac.png"
	request := &newvcms.InstanceFindRequest{
		Identifier: "FileMetaData",
		Sorters:    nil,
		Page:       1,
		PageSize:   10,
		Condition: map[string]string{
			"instance.identifier": "bf91a92aa29f34911540557b41ce79ac.png",
			"FileMetaData.status": "1",
		},
	}
	response, err := grpcClient.FindInstance("10.1.68.60", 9090, request)
	if err != nil {
		fmt.Printf("Printf err Content: %#v\n", err.Error())
	}
	data := &fileUtil.FileInstanceList{}
	err = json.Unmarshal([]byte(response.Data), data)
	if err != nil {
		return
	}
	fmt.Printf("Printf s Content: %v\n", data.Instances[0].FileMetaDataUri)
}

func createFileMetaData() {
	isa, _ := json.Marshal([]string{"FileMetaData"})
	request := &newvcms.InstanceCreateRequest{
		Identifier: "FileMetaData",
		Data: map[string]string{
			"instance.identifier":   "bf91a92aa29f34911540557b41ce79ac.png",
			"instance.isa":          string(isa),
			"FileMetaData.fileName": "test.png",
			"FileMetaData.uri":      "bf91a92aa29f34911540557b41ce79ac.png",
			"FileMetaData.type":     ".png",
			"FileMetaData.hash":     "bf91a92aa29f34911540557b41ce79ac",
			"FileMetaData.fileSize": "78899",
		},
	}
	response, err := grpcClient.AddFileInstence("10.1.68.60", 9090, request)
	if err != nil {
		fmt.Printf("Printf err Content: %#v\n", err.Error())
	}
	fmt.Printf("Printf response Content: %#v\n", response)
}

func instanceCreate() {
	isaProperty, _ := json.Marshal([]string{"property"})
	propreties, _ := json.Marshal([]map[string]string{
		{
			"instance.identifier": "文件名",
			"instance.isa":        string(isaProperty),
			"property.identifier": "fileName",
			"property.type":       "string",
			"property.status":     "1",
		},
		{
			"instance.identifier": "文件唯一标识",
			"instance.isa":        string(isaProperty),
			"property.identifier": "uri",
			"property.type":       "string",
			"property.status":     "1",
		},
		{
			"instance.identifier": "文件类型",
			"instance.isa":        string(isaProperty),
			"property.identifier": "type",
			"property.type":       "string",
			"property.status":     "1",
		},
		{
			"instance.identifier": "hash",
			"instance.isa":        string(isaProperty),
			"property.identifier": "hash",
			"property.type":       "string",
			"property.status":     "1",
		},
		{
			"instance.identifier": "文件大小",
			"instance.isa":        string(isaProperty),
			"property.identifier": "fileSize",
			"property.type":       "int",
			"property.status":     "1",
		},
	})
	isa, _ := json.Marshal([]string{"type"})
	FileMetaData := map[string]string{
		"instance.identifier": "文件上传模型",
		"instance.isa":        string(isa),
		"type.identifier":     "FileMetaData",
		"type.property":       string(propreties),
		"type.status":         "1",
	}
	request := &newvcms.InstanceCreateRequest{
		Identifier: "type",
		Data:       FileMetaData,
	}
	response, err := grpcClient.AddFileInstence("10.1.68.60", 9090, request)
	if err != nil {
		fmt.Printf("Printf err Content: %#v\n", err.Error())
	}
	fmt.Printf("Printf response Content: %#v\n", response)
}

func deleteInstance() {
	//uids := []string{"0x78c9", "0x78ce", "0x78cd", "0x78cc", "0x78cb", "0x78ca"}
	uids := []string{"0x78d1", "0x78d2", "0x78d3", "0x78d4", "0x78d5", "0x78d6"}
	request := &newvcms.InstanceDeleteRequest{
		Identifier: "type",
		Ids:        uids,
	}
	response, err := grpcClient.DelFileInstence("10.1.68.60", 9090, request)
	if err != nil {
		fmt.Printf("Printf err Content: %#v\n", err.Error())
	}
	fmt.Printf("Printf response Content: %#v\n", response)
}

func upload() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
	conn := grpcClient.GetGrpcClient("127.0.0.1", 8200)
	if conn == nil {
		fmt.Println("client error")
	}
	defer conn.Close()
	fileName := "test.png"
	fileBytes, err := fileUtil.FileRead(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	metaData, err := fileUtil.ParseFile(fileName, "vfile")
	if err != nil {
		fmt.Println(err.Error())
	}
	metaData.FileSize = int64(len(fileBytes))
	uploadRequest := &vfile.UploadRequest{
		Metadata: &metaData,
		File:     &vfile.File{Content: fileBytes},
	}
	rsp, err := vfile.NewFileServiceClient(conn).Upload(context.TODO(), uploadRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Printf rsp Content: %#v\n", rsp)
}

func download() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
	conn := grpcClient.GetGrpcClient("127.0.0.1", 8200)
	if conn == nil {
		fmt.Println("client error")
	}
	defer conn.Close()
	downloadRequest := &vfile.DownloadRequest{
		Uri:    "vfile://bf91a92aa29f34911540557b41ce79ac.png",
		Option: nil,
	}
	rsp, err := vfile.NewFileServiceClient(conn).Download(context.TODO(), downloadRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
	fileName, err := fileUtil.FileWrite(rsp.File.Content, rsp.Metadata.Uri)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Printf fileName Content: %#v\n", fileName)
}

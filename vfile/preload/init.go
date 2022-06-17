package main

import (
	"encoding/json"
	"file-service/vfile/proto/newvcms"
	"file-service/vfile/utils/config"
	"file-service/vfile/utils/grpcClient"
	"fmt"
)

func main() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
	instanceCreate()
}

// 创建文件模型
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
	response, err := grpcClient.AddFileInstence(config.GetString("grpc.instanceService.host"), config.GetInt("grpc.instanceService.host"), request)
	if err != nil {
		fmt.Printf("Printf err Content: %#v\n", err.Error())
	}
	fmt.Printf("Printf response Content: %#v\n", response)
}

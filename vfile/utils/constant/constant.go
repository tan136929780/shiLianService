/*
@Time : 2022/6/7 下午5:07
@Author : tan
@File : const
@Software: GoLand
*/
package constant

import "file-service/vfile/proto/vfile"

const (
	SUCCESS int32 = iota
	FAILED
)

var Version = vfile.FileServiceVersion_V_0_1_0
var Os *string
var HostName *string
var AcceptProtocol *string

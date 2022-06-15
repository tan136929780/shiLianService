/*
@Time : 2022/6/14 下午5:55
@Author : tan
@File : fileRead
@Software: GoLand
*/
package fileUtil

import (
	"crypto/md5"
	"encoding/hex"
	"file-service/vfile/proto/vfile"
	"file-service/vfile/utils/config"
	"file-service/vfile/utils/errorUtil"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func FileRead(fileName string) ([]byte, error) {
	pathName := GetStorePath()
	fileName = pathName + strings.Trim(strings.TrimPrefix(fileName, "vfile://"), "/")
	fileExist := CheckFileExist(fileName)
	if !fileExist {
		return nil, errorUtil.NewStringError("下载文件不存在")
	}
	fp, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 6) // 读写方式打开
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	bytes, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}
	return bytes, err
}

func FileWrite(fileBytes []byte, fileName string) (string, error) {
	pathName := GetStorePath()
	_, error := os.Stat(pathName)
	if os.IsNotExist(error) {
		os.MkdirAll(pathName, 0755)
	}
	isExist := CheckFileExist(fileName)
	if isExist {
		return fileName, nil
	}
	fileName = pathName + fileName
	err := ioutil.WriteFile(fileName, fileBytes, 0755)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GetFileMd5(fileName string) (md5Str string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)
	md5Str = hex.EncodeToString(hashInBytes)
	return
}

func ParseFile(fileName, proto string) (vfile.MetaData, error) {
	fileNameWithSuffix := path.Base(fileName)
	fileType := path.Ext(fileNameWithSuffix)
	fileNameOnly := strings.TrimSuffix(fileNameWithSuffix, fileType)
	md5Str, err := GetFileMd5(fileName)
	if err != nil {
		return vfile.MetaData{}, err
	}
	return vfile.MetaData{
		FileName: proto + "://" + fileNameOnly,
		Uri:      "",
		Type:     fileType,
		Hash:     md5Str,
		FileSize: 0,
	}, nil
}

func GetStorePath() string {
	return strings.TrimSuffix(config.GetString("upload.fileWriteDir"), "/") + "/"
}

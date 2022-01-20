package utils

import (
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

//SaveFile2Dst 将文件保存到目标地址(拷贝文件)
//
//file *multipart.FileHeader 文件
//
//dst string 拷贝到的目标地址
func SaveFile2Dst(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func() {
		_ = src.Close()
	}()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	_, err = io.Copy(out, src)

	return err
}

//FileGetContents 获取文件内容
//
//filename string 文件地址
func FileGetContents(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

//FilePutContents 将内容写入文件(覆盖写)
//
//content []byte 写入的内容
//
//dst string 写入的目标地址
func FilePutContents(content []byte, dst string) error {
	return ioutil.WriteFile(dst, content, 0644)
}

//FileExists 检查文件上是否存在
//
//dst string 目标地址
func FileExists(dst string) bool {
	ok, _ := FileExistsWithError(dst)

	return ok
}

//FileExistsWithError 检查文件上是否存在(会返回错误信息)
//
//dst string 目标地址
func FileExistsWithError(dst string) (bool, error) {
	_, err := os.Stat(dst)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, err
}

//FileExt 获取文件扩展名
//
//根据文件名最后一个.分隔来切分获取
func FileExt(filename string) string {
	filenameSplit := strings.Split(filename, ".")
	return filenameSplit[len(filenameSplit)-1]
}

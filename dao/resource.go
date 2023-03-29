package dao

import (
	"LipLanguage/common"
	"LipLanguage/model"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// CreateResourceData 创建资源数据
func CreateResourceData(data []byte) (res model.Resource, err error) {
	// 写入文件
	filename := fmt.Sprintf("%x", md5.Sum(data))
	res.SrcID = filename
	path := common.SrcPath + filename
	res.Path = path
	f, err := os.Create(path)
	if err != nil {
		return
	}
	_, err = f.Write(data)
	if err != nil {
		return
	}
	err = f.Close()
	if err != nil {
		return
	}

	// 写入数据库
	err = DB.Create(&res).Error
	return
}

// GetResourceDataBySrcID 通过SrcID获取资源数据
func GetResourceDataBySrcID(SrcID string) ([]byte, error) {
	// 从数据库中获取资源
	ret := &model.Resource{}
	err := DB.Model(model.Resource{}).Where("src_id=?", SrcID).Take(ret).Error
	// 读取文件
	if err != nil {
		return nil, err
	}
	f, err := os.Open(ret.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

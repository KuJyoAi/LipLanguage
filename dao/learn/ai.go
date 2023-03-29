package learn

import (
	"LipLanguage/common"
	"LipLanguage/model"
	"bytes"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"net/http"
	"sync"
	"time"
)

var AiLock sync.Mutex

// PostToAi 把视频文件post过去, 发送路径
func PostToAi(data []byte) (ret model.AiPostResponse, err error) {
	// 加锁防止AI被高并发请求
	AiLock.Lock()
	defer AiLock.Unlock()

	// 请求部分:
	URL := common.AIUrl
	file := bytes.NewReader(data)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("video", "file.webm")
	if err != nil {
		logrus.Errorf("[util.PostVideoPath] CreateFormFile %v", err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		logrus.Errorf("[util.PostVideoPath] Copy %v", err)
		return
	}
	err = writer.Close()
	if err != nil {
		logrus.Errorf("[util.PostVideoPath] writer.Close%v", err)
		return
	}

	// 发送请求
	logrus.Infof("[util.PostVideoPath] send request to %v", URL)
	request, err := http.NewRequest("POST", URL, body)
	if err != nil {
		logrus.Errorf("[util.PostVideoPath] NewRequest%v", err)
		return
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		logrus.Errorf("[util.PostVideoPath] %v", err)
		return
	}
	defer resp.Body.Close()

	// 读取返回
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("[util.PostVideoPath] %v", err)
		return
	}

	if resp.StatusCode != 200 {
		//有错误
		logrus.Errorf("[util.PostVideoPath] %v", resp.Status)
		return ret, errors.New(fmt.Sprintf("AI返回错误:%v", resp.Status))
	}

	// 数据格式: 00 00 结果 视频数据
	data[0] = 0    //去掉第一个字节
	data[1] = 0x0A //去掉第二个字节
	DataStart := 1 //数据开始位置
	//webm: 1A 45 DF A3
	//查找数据开始位置
	for i := 0; i < len(data)-4; i++ {
		if data[i] == 0x1A && data[i+1] == 0x45 && data[i+2] == 0xDF && data[i+3] == 0xA3 {
			DataStart = i
			break
		}
	}

	// 获取数据
	Res := data[2:DataStart]
	data = data[DataStart:]
	ret = model.AiPostResponse{
		Result: string(Res),
		Data:   data,
	}

	logrus.Infof(`[util.PostVideoPath] AI Response: data[1]: %d ResLen:%v result_len:%v`, data[1], len(Res), len(ret.Result))
	return ret, err
}

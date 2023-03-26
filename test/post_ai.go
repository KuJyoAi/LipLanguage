package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	path := "C:\\Users\\KuJyo\\Desktop\\北京时间.mp4"

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("video", path)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	// 读取返回
	req, err := http.NewRequest("POST", "http://103.222.190.10:25555/lip", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	logrus.Infof("response Status: %v", resp.Status)
	logrus.Infof("response Headers: %v", resp.Header)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	logrus.Infof("response Body: %v", len(data))
}

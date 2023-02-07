package test

import (
	"backend/core/define"
	"bytes"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func Test_cos(t *testing.T) {

	u, _ := url.Parse("https://cloud-1307889700.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	key := "cloud-disk/test.jpg"
	_, _, err := c.Object.Upload(context.Background(), key, "./img/test2.jpg", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func TestFileUploadByPut(t *testing.T) {
	u, _ := url.Parse("https://cloud-1307889700.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	key := "cloud-disk/test2.jpg"

	file, err := os.ReadFile("./img/test2.jpg")

	if err != nil {
		log.Fatal(err)
	}

	_, err = c.Object.Put(context.Background(), key, bytes.NewReader(file), nil)
	if err != nil {
		log.Fatal(err)
	}
}

// 分片上传初始化
func TestInitPartUpload(t *testing.T) {
	u, _ := url.Parse("https://cloud-1307889700.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/example2.jpg"
	v, _, err := c.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Fatal(err)
	}
	UploadID := v.UploadID
	fmt.Println(UploadID)
	// 1675755492a246050a07f7ca17d909cfb9b62ee61d6fef4afc4ff205cb68dbc39c6b92551d
}

// 分片上传

func TestPartUpload(t *testing.T) {

	u, _ := url.Parse("https://cloud-1307889700.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/example2.jpg"

	UploadID := "1675755492a246050a07f7ca17d909cfb9b62ee61d6fef4afc4ff205cb68dbc39c6b92551d"

	//f, err := os.ReadFile("0.chunk") //ca4d4d7856d8f86fc75d7d9740c9194a
	//f, err := os.ReadFile("1.chunk") // 87bcb2aea4cd7ff7e8c91f8cfba84288
	f, err := os.ReadFile("2.chunk") // de83d254c96ffb11bef3b70f61f4099e
	if err != nil {
		t.Fatal(err)
	}
	resp, err := c.Object.UploadPart(
		context.Background(), key, UploadID, 3, bytes.NewReader(f), nil,
	)
	if err != nil {
		panic(err)
	}
	PartETag := resp.Header.Get("ETag") // md5(0.chunk)
	fmt.Println(PartETag)
}

// 分片上传完成
func TestPartUploadComplete(t *testing.T) {
	u, _ := url.Parse("https://cloud-1307889700.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/example2.jpg"
	UploadID := "1675755492a246050a07f7ca17d909cfb9b62ee61d6fef4afc4ff205cb68dbc39c6b92551d"
	//md5 := "ca4d4d7856d8f86fc75d7d9740c9194a"

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: "ca4d4d7856d8f86fc75d7d9740c9194a"}, cos.Object{
		PartNumber: 2, ETag: "87bcb2aea4cd7ff7e8c91f8cfba84288"}, cos.Object{
		PartNumber: 3, ETag: "de83d254c96ffb11bef3b70f61f4099e"},
	)
	_, _, err := c.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}
}

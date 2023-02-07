package test

import (
	"backend/core/define"
	"bytes"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func Test_cos(t *testing.T) {

	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.tencentcloud.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.tencentcloud.com/cos5/bucket, 关于地域的详情见 https://www.tencentcloud.com/document/product/436/6224?from_cn_redirect=1
	u, _ := url.Parse("https://cloud-1307889700.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.tencentcloud.com/cam/capi
			SecretKey: define.TencentSecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.tencentcloud.com/cam/capi
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	key := "cloud-disk/test.jpg"
	_, _, err := c.Object.Upload(context.Background(), key, "./img/test2.jpg", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func TestFileUploadByPut(t *testing.T) {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.tencentcloud.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.tencentcloud.com/cos5/bucket, 关于地域的详情见 https://www.tencentcloud.com/document/product/436/6224?from_cn_redirect=1
	u, _ := url.Parse("https://cloud-1307889700.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.tencentcloud.com/cam/capi
			SecretKey: define.TencentSecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.tencentcloud.com/cam/capi
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
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

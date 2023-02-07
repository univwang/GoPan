package helper

import (
	"backend/core/define"
	"bytes"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// CosPartUploadComplete 分片上传完成
func CosPartUploadComplete(key, uploadId string, cs []cos.Object) error {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cs...)
	_, _, err := c.Object.CompleteMultipartUpload(
		context.Background(), key, uploadId, opt,
	)
	return err
}

// CosPartUpload 分片上传
func CosPartUpload(r *http.Request) (string, error) {

	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := r.PostForm.Get("key")
	UploadId := r.PostForm.Get("upload_id")
	part_number, err := strconv.Atoi(r.PostForm.Get("part_number"))
	if err != nil {
		return "", err
	}

	f, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)

	resp, err := c.Object.UploadPart(
		context.Background(), key, UploadId, part_number, bytes.NewReader(buf.Bytes()), nil,
	)
	if err != nil {
		return "", err
	}

	return strings.Trim(resp.Header.Get("ETag"), "\""), nil
}

// CosInitPart 分片上传初始化
func CosInitPart(ext string) (string, string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/" + UUID() + ext
	v, _, err := c.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return "", "", err
	}

	return key, v.UploadID, nil
}

func GenerateToken(id int, identity string, name string) (string, error) {
	// id
	// identity
	// name
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyzeToken
// Token 解析
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}

	return uc, nil
}

// CosUpload 文件上传
func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	file, header, err := r.FormFile("file")
	key := "cloud-disk/" + UUID() + path.Ext(header.Filename)

	_, err = c.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		panic(err)
	}
	return define.CosBucket + "/" + key, nil
}
func UUID() string {
	return uuid.NewV4().String()
}

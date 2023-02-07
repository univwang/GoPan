package helper

import (
	"backend/core/define"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"path"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
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

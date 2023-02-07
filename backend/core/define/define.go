package define

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var CosBucket = "https://cloud-1307889700.cos.ap-nanjing.myqcloud.com"
var TencentSecretID = "AKIDAD3f7NODm6B2kIXZQ4UecLr8B4lFUWL2"
var TencentSecretKey = "VlodDZieLo335sQdeS2PXQ8ctNdfT5Uh"
var JwtKey = "cloud-disk-key"

// 分页默认参数
var PageSize = 10

var Datetime = "2006:01:02 15:04:05"

package test

import (
	"backend/core/helper"
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {

	md5 := helper.Md5("123456")
	fmt.Println(md5)
}

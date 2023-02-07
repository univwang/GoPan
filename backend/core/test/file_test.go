package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"testing"
)

// 分片大小
const chunkSize = 1 * 1024 * 1024 // 1MB
// 文件分片

func TestGenerateChunkFile(t *testing.T) {
	//fileInfo, err := os.Stat("test.rar")
	fileInfo, err := os.Stat("/img/test2.jpg")
	if err != nil {
		t.Fatal(err)
	}
	// 分片个数
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	file, err := os.OpenFile("test.rar", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunkSize)

	for i := 0; i < int(chunkNum); i++ {
		// 指定读取文件的起始位置
		file.Seek(int64(i*chunkSize), 0)

		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		file.Read(b)

		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, 0777)

		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
	file.Close()
}

// 分片合并
func TestMergeChunkFile(t *testing.T) {

	file, err := os.OpenFile("test2.rar", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	fileInfo, err := os.Stat("test.rar")
	if err != nil {
		t.Fatal(err)
	}
	// 分片个数
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))

	for i := 0; i < int(chunkNum); i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_RDONLY, 0777)
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		file.Write(b)
		f.Close()
	}
	file.Close()
}

// 一致性校验

func TestCheckFile(t *testing.T) {
	file1, err := os.OpenFile("test.rar", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := ioutil.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}

	file2, err := os.OpenFile("test2.rar", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := ioutil.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}

	bytes1 := md5.Sum(b1)
	bytes2 := md5.Sum(b2)

	fmt.Println(bytes1, bytes2, bytes1 == bytes2)
}

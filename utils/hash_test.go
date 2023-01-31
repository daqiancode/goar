package utils

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestDeepHash(t *testing.T) {
	dataList := []interface{}{}
	dataList = append(dataList, Base64Encode([]byte("123")))
	assert.Equal(t, [48]uint8([48]uint8{0xfd, 0x33, 0x8d, 0x7c, 0x23, 0x3, 0xe5, 0xed, 0xb3, 0xce, 0x4b, 0x64, 0xc2, 0xbc, 0xd8, 0xa9, 0x4b, 0xd0, 0xee, 0x58, 0x79, 0x33, 0x6a, 0x3c, 0x11, 0xc5, 0x1, 0xa5, 0x6d, 0x79, 0xc7, 0xc7, 0x46, 0x1b, 0xca, 0x4a, 0x2a, 0xe3, 0xdc, 0xa9, 0xa, 0x8b, 0x66, 0xc8, 0xd, 0x5e, 0x53, 0x61}), DeepHash(dataList))

	dataList = append(dataList, Base64Encode([]byte("123")))
	assert.Equal(t, [48]uint8{0xe7, 0xb, 0xed, 0x22, 0x75, 0xe0, 0x4, 0xd, 0xf5, 0x3c, 0x97, 0xe9, 0x3f, 0x97, 0xa9, 0x5f, 0xa, 0x4d, 0x29, 0xc8, 0x5d, 0x76, 0x5f, 0x54, 0x9a, 0x32, 0x74, 0xb6, 0x6b, 0xf, 0x55, 0xa4, 0xb9, 0x2f, 0x23, 0x73, 0x34, 0xda, 0x1f, 0x25, 0xc1, 0xac, 0x59, 0x14, 0xfb, 0xe0, 0xfb, 0x89}, DeepHash(dataList))

	dataList = append(dataList, [][]string{
		[]string{
			Base64Encode([]byte("APP")),
			Base64Encode([]byte("1.0")),
		},
		[]string{
			Base64Encode([]byte("Contract")),
			Base64Encode([]byte("0x000")),
		},
	})
	assert.Equal(t, [48]uint8([48]uint8{0xfb, 0x9b, 0xee, 0xfc, 0x83, 0x50, 0xa, 0xe9, 0xc9, 0x41, 0x9f, 0xb1, 0x58, 0x9b, 0xe3, 0x9c, 0x55, 0x73, 0xfd, 0xad, 0xed, 0xbd, 0x85, 0xc9, 0x87, 0x42, 0x9, 0xfb, 0x4b, 0xb, 0x45, 0x83, 0xb8, 0x5c, 0xfa, 0xa, 0xf2, 0x21, 0x7c, 0xab, 0x65, 0x71, 0x5b, 0xf0, 0x88, 0xc6, 0xc2, 0x27}), DeepHash(dataList))
}

func TestDeepHashStr(t *testing.T) {
	hash := deepHashStr(Base64Encode([]byte("123")))
	assert.Equal(t, [48]uint8([48]byte{0x27, 0xbb, 0xb6, 0x32, 0x9a, 0xcb, 0x39, 0x7e, 0x10, 0x59, 0x85, 0xd2, 0x66, 0xa0, 0xe2, 0x1b, 0x4e, 0x64, 0xf1, 0xe0, 0x97, 0x8d, 0x2a, 0x31, 0x16, 0x88, 0xca, 0x99, 0xa7, 0xf5, 0xa4, 0x94, 0xd5, 0x8f, 0xf9, 0xf6, 0xc4, 0xc8, 0x66, 0x33, 0x36, 0x9, 0x83, 0x32, 0xd6, 0x88, 0x59, 0xc3}), hash)
}

func TestHashStream(t *testing.T) {
	file, err := os.Open("bytes.go")
	assert.NoError(t, err)
	var data interface{}
	data = file
	_, ok := data.(io.Reader)
	assert.Equal(t, ok, true)
	dataBy, err := ioutil.ReadAll(file)
	assert.NoError(t, err)
	file.Close()
	file2, err := os.Open("bytes.go")
	defer file2.Close()
	hashStr := deepHashStr(Base64Encode(dataBy))
	hashStream := deepHashStream(file2)
	assert.Equal(t, hashStr, hashStream)
}

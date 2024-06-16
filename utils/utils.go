package utils

import (
	"encoding/base64"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
	"unsafe"
)

var (
	HexArray = []byte("0123456789abcdef")
)

func Address(ctx *app.RequestContext) string {
	// Check headers.
	address := ctx.GetHeader("CF-Connecting-IP")
	if len(address) != 0 {
		return string(address)
	}
	address = ctx.GetHeader("X-Forwarded-For")
	if len(address) != 0 {
		return string(address)
	}
	address = ctx.GetHeader("X-Real-IP")
	if len(address) != 0 {
		return string(address)
	}
	// Return the request IP.
	return ctx.ClientIP()
}

func BytesToHex(bytes []byte) string {
	if bytes == nil {
		return ""
	}
	hexChars := make([]byte, len(bytes)*2)
	for i := range bytes {
		val := bytes[i] & 0xFF
		hexChars[i*2] = HexArray[val>>4]
		hexChars[i*2+1] = HexArray[val&0x0F]
	}
	return string(hexChars)
}

// Base64Encode use Encoder.RFC4648
func Base64Encode(toEncode []byte) string {
	return base64.StdEncoding.EncodeToString(toEncode)
}

func Base64Decode(toDecode string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(toDecode)
}

func AbilityHash(str string) int {
	bytes := []rune(str)
	v7 := 0
	v8 := 0
	for v8 < len(str) {
		v7 = int(bytes[v8]) + 131*v7
		v8++
	}
	return v7
}

func GetCurrentSeconds() int {
	return time.Now().Second()
}

func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToByte(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func ToUint32Slice(ints []int32) []uint32 {
	return *(*[]uint32)(unsafe.Pointer(&ints))
}

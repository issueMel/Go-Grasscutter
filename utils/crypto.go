package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

const prefix = "D:/IdeaProjects/Go-Grasscutter/resources"

var (
	DispatchKey       []byte
	DispatchSeed      []byte
	EncryptKey        []byte
	EncryptSeedBuffer []byte
	CurSigningKey     *rsa.PrivateKey
	EncryptionKeys    = make(map[int]*rsa.PublicKey)
)

func LoadKeys() {
	DispatchKey = readResource(prefix + "/keys/dispatchKey.bin")
	DispatchSeed = readResource(prefix + "/keys/dispatchSeed.bin")
	EncryptKey = readResource(prefix + "/keys/secretKey.bin")
	EncryptSeedBuffer = readResource(prefix + "/keys/secretKeyBuffer.bin")
	initCurSigningKey(prefix + "/keys/SigningKey.der")
	initGameKeys()
}

func initCurSigningKey(resourcePath string) {
	// 读取DER格式的私钥文件
	derBytes, err := os.ReadFile(resourcePath)
	if err != nil {
		log.Println("Error:", err)
	}
	// 解析私钥
	privateKey, err := x509.ParsePKCS8PrivateKey(derBytes)
	if err != nil {
		log.Println("Error:", err)
	}
	ok := true
	// 转换为RSA私钥类型
	CurSigningKey, ok = privateKey.(*rsa.PrivateKey)
	if !ok {
		log.Println("Error:", err)
	}
}

func initGameKeys() {
	pattern := regexp.MustCompile(`([0-9]*)_Pub\.der`)
	files, err := filepath.Glob(prefix + "/keys/game_keys/*_Pub.der")
	if err != nil {
		log.Println("Error:", err)
		return
	}

	for _, filePath := range files {
		m := pattern.FindStringSubmatch(filepath.Base(filePath))
		if len(m) > 1 {
			keyBytes, err := os.ReadFile(filePath)
			if err != nil {
				log.Println("Error reading file:", err)
				continue
			}
			key, err := x509.ParsePKIXPublicKey(keyBytes)
			if err != nil {
				log.Println("Error parsing public key:", err)
				continue
			}

			rsaKey, ok := key.(*rsa.PublicKey)
			if !ok {
				log.Println("Key is not an RSA public key")
				continue
			}

			id, err := strconv.Atoi(m[1])
			if err != nil {
				log.Println("Error converting ID to integer:", err)
				continue
			}
			EncryptionKeys[id] = rsaKey
		}
	}
}

func CreateSessionKey(length int) []byte {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println(err)
	}
	return bytes
}

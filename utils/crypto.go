package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
	"regexp"
	"strconv"
)

var (
	DispatchKey       []byte
	DispatchSeed      []byte
	EncryptKey        []byte
	EncryptSeedBuffer []byte
	CurSigningKey     *rsa.PrivateKey
	EncryptionKeys    = make(map[int]*rsa.PublicKey)
)

func LoadKeys() {
	DispatchKey = ReadResource("/keys/dispatchKey.bin")
	DispatchSeed = ReadResource("/keys/dispatchSeed.bin")
	EncryptKey = ReadResource("/keys/secretKey.bin")
	EncryptSeedBuffer = ReadResource("/keys/secretKeyBuffer.bin")
	initCurSigningKey("/keys/SigningKey.der")
	initGameKeys()
}

func initCurSigningKey(resourcePath string) {
	// 读取DER格式的私钥文件
	derBytes := ReadResource(resourcePath)
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
	files, err := r.ReadDir("resources/keys/game_keys")
	if err != nil {
		log.Println("Error:", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		m := pattern.FindStringSubmatch(file.Name())
		if len(m) > 1 {
			// filePath := filepath.Join("resources/keys/game_keys/", file.Name())
			// filepath.Join() give wrong path in Windows with embed
			filePath := "resources/keys/game_keys/" + file.Name()
			keyBytes, err := r.ReadFile(filePath)
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

package crypto

import (
	"Go-Grasscutter/src/utils"
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
	DispatchKey = utils.ReadResource("/keys/dispatchKey.bin")
	DispatchSeed = utils.ReadResource("/keys/dispatchSeed.bin")
	EncryptKey = utils.ReadResource("/keys/secretKey.bin")
	EncryptSeedBuffer = utils.ReadResource("/keys/secretKeyBuffer.bin")
	initCurSigningKey("/keys/SigningKey.der")
	initGameKeys()
}

func initCurSigningKey(resourcePath string) {
	// Read private key files in DER format
	derBytes := utils.ReadResource(resourcePath)
	// Parse private key
	privateKey, err := x509.ParsePKCS8PrivateKey(derBytes)
	if err != nil {
		log.Println("Error:", err)
	}
	ok := true
	// Convert to RSA private key type
	CurSigningKey, ok = privateKey.(*rsa.PrivateKey)
	if !ok {
		log.Println("Error:", err)
	}
}

func initGameKeys() {
	pattern := regexp.MustCompile(`([0-9]*)_Pub\.der`)
	r := utils.GetResource()
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
			keyBytes := utils.ReadResource("/keys/game_keys/" + file.Name())
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

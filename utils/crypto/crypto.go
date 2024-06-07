package crypto

import (
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/http/object"
	"Go-Grasscutter/utils"
	"Go-Grasscutter/utils/algorithms/mt19937"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	mrand "math/rand/v2"
	"regexp"
	"strconv"
)

var (
	DispatchKey       []byte
	DispatchSeed      []byte
	EncryptKey        []byte
	EncryptSeed       = uint64(11468049314633205968)
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
		log.SugaredLogger.Error("Error:", err)
	}
	ok := true
	// Convert to RSA private key type
	CurSigningKey, ok = privateKey.(*rsa.PrivateKey)
	if !ok {
		log.SugaredLogger.Error("Error:", err)
	}
}

func initGameKeys() {
	pattern := regexp.MustCompile(`([0-9]*)_Pub\.der`)
	r := utils.GetResource()
	files, err := r.ReadDir("res/keys/game_keys")
	if err != nil {
		log.SugaredLogger.Error("Error:", err)
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
				log.SugaredLogger.Error("Error parsing public key:", err)
				continue
			}
			rsaKey, ok := key.(*rsa.PublicKey)
			if !ok {
				log.SugaredLogger.Error("Key is not an RSA public key")
				continue
			}
			id, err := strconv.Atoi(m[1])
			if err != nil {
				log.SugaredLogger.Error("Error converting ID to integer:", err)
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
		log.SugaredLogger.Error("CreateSessionKey error:", err)
	}
	return bytes
}

func Xor(packet, key []byte) []byte {
	for i := 0; i < len(packet); i++ {
		packet[i] ^= key[i%len(key)]
	}
	return packet
}

func toInt(s string) int {
	i, _ := new(big.Int).SetString(s, 10)
	return int(i.Int64())
}

func EncryptAndSignRegionData(regionInfo []byte, keyID string) (*object.QueryCurRegionRspJson, error) {
	if keyID == "" {
		return nil, errors.New("key ID was not set")
	}
	// Get encryption key
	val, err := strconv.Atoi(keyID)
	if err != nil {
		return nil, err
	}
	encryptionKey, ok := EncryptionKeys[val]
	if !ok {
		return nil, errors.New("encryption key not found")
	}

	// Encrypt regionInfo in chunks
	var encryptedRegionInfo []byte
	chunkSize := 256 - 11
	regionInfoLength := len(regionInfo)
	// math.Ceil()
	numChunks := (regionInfoLength + chunkSize - 1) / chunkSize
	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > regionInfoLength {
			end = regionInfoLength
		}
		chunk := regionInfo[start:end]

		encryptedChunk, err := rsa.EncryptPKCS1v15(rand.Reader, encryptionKey, chunk)
		if err != nil {
			return nil, fmt.Errorf("encryption error: %v", err)
		}
		encryptedRegionInfo = append(encryptedRegionInfo, encryptedChunk...)
	}

	// Sign regionInfo
	hashed := sha256.Sum256(regionInfo)
	signature, err := rsa.SignPKCS1v15(rand.Reader, CurSigningKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, fmt.Errorf("signing error: %v", err)
	}

	// Base64 encode encrypted data and signature
	encodedContent := base64.StdEncoding.EncodeToString(encryptedRegionInfo)
	encodedSign := base64.StdEncoding.EncodeToString(signature)

	return &object.QueryCurRegionRspJson{
		Content: encodedContent,
		Sign:    encodedSign,
	}, nil
}

func GenerateEncryptKeyAndSeed(encryptKey []byte) ([]byte, uint64) {
	encryptSeed := mrand.Uint64()
	mt := mt19937.New()
	mt.Seed(encryptSeed)
	mt.Seed(mt.NextUint64())
	mt.NextUint64()

	for i := 0; i < 4096>>3; i++ {
		rand := mt.NextUint64()
		encryptKey[i<<3] = byte(rand >> 56)
		encryptKey[(i<<3)+1] = byte(rand >> 48)
		encryptKey[(i<<3)+2] = byte(rand >> 40)
		encryptKey[(i<<3)+3] = byte(rand >> 32)
		encryptKey[(i<<3)+4] = byte(rand >> 24)
		encryptKey[(i<<3)+5] = byte(rand >> 16)
		encryptKey[(i<<3)+6] = byte(rand >> 8)
		encryptKey[(i<<3)+7] = byte(rand)
	}
	return encryptKey, encryptSeed
}

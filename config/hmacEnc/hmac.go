package hmacEnc

import (
	hmacCrypto "crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"strings"
)

type HmacInterface interface {
	HmacEncrypt(payload string) string
	VerifyHMAC(payload string, targetHmac string) bool
}

type hmac struct {
	PRIVATE_KEY string
	ALGORYTHM   func() hash.Hash
}

func NewHmac(privateKey string, algo ...string) *hmac {
	var algorythm func() hash.Hash

	algorythms := map[string]func() hash.Hash{
		"SHA256": sha256.New,
		"SHA1":   sha1.New,
	}

	if len(algo) < 1 {
		algorythm = sha256.New
	} else if algorythms[strings.ToUpper(algo[0])] == nil {
		algorythm = sha256.New
	} else {
		algorythm = algorythms[strings.ToUpper(algo[0])]
	}

	return &hmac{
		PRIVATE_KEY: privateKey,
		ALGORYTHM:   algorythm,
	}
}

func (h *hmac) HmacEncrypt(payload string) string {
	// Konversi kunci ke byte array.
	keyBytes := []byte(h.PRIVATE_KEY)

	// Buat objek HMAC dengan fungsi hash SHA-256 dan kunci.
	enc := hmacCrypto.New(h.ALGORYTHM, keyBytes)

	// Tulis pesan ke objek HMAC.
	enc.Write([]byte(payload))

	// Ambil nilai HMAC yang dihasilkan dan konversi ke bentuk string heksadesimal.
	hmacBytes := enc.Sum(nil)
	hmacString := hex.EncodeToString(hmacBytes)

	return hmacString
}

func (h *hmac) VerifyHMAC(payload []byte, targetHmac string) bool {
	// Konversi kunci ke byte array.
	keyBytes := []byte(h.PRIVATE_KEY)

	// Buat objek HMAC dengan fungsi hash SHA-256 dan kunci.
	enc := hmacCrypto.New(h.ALGORYTHM, keyBytes)

	// Tulis pesan ke objek HMAC.
	enc.Write(payload)

	// Ambil nilai HMAC yang dihasilkan dan konversi ke bentuk string heksadesimal.
	hmacBytes := enc.Sum(nil)
	hmacString := hex.EncodeToString(hmacBytes)

	return hmacCrypto.Equal([]byte(hmacString), []byte(targetHmac))
}

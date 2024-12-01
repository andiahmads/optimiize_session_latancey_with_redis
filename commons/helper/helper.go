package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

func DynamicDir() string {
	_, b, _, _ := runtime.Caller(0)
	bStr := filepath.Dir(b)
	baseDir := strings.Replace(bStr, "commons/helper", "", -1)
	return baseDir
}

func RandomString(n int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPXRSTU"
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func SplitToken(r *http.Request) string {
	return strings.Fields(r.Header.Get("Authorization"))[1]
}

func DoSHA256(str string) string {
	hash := sha256.Sum256([]byte(str))
	hashString := hex.EncodeToString(hash[:])
	// Cetak hasil hash
	return hashString
}

package helpers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/Daniel-Sogbey/short_it/internals/models"
)

func GenerateToken(url models.OriginalUrl) string {
	hash := sha256.New()

	hash.Write([]byte(url.URL))

	timestamp := url.CreatedAt

	combinedHash := fmt.Sprintf("%s%s", base64.URLEncoding.EncodeToString(hash.Sum(nil)), timestamp)

	return combinedHash[:6]
}

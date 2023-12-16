package helpers

import (
	"strings"
	"testing"
	"time"

	"github.com/Daniel-Sogbey/short_it/internals/models"
)

func TestGenerateToken(t *testing.T) {
	token := GenerateToken(models.OriginalUrl{
		URL:       "https://www.google.com/ghanaweb.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if strings.TrimSpace(token) == "" {
		t.Errorf("Expected a token, this %s but got this %s", token, "")
	}
}

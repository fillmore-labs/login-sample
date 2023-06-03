package pkce

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/oauth2"
)

func NewVerifier(length int) (string, error) {
	// RFC7636 4.1: Result must be between 43 and 128 characters.
	if length < 32 || length > 96 {
		return "", fmt.Errorf("expected length between 32 and 96, got %d", length)
	}

	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(buf), nil
}

func AuthCodeOptions(verifier string) []oauth2.AuthCodeOption {
	sha := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(sha[:])

	return []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam(codeChallengeMethodKey, "S256"),
		oauth2.SetAuthURLParam(codeChallengeKey, challenge),
	}
}

func ExchangeOptions(verifier string) []oauth2.AuthCodeOption {
	return []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam(codeVerifierKey, verifier),
	}
}

const (
	codeChallengeKey       = "code_challenge"
	codeChallengeMethodKey = "code_challenge_method"

	codeVerifierKey = "code_verifier"
)

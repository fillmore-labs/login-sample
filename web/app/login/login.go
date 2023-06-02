package login

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"login-sample/platform/authenticator"
)

// Handler for our login.
func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		verifier := randString(43)

		sha := sha256.Sum256([]byte(verifier))
		challenge := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sha[:])
		challengeMethod := "S256"

		authCodeUrlOptions := []oauth2.AuthCodeOption{
			oauth2.SetAuthURLParam(codeChallengeKey, challenge),
			oauth2.SetAuthURLParam(codeChallengeMethodKey, challengeMethod),
		}

		// Save the state inside the session.
		session := sessions.Default(ctx)
		session.Set("state", state)
		session.Set("verifier", verifier)
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state, authCodeUrlOptions...))
	}
}

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890-._~"

func randString(n int) string {
	b := make([]byte, n)
	max := new(big.Int).SetInt64(int64(len(characters)))
	for i := range b {
		idx, _ := rand.Int(rand.Reader, max)
		b[i] = characters[idx.Int64()]
	}
	return string(b)
}

const (
	codeChallengeKey       = "code_challenge"
	codeChallengeMethodKey = "code_challenge_method"
)

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

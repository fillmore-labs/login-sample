package logout

import (
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

// Handler for our logout.
func Handler(ctx *gin.Context) {
	logoutUrl, err := url.Parse("https://" + os.Getenv("OAUTH_DOMAIN") + "/v2/logout")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", "https://"+ctx.Request.Host)
	parameters.Add("client_id", os.Getenv("OAUTH_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	ctx.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}

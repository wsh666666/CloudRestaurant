package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const CookieName = "cookie_user"
const CookieTimeLength = 10 * 60

func CookieAuth(context *gin.Context) (*http.Cookie, error) {
	cookie, err := context.Request.Cookie(CookieName)
	if err == nil {
		context.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain,cookie.SameSite, cookie.Secure, cookie.HttpOnly)
	} else {
		return nil, err
	}
	return cookie, nil
}

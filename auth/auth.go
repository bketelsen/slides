package auth

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Header(c *gin.Context, key string) string {
	if values, _ := c.Request.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}

func BasicAuth() gin.HandlerFunc {
	realm := "Authorization Required"
	realm = "Basic realm=" + strconv.Quote(realm)
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	enabled := isEnabled(user, password)
	if enabled {
		log.Warn("Auth mode enabled")
		log.Warn(fmt.Sprintf("Visit http://%s:%s@0.0.0.0:8080", user, password))
	}
	return func(c *gin.Context) {
		header := Header(c, "Authorization")
		if enabled && header != authorizationHeader(user, password) {
			// Credentials doesn't match, we return 401 and abort handlers chain.
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}

func isEnabled(user, password string) bool {
	switch {
	case user == "":
		return false
	case password == "":
		return false
	default:
		return true
	}
}

func authorizationHeader(user, password string) string {
	base := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(base))
}

package web

import (
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	smemory "github.com/ulule/limiter/v3/drivers/store/memory"
)

var sessionSecret = generateSecureToken(32)

func rateLimiterMiddleware() gin.HandlerFunc {
	return mgin.NewMiddleware(limiter.New(
		smemory.NewStore(),
		limiter.Rate{
			Period: 1 * time.Second,
			Limit:  5,
		},
	))
}

func sessionsMiddleware() gin.HandlerFunc {
	return sessions.Sessions(
		"mysession",
		cookie.NewStore([]byte(sessionSecret)),
	)
}

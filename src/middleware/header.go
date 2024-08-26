package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

const (
	headerToken    = "x-NBDS-token"
	tokenValue     = "123"
	userAgentToken = "NBDSClient/62.0.4.4878570.4789131 (Windows;10;;Professional, x64)"
)

const (
	banTime  = time.Hour
	requests = 100
)

func Headers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}

func RateLimit() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(banTime), requests)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			abortWithError(c, http.StatusTooManyRequests, http.StatusTooManyRequests, "Rate limit exceeded. Please try again later.")
			return
		}
		c.Next()
	}
}

func AuthenticateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if header := c.GetHeader(headerToken); header != tokenValue {
			abortWithError(c, http.StatusUnauthorized, http.StatusUnauthorized, "Missing or invalid auth token.")
			return
		}
		c.Next()
	}
}

func ValidateUserAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.GetHeader("User-Agent"))
		if userAgent := c.GetHeader("User-Agent"); userAgent != userAgentToken {
			abortWithError(c, http.StatusBadRequest, http.StatusBadRequest, "Missing or invalid User-Agent.")
			return
		}
		c.Next()
	}
}

func ValidateMethod(method string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != method {
			abortWithError(c, http.StatusMethodNotAllowed, http.StatusMethodNotAllowed, "invalid method.")
			return
		}
		c.Next()
	}
}

func abortWithError(c *gin.Context, statusCode int, errorCode int, errorMessage string) {
	c.JSON(statusCode, gin.H{
		"code":    errorCode,
		"message": errorMessage,
	})
	c.Abort()
}

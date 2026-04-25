package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sushantpardhi/shared/jwtToken"
)

func AuthMiddleware() gin.HandlerFunc {
	// check if user is authenticated
	return func(c *gin.Context) {
		tokenStr := ""

		// Prefer Authorization bearer token if provided
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" && strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			tokenStr = strings.TrimSpace(authHeader[7:])
		}

		if tokenStr == "" {
			var err error
			tokenStr, err = c.Cookie("access_token")
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				return
			}
		}

		// validate token
		token, err := jwt.ParseWithClaims(tokenStr, &jwtToken.Claims{}, func(t *jwt.Token) (any, error) {
			return jwtToken.GetSecretKey(), nil
		})

		// check if token is valid
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		// check if claims are valid
		claims, ok := token.Claims.(*jwtToken.Claims)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid claims"})
			return
		}

		// store in context
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	// check if user is admin
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		// check if role exists and is admin
		if !exists || role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "forbidden!! Can only be accessed by Admins"})
			return
		}
		c.Next()
	}
}

func SuperAdminOnly() gin.HandlerFunc {
	// check if user is super admin
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		// check if role exists and is super_admin
		if !exists || role != "super_admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "forbidden!! Can only be accessed by Super Admin"})
			return
		}
		c.Next()
	}
}

func AdminOrSuperAdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || (role != "admin" && role != "super_admin") {
			c.AbortWithStatusJSON(403, gin.H{"error": "forbidden!! Can only be accessed by Admin or Super Admin"})
			return
		}
		c.Next()
	}
}

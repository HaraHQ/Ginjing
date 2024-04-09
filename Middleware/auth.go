package Middleware

import (
	"encoding/json"
	"ginjing/Controllers/Authentication"
	"strings"

	"github.com/kataras/iris/v12"
)

func AuthMiddleware(ctx iris.Context) {
	// Validate token
	h := ctx.Request().Header
	authHeader := h.Get("Authorization")
	if authHeader == "" {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(map[string]string{"error": "missing Authorization header"})
		return
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(map[string]string{"error": "invalid or missing Bearer token"})
		return
	}

	// Verify token
	verifyingToken, err := Authentication.VerifyToken(tokenParts[1])
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": "failed to verify token"})
		return
	}

	// Marshal token claims to JSON
	jsonData, err := json.Marshal(verifyingToken.Claims)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": "failed to marshal token claims"})
		return
	}

	// Set user data in context
	ctx.Values().Set("user", jsonData)

	// Continue
	ctx.Next()
}

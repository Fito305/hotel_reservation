package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/Fito305/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error { // THIS IS A DECORATOR - means that we are going to modify the function so it fits what we want. It'll still return the func that it needs.
	token, ok := c.GetReqHeaders()["X-Api-Token"] 
	if !ok {
		return fmt.Errorf("Unathorized")
	}
	// I added this code. token is of type []string but needs to be a string
	stringifyToken := ""
	for _, piece := range token {
		stringifyToken += piece
	}

	 claims, err := validateToken(stringifyToken) // Was just token but gave me an error.
	 if err != nil {
		return err
	}
	expiresFloat := claims["expires"].(float64)
	expires := int64(expiresFloat)
	if time.Now().Unix() > expires {
		return fmt.Errorf("token expired")
	}
	userID := claims["id"].(string)
	user, err := userStore.GetUserByID(c.Context(), userID)
	if err != nil {
		return fmt.Errorf("unauthorized")
	}
	// Set the current authenticated user to the context.
	c.Context().SetUserValue("user", user)
	return c.Next()
  }
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		fmt.Println("NEVER PRINT SECRET", secret)
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse JWT token:", err)
		return nil, fmt.Errorf("unauthorized")
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("unauthorized")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	} 
	return claims, nil
}

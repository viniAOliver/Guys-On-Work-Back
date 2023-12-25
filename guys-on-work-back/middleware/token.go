package middleware

import (
	"guys_on_work_back/repository"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var userSystemRepository = repository.NewUserSystemRepository()

var secretKey = []byte("caiocaio")

type Login struct {

	// Login Model Attributes
	UserSystemEmail    string `json:"user_system_email"`
	UserSystemPassword string `json:"user_system_password"`
}

type Claims struct {
	UserSystemID int64 `json:"id"`
	UserSystemEmail string `json:"user_system_email"`
	jwt.StandardClaims
}

// Middleware to verified the token JWT
func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Get header from request - token JWT
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {

			c.JSON(401, gin.H{"error": "Token não fornecido na requisição"})
			c.Abort()
			return

		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.JSON(401, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("user_system_email", claims.UserSystemEmail)
			c.Next()
		} else {
			c.JSON(401, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}
	}
}

func LoginHandler(ctx *gin.Context) {
	var login Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := userSystemRepository.UserSystemByEmailDetail(login.UserSystemEmail)

	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}

	// If there is an error
	if err != nil {

		// Return JSON with the error found and status code 500 - INTERNAL ERROR
		ctx.JSON(500, gin.H{"error": err})
		return

	}

	if login.UserSystemEmail == user.UserSystemEmail {

		err := bcrypt.CompareHashAndPassword([]byte(user.UserSystemPassword), []byte(login.UserSystemPassword))

		if err == nil {

			// Create the token JWT
			expirationTime := time.Now().Add(24 * time.Hour)
			claims := &Claims{
				UserSystemID: int64(user.ID),
				UserSystemEmail: user.UserSystemEmail,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(secretKey)
			if err != nil {
				ctx.JSON(500, gin.H{"error": "Erro ao criar token"})
				return
			}

			ctx.JSON(200, gin.H{"token": tokenString})
		} else {
			ctx.JSON(401, gin.H{"error": err.Error()})
		}

	} else {

		ctx.JSON(401, gin.H{"error": "Credenciais fornecidas inválidas!"})

	}
}

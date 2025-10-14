package middlerware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type userClaims struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func Authenticate(ctx *gin.Context) {
	header := ctx.Request.Header.Get("Authorization")
	if header == "" {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
		return
	}
	tokenstring := strings.TrimPrefix(header, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenstring, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
		return
	}
	if claims, ok := token.Claims.(*userClaims); ok && token.Valid {
		ctx.Set("_id", claims.ID)
		ctx.Set("user", claims.Name)
		ctx.Set("email", claims.Email)
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		ctx.Abort()
		return
	}
	ctx.Next()
}

func GenarateToken(id, name, email string) (string, error) {
	claims := userClaims{
		ID:    id,
		Name:  name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "golang_pov",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

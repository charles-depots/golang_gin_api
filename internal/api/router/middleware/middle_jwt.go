package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "charles.shao.101" // The signature information should be set to be dynamically obtained from the library
)

// JWTAuth Middleware, check token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "The request does not carry the token and has no permission to access",
				"data":   nil,
			})
			c.Abort()
			return
		}

		log.Print("get token: ", token)
		j := NewJWT()
		// Parse the relevant information contained in the token
		claims, err := j.ParserToken(token)

		fmt.Println(claims, err)
		if err != nil {
			// token expired
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "The token authorization has expired, please reapply for authorization",
					"data":   nil,
				})
				c.Abort()
				return
			}
			// Other errors
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
				"data":   nil,
			})
			c.Abort()
			return
		}

		// Resolve to specific claims related information
		c.Set("claims", claims)

	}
}

// JWT basic data structure
// Signed signkey
type JWT struct {
	SigningKey []byte
}

// Define claim
type CustomClaims struct {
	Name  string `json:"userName"`
	Email string `json:"email"`
	// StandardClaim structure implements the Claims interface (Valid() function)
	jwt.StandardClaims
}

// Initialize the JWT instance
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// Get the signkey
func GetSignKey() string {
	return SignKey
}

func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// Create Token (based on the user's basic information claims)
// Use HS256 algorithm for token generation
// Use user basic information claims and signature key (signkey) to generate token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	// https://gowalker.org/github.com/dgrijalva/jwt-go#Token
	// Return a token structure pointer
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// Token parsing
// Couldn't handle this token:
func (j *JWT) ParserToken(tokenString string) (*CustomClaims, error) {
	// https://gowalker.org/github.com/dgrijalva/jwt-go#ParseWithClaims
	// Enter the user-defined Claims structure object, token, and custom function to parse the token string into jwt's Token structure pointer
	// Keyfunc is an anonymous function type: type Keyfunc func(*Token) (interface{}, error)
	// func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	fmt.Println(token, err)
	if err != nil {
		// https://gowalker.org/github.com/dgrijalva/jwt-go#ValidationError
		// jwt.ValidationError is an incorrect structure of an invalid token
		if ve, ok := err.(*jwt.ValidationError); ok {
			// ValidationErrorMalformed is a uint constant, indicating that the token is not available
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
				// ValidationErrorExpired表示Token过期
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
				// ValidationErrorNotValidYet indicates invalid token
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	// Parse out the claims information in the token and verify the original data of the user
	// Make the following type assertions to convert token.Claims into a specific user-defined Claims structure
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, TokenInvalid

}

// Update Token
func (j *JWT) UpdateToken(tokenString string) (string, error) {
	// TimeFunc is a current time variable whose default value is time.Now, which is used to verify the expiration time after parsing the token
	// You can use other time values to override
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	// Acquiring basic data token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil

	})

	// Current also effective verification token
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		// Modify the expiration time of Claims(int64)
		// https://gowalker.org/github.com/dgrijalva/jwt-go#StandardClaims
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", fmt.Errorf("Token acquisition failed:%v", err)
}

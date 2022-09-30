package middleware

import (
	"time"

	"github.com/kataras/iris/v12"

	"github.com/iris-contrib/middleware/jwt"
)

var mySecret = []byte("secret")

var J = jwt.New(jwt.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	},
	Expiration: true,

	Extractor: jwt.FromAuthHeader,

	SigningMethod: jwt.SigningMethodHS256,
})

// generate token to use.
func GetTokenHandler(email string) string {
	now := time.Now()
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"iat":   now.Unix(),
		"exp":   now.Add(15 * time.Minute).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(mySecret)
	return tokenString

}

func MyAuthenticatedHandler(ctx iris.Context) string {
	if err := J.CheckJWT(ctx); err != nil {
		J.Config.ErrorHandler(ctx, err)

	}

	token := ctx.Values().Get("jwt").(*jwt.Token)

	tokenResult := token.Claims.(jwt.MapClaims)["email"].(string)

	ctx.Next()
	// for key, value := range foobar {
	// 	ctx.Writef("%s = %s", key, value)
	// }

	return tokenResult
}

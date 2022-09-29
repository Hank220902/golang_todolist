package middleware

import (
	"time"

	"github.com/kataras/iris/v12"

	"github.com/iris-contrib/middleware/jwt"
)

var mySecret = []byte("secret")

var j = jwt.New(jwt.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	},
	Expiration: true,

	Extractor: jwt.FromParameter("token"),

	SigningMethod: jwt.SigningMethodHS256,
})

// generate token to use.
func GetTokenHandler(email string) string {
	now := time.Now()
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"iat": now.Unix(),
		"exp": now.Add(15 * time.Minute).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(mySecret)
	return tokenString
	
}

func MyAuthenticatedHandler(ctx iris.Context) {
	if err := j.CheckJWT(ctx); err != nil {
		j.Config.ErrorHandler(ctx, err)
		return
	}

	token := ctx.Values().Get("jwt").(*jwt.Token)

	ctx.Writef("This is an authenticated request\n\n")
	// ctx.Writef("Claim content:\n")

	foobar := token.Claims.(jwt.MapClaims)
	ctx.Writef("foo=%s\n", foobar["foo"])
	// for key, value := range foobar {
	// 	ctx.Writef("%s = %s", key, value)
	// }
}

// func main() {
// 	app := iris.New()

// 	app.Get("/", getTokenHandler)
// 	app.Get("/protected", myAuthenticatedHandler)

// 	// j.CheckJWT(Context) error can be also used inside handlers.

// 	app.Listen(":8080")
// }

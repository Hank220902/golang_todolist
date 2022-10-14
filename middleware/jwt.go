package middleware

import (
	"context"
	"fmt"
	"time"

	"todolist/connect"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

var mySecret = []byte("secret")

var J = jwt.New(jwt.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	},
	Expiration: false,

	Extractor: jwt.FromAuthHeader,

	SigningMethod: jwt.SigningMethodHS256,
})

// generate token to use.
func GetTokenHandler(email string) string {
	var Redis = connect.Redis()
	ctx := context.Background()

	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,

	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(mySecret)

	n, err := Redis.Exists(ctx, email).Result()
	if err != nil {
		panic(err)
	}
	if n > 0 {
		n := Redis.Del(ctx, email)
		fmt.Println(n.Result())
		err := Redis.Set(ctx, email, tokenString, 15*time.Minute).Err()
		if err != nil {
			panic(err)
		}
	} else {
		err := Redis.Set(ctx, email, tokenString, 15*time.Minute).Err()
		if err != nil {
			panic(err)
		}
	}

	return tokenString

}

func MyAuthenticatedHandler(ctx iris.Context) string {
	var Redis = connect.Redis()
	if err := J.CheckJWT(ctx); err != nil {
		J.Config.ErrorHandler(ctx, err)

	}

	token := ctx.Values().Get("jwt").(*jwt.Token)

	tokenResult := token.Claims.(jwt.MapClaims)["email"].(string)
	n, err := Redis.Exists(ctx, tokenResult).Result()
	if err != nil {
		panic(err)
	}
	if n > 0 {
		ctx.Next()

		return tokenResult
	} else {
		return "token not found"
	}

}

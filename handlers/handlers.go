package handlers

import (
	"context"
	"fmt"

	"github.com/Jonda-HR/Goland_twitter/v2/jwt"
	"github.com/Jonda-HR/Goland_twitter/v2/models"
	"github.com/Jonda-HR/Goland_twitter/v2/routers"
	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("I Will process" + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var respons models.RespApi

	respons.Status = 400

	isOk, statusCode, msg, claim := validAuthorization(ctx, request)
	if !isOk {
		respons.Status = statusCode
		respons.Message = msg
		return respons
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "signup":
			return routers.SignIn(ctx)

		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
	}

	respons.Message = "Method Invalid"
	return respons
}

func validAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "signup" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]

	if len(token) == 0 {
		return false, 401, "required token", models.Claim{}
	}

	claim, isOk, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtSing")).(string))

	if !isOk {
		if err != nil {
			fmt.Println("Token Error: " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Token Error " + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg, *claim
}

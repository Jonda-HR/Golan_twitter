package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Jonda-HR/Goland_twitter/v2/db"
	"github.com/Jonda-HR/Goland_twitter/v2/models"
)

func SignIn(ctx context.Context) models.RespApi {
	var user models.User
	var response models.RespApi
	response.Status = 400

	fmt.Println("SIGN UP")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)
	if err != nil {
		response.Message = err.Error()
		fmt.Println(response.Message)
		return response
	}

	if len(user.Email) == 0 {
		response.Message = "mail is needed"
		fmt.Println(response.Message)
		return response
	}

	if len(user.Password) < 8 {
		response.Message = "A password of at least 8 characters is required"
		fmt.Println(response.Message)
		return response
	}

	_, exist, _ := db.UserExist(user.Email)

	if exist {
		response.Message = "User Exists"
		fmt.Println(response.Message)
		return response
	}

	_, status, err := db.InsertSignUp(user)
	if err != nil {
		response.Message = "An ocurred error while inserting signup"
		fmt.Println(response.Message)
		return response
	}

	if !status {
		response.Message = "Failed to insert user record"
	}

	response.Status = 200
	response.Message = "signUp Ok"
	return response
}

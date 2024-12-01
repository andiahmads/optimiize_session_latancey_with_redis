package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"gored/commons/exception"
	"gored/commons/helper"
	"gored/pkg"
	"time"
)

func loginUsescase(ctx context.Context, userNama, password string) (userItems, error) {
	data, err := getUserByEmail(ctx, pkg.Dbx(), userNama)
	if err != nil {
		fmt.Println(err)
		return userItems{}, exception.ErrAccountNotFound
	}

	hashPass := helper.DoSHA256(password)
	if data.Password != hashPass {
		return userItems{}, exception.ErrAccountNotFound
	}

	// set token and redis
	token := helper.RandomString(50)
	toMarshal, err := json.Marshal(data)
	if err != nil {
		return userItems{}, err
	}

	if err := pkg.Rdb().Set(ctx, token, toMarshal, 1*time.Minute).Err(); err != nil {
		return userItems{}, err
	}

	response := userItems{
		UserID:      data.UserID,
		Username:    data.Username,
		AccessToken: token,
		CreatedAt:   data.CreatedAt,
	}

	return response, nil
}

func getProductUseCase(ctx context.Context) ([]productItems, error) {
	return getProduct(ctx, pkg.Dbx())
}

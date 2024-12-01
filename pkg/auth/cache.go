package auth

import (
	"context"
	"encoding/json"
	"gored/commons/exception"
	"gored/pkg"
)

func User(ctx context.Context, token string) (userItems, error) {
	data, err := pkg.Rdb().Get(ctx, token).Bytes()
	if err != nil {
		return userItems{}, exception.ErrExpToken
	}

	var userData userItems

	if err := json.Unmarshal(data, &userData); err != nil {
		return userItems{}, err
	}

	return userData, nil
}

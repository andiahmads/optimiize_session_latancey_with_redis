package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"gored/commons/exception"
	"gored/commons/response"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JsonEncoder[string](w, http.StatusBadRequest, exception.ErrBadRequest.Error(), false, "")
		return
	}

	data, err := loginUsescase(context.Background(), req.Username, req.Password)
	fmt.Println(err)
	if err != nil {
		response.JsonEncoder[string](w, http.StatusBadRequest, exception.ErrBadRequest.Error(), false, "")
		return
	}
	response.JsonEncoder[userItems](w, http.StatusOK, exception.Success, true, data)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getProductUseCase(context.Background())
	if err != nil {
		response.JsonEncoder[string](w, http.StatusBadRequest, exception.ErrBadRequest.Error(), false, "")
		return
	}

	response.JsonEncoder[[]productItems](w, http.StatusOK, exception.Success, false, data)
}

package auth

type (
	userItems struct {
		UserID      string `json:"id"`
		Username    string `json:"username"`
		Password    string `json:"password,omitempty"`
		AccessToken string `json:"access_token"`
		CreatedAt   string `json:"created_at"`
	}

	productItems struct {
		ProductId string `json:"id"`
		Name      string `json:"name"`
		Stock     int    `json:"stock"`
		CreatedAt string `json:"created_at"`
	}
)

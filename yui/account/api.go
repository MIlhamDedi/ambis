package account

type (
	CreateUser struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	VerifyUser struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
	}
)

package auth

type (
	CreateToken struct {
		Authentication Authentication
		ExpirationTime int64
	}

	VerifyToken struct {
		Token string
	}
)

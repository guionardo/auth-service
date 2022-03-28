package domain

type (
	UserData struct {
		UserID       string
		PasswordHash string
		Payload      interface{}
	}
)

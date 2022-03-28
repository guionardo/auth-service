package dto

type PostUserDTO struct {
	UserId       string      `json:"userid"`
	PasswordHash string      `json:"password_hash"`
	Payload      interface{} `json:"payload"`
}

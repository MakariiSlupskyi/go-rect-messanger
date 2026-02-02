package user

type User struct {
	ID           int64
	Username     string
	DisplayName  *string
	PasswordHash string
}

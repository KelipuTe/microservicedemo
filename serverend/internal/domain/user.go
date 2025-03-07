package domain

type User struct {
	Id       int64
	Username string
	Password string
	Phone    string
	Email    string
}

type VerifyCode struct {
	Id        int64
	Biz       string
	Target    string
	Code      string
	ExpiresAt int64
	IsUsed    int
}

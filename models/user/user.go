package user

import "time"

// email unique

type UserEntity struct {
	Email    string // email unique
	Username string // user name not unique
	Password string
	Phone    string

	RegisterTime time.Time
	RegisterIp   string

	LastLoginTime time.Time
	LastLoginIp   string
}

type UserToken = string

type UserModel interface {
	Login(user, password string) (UserToken, error)
	IsLogin(ut UserToken) *UserEntity
	Logout(ut UserToken)
	Register(ue *UserEntity) error
	SendVerifyCodeByEmail(email string) error
	VerifyEmail(verifyCode string) error

	ResetPassword(ue *UserEntity, pwd string) error
}

// interface for controller or biz
// for register
func Create(e *UserEntity) error {
	return nil
}

// for login
func CheckPwd(email, pwd string) error {
	return nil
}

func GetByEmail(email string) (*UserEntity, error) {
	return nil, nil
}

func VerifyEmail(verifyCode string) error {
	return nil
}

func SendVerifyCodeByEmail(email string) error {
	return nil
}

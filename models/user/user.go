package user

import "time"

// email unique

type UserEntity struct {
	Email string	// email unique
	Username string	// user name not unique
	Password string
	Phone string

	RegisterTime time.Time
	RegisterIp string

	LastLoginTime time.Time
	LastLoginIp string
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


package user

import "time"

// email unique

type UserEntity struct {
	Email    string // email unique
	Nickname string // nick name not unique
	Password string
	Phone    string

	RegisterTime time.Time
	RegisterIp   string

	LastLoginTime time.Time
	LastLoginIp   string
}



// interface for controller or biz
// for register
func Add(e *UserEntity) error {
	return nil
}

func Update(e *UserEntity) error {
	return nil
}

func Delete(e *UserEntity) error {
	return nil
}

func Get(e *UserEntity) error {
	return nil
}

func GetByEmail(email string) (*UserEntity, error) {
	return nil, nil
}

func GetByPhone(phone string) (*UserEntity, error) {
	return nil, nil
}


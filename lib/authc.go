package lib

import (
	"errors"
	"time"

	"github.com/ikeikeikeike/gopkg/convert"

	"github.com/uxff/taniago/models"
)

/*
 Get authenticated user and update logintime
*/
func Authenticate(email string, password string) (user *models.User, err error) {
	msg := "invalid email or password."
	user = &models.User{Email: email}

	if err := user.Read("Email"); err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			err = errors.New(msg)
		}
		return user, err
	} else if user.Id < 1 {
		// No user
		return user, errors.New(msg)
	} else if user.Password != convert.StrTo(password).Md5() {
		// No matched password
		return user, errors.New(msg)
	} else {
		user.Lastlogintime = time.Now()
		user.Update("Lastlogintime")
		return user, nil
	}
}

func SignupUser(u *models.User) (int64, error) {
	var (
		err error
		msg string
	)

	if models.Users().Filter("email", u.Email).Exist() {
		msg = "was already regsitered input email address."
		return 0, errors.New(msg)
	}

	u.Password = convert.StrTo(u.Password).Md5()

	err = u.Insert()
	if err != nil {
		return 0, err
	}

	return u.Id, err
}

package user

import (
	"errors"
	"fmt"
	model "lets-blog/pkg/models/user"
	repos "lets-blog/pkg/repositories/user"
	"time"
)

func SignUpUser(user *model.UserReqModel) error {
	if d,_ := repos.GetUserUsingEmail(user.UserEmail); d != nil {
		return errors.New("email already exists")
	}

	userId := fmt.Sprintln(time.Now())
	createdAt := time.Now()

	userToSave := model.UserModel{
		UserId: userId,
		UserName: user.UserName,
		UserEmail: user.UserEmail,
		Password: user.Password,
		CreatedAt: createdAt.Format("02/01/2006"),
	}

	if err := repos.CreateNewUser(userToSave); err != nil {
		return err
	}
	return nil
}

func LoginUser()  {
	
}

func LogoutUser()  {
	
}

func GetUser()  {
	
}
package usecase

import (
	model "api/models/user"
	repository "api/repositories/user"
)

type UserUseCase struct{
  repository repository.UserRepository
}


func NewUserUseCase(repo repository.UserRepository) UserUseCase  {
	return UserUseCase{
		repository: repo,
	}
}


func (users * UserUseCase) GetUsers()([]model.User, error){
	return users.repository.GetUsers();
}


func (us * UserUseCase) CreateUser(user model.User)(model.User, error){
	newUser, err := us.repository.CreateUser(user)

	if err != nil{
		return model.User{},err
	}

	return newUser, nil
}


func (us *UserUseCase) GetUserById(id int) (*model.User, error) {
	return us.repository.GetUserById(id)
}



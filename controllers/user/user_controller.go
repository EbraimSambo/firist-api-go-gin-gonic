package controller

import (
	model "api/models/user"
	usecase "api/usecases/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) UserController {
	return UserController{
		userUsecase: usecase,
	}
}

func (u *UserController) GetUsers(ctx *gin.Context) {
	
	users, err := u.userUsecase.GetUsers();

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}


func (uc *UserController) CreateUser(ctx *gin.Context){
	var user = model.User{}
	err := ctx.BindJSON(&user)

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err)
		return
	}

	insertUser, err := uc.userUsecase.CreateUser(user) 

	if err != nil{
		ctx.JSON(http.StatusInternalServerError,err)
		return
	}

	ctx.JSON(http.StatusCreated, insertUser)
}
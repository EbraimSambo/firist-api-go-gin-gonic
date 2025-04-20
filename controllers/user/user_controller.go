package controller

import (
	"api/models/response"
	model "api/models/user"
	usecase "api/usecases/user"
	"net/http"
	"strconv"

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

	users, err := u.userUsecase.GetUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user = model.User{}
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	insertUser, err := uc.userUsecase.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertUser)
}
func (u *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, response.ResponseError{
			Status:  400,
			Message: "O parâmetro id é obrigatório",
		})
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseError{
			Status:  400,
			Message: "ID inválido",
		})
		return
	}

	user, err := u.userUsecase.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseError{
			Status:  500,
			Message: "Erro interno ao buscar o usuário",
		})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, response.ResponseError{
			Status:  404,
			Message: "Usuário não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

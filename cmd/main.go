package main

import (
	controller "api/controllers/user"
	"api/database"
	repository "api/repositories/user"
	usecase "api/usecases/user"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnectin,err := database.ConnectDB()
	if err != nil{
		panic(err)
	}
	UserRepository := repository.NewUserRepository(dbConnectin)
	UserUsecase := usecase.NewUserUseCase(UserRepository)
	UserController := controller.NewUserController(UserUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World GIn Gonic",
		})
	})

	server.GET("/users", UserController.GetUsers);
	server.POST("/users", UserController.CreateUser);
	server.POST("/users/:id", UserController.GetUser);
	
	server.Run(":8000")
}

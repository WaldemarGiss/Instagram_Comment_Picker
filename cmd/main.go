package main

import (
	"Instagram_Comment_Picker/api"
	"Instagram_Comment_Picker/internal/repository"
	"Instagram_Comment_Picker/internal/service"
	"github.com/gin-gonic/gin"
)

// init is invoked before main()
/*func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}*/

func main() {
	router := gin.Default()

	repository := repository.ProvideInstaPickerRepository()

	service := service.ProvideInstaPickerService(repository)

	api.CreateRouter(router, service)

	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}

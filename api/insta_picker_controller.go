package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type instaPickerService interface {
	DoSomething()
}

type InstaPickerController struct {
	service instaPickerService
}

func ProvideInstaPickerController(service instaPickerService) *InstaPickerController {
	return &InstaPickerController{service: service}
}

func (controller *InstaPickerController) DoSomething(req *gin.Context) {
	fmt.Print("nice")
}

//Root Router to relative Path
func CreateRouter(router *gin.Engine, Service instaPickerService) {
	controller := ProvideInstaPickerController(Service)
	router.GET("HVB_Stock_API/calcGain", controller.DoSomething)
}

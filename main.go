package main

import (
	"io"
	"net/http"
	"os"

	repository "github.com/Angieuski/Desafio-Edu/Repository"
	"github.com/Angieuski/Desafio-Edu/controller"
	"github.com/Angieuski/Desafio-Edu/middlewares"
	"github.com/Angieuski/Desafio-Edu/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	uiRepository    repository.UiRepository = repository.NewRepository()
	UiServiceStruct service.UiService       = service.New(uiRepository)
	UiController    controller.UiController = controller.New(UiServiceStruct)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	defer uiRepository.CloseDB()

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(),
		middlewares.BasicAuth(), gindump.Dump())

	server.GET("/uis", func(ctx *gin.Context) {
		ctx.JSON(200, UiController.FindAll())
	})

	server.POST("/uis", func(ctx *gin.Context) {
		err := UiController.Save(ctx)
		if err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Entrada de Ui válida"})
		}
	})

	server.PUT("/uis/:id", func(ctx *gin.Context) {
		err := UiController.Update(ctx)
		if err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Entrada de Ui válida"})
		}
	})

	server.DELETE("/uis/:id", func(ctx *gin.Context) {
		err := UiController.Delete(ctx)
		if err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Entrada de Ui válida"})
		}
	})

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}

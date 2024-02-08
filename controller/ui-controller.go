package controller

import (
	"errors"
	"strconv"

	"github.com/Angieuski/Desafio-Edu/entity"
	"github.com/Angieuski/Desafio-Edu/service"
	"github.com/gin-gonic/gin"
)

type UiController interface {
	FindAll() []entity.Ui
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.UiService
}

func New(service service.UiService) UiController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Ui {
	return c.service.FindAll()
}

func (c *controller) Update(ctx *gin.Context) error {
	var ui entity.Ui

	err := ctx.ShouldBindJSON(&ui)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	ui.ID = id

	existingUser := c.service.FindByEmail(ui.Email)
	if existingUser.ID > 0 && existingUser.ID != ui.ID {
		return errors.New("E-mail já está em uso por outro usuário")
	}

	c.service.Update(ui)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var ui entity.Ui
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	ui.ID = id
	c.service.Delete(ui)
	return nil
}

func (c *controller) Save(ctx *gin.Context) error {
	var ui entity.Ui
	err := ctx.ShouldBindJSON((&ui))
	if err != nil {
		return err
	}
	existingUser := c.service.FindByEmail(ui.Email)
	if existingUser.ID > 0 {
		return errors.New("E-mail já está em uso por outro usuário")
	}
	c.service.Save(ui)
	return nil
}

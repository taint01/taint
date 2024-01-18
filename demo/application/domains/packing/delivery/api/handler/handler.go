package handler

import (
	"demo/application/domains/packing/entity"
	"demo/application/domains/packing/models"
	"demo/application/domains/packing/usecase"
	"demo/application/model"
	"demo/config"
	"github.com/gin-gonic/gin"
	"strconv"

	"context"
)

type Handler struct {
	lib     *model.Lib
	cfg     *config.Config
	usecase usecase.IUseCase
}

func InitHandler(
	lib *model.Lib,
	cfg *config.Config,
	usecase usecase.IUseCase,
) Handler {
	return Handler{
		lib,
		cfg,
		usecase,
	}
}

func (h *Handler) Read(ctx context.Context, c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	name := c.Query("name")
	age, _ := strconv.Atoi(c.Query("age"))
	res, err := h.usecase.Read(ctx, &models.Params{Id: int64(id), Name: name, Age: int64(age)})
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	}

	c.JSON(200, gin.H{
		"response": res,
	})
}

func (h *Handler) Insert(ctx context.Context, c *gin.Context) {

	var user *entity.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(404, gin.H{
			"error": "can't bind JSON insert",
		})
	}
	err := h.usecase.Insert(ctx, user)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	}
	if err == nil {
		c.JSON(200, gin.H{
			"response": "insert ok",
		})
	}

}

func (h *Handler) Update(ctx context.Context, c *gin.Context) {

	var user *entity.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(404, gin.H{
			"error": "can't bind JSON update",
		})
	}
	err := h.usecase.Update(ctx, user)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	}
	if err == nil {
		c.JSON(200, gin.H{
			"response": "update ok",
		})
	}
}

func (h *Handler) Delete(ctx context.Context, c *gin.Context) {

	var user *entity.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(404, gin.H{
			"error": "can't bind JSON delete",
		})
	}
	err := h.usecase.Delete(ctx, user)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
	}
	if err == nil {
		c.JSON(200, gin.H{
			"response": "delete ok",
		})
	}
}

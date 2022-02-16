package handler

import (
	"log"

	"github.com/anime454/go_hexagonal_architecture/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userSv service.UserService
}

func NewUserHandler(us service.UserService) userHandler {
	return userHandler{userSv: us}
}

func (uhdl userHandler) Register() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		res, err := uhdl.userSv.GetAllUsers()
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{
				"code":    50000,
				"message": "Internal Server Error",
				"data":    nil,
			})
		}
		c.JSON(200, gin.H{
			"code":    20000,
			"message": "Success",
			"data":    res,
		})
	}
	return fn
}

func (uhdl userHandler) GetAll() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		res, err := uhdl.userSv.GetAllUsers()
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{
				"code":    50000,
				"message": "Internal Server Error",
				"data":    nil,
			})
		}
		c.JSON(200, gin.H{
			"code":    20000,
			"message": "Success",
			"data":    res,
		})
	}
	return fn
}

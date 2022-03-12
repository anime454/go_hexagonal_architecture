package handler

import (
	"log"

	"github.com/anime454/go_hexagonal_architecture/logs"
	"github.com/anime454/go_hexagonal_architecture/responses"
	"github.com/anime454/go_hexagonal_architecture/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userSv service.UserService
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserDetail struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func NewUserHandler(us service.UserService) userHandler {
	return userHandler{userSv: us}
}

var Err500 = map[string]interface{}{
	"Code":    50000,
	"Message": "Internal Server Error",
	"Data":    nil,
}

func (uhdl userHandler) Register() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		u := User{}
		err := c.Bind(&u)
		if err != nil {
			logs.Error(err)
			c.JSON(500, responses.InternalServerError())
			return
		}

		user := service.User{
			Username: u.Username,
			Password: u.Password,
			FullName: u.FullName,
			Email:    u.Email,
			Role:     u.Role,
		}

		res, err := uhdl.userSv.Register(user)
		if err != nil {
			logs.Error(err)
			c.JSON(500, Err500)
			return
		}

		c.JSON(200, gin.H{
			"code":    20000,
			"message": "Success",
			"data":    res,
		})
	}
	return fn
}

func (uhdl userHandler) GetAllUser() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		res, err := uhdl.userSv.GetAllUsers()
		if err != nil {
			log.Fatal(err)
			c.JSON(500, Err500)
			return
		}
		c.JSON(200, gin.H{
			"code":    20000,
			"message": "Success",
			"data":    res,
		})
	}
	return fn
}

func (uhdl userHandler) GetUserById() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		res, err := uhdl.userSv.GetUserById(id)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, Err500)
			return
		}
		c.JSON(200, gin.H{
			"code":    20000,
			"message": "Success",
			"data":    res,
		})
	}
	return fn
}

func (uhdl userHandler) UpdateUser() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		u := UserDetail{}
		err := c.Bind(&u)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, Err500)
			return
		}

		user := service.UserDetail{
			Id:       u.Id,
			Username: u.Username,
			FullName: u.FullName,
			Email:    u.Email,
			Role:     u.Role,
		}

		res, err := uhdl.userSv.UpdateUser(user)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, Err500)
			return
		}
		c.JSON(200, gin.H{
			"code":    20000,
			"message": "Success",
			"data":    res,
		})

	}
	return fn
}

func (uhdl userHandler) DeleteUserById() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		err := uhdl.userSv.DeleteUser(id)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, Err500)
			return
		}
		c.JSON(200, gin.H{
			"code":    20000,
			"message": "Success",
			"data":    nil,
		})
	}
	return fn
}

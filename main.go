package main

import (
	"database/sql"
	"net/http"

	"github.com/anime454/go_hexagonal_architecture/handler"
	"github.com/anime454/go_hexagonal_architecture/logs"
	"github.com/anime454/go_hexagonal_architecture/middleware"
	"github.com/anime454/go_hexagonal_architecture/repository"
	"github.com/anime454/go_hexagonal_architecture/service"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./tmp.db")
	if err != nil {
		logs.Error(err)
	}

	const createTable string = `
	CREATE TABLE IF NOT EXISTS user (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		fullname VARCHAR(255) NOT NULL,
		email VARCHAR(255) NULL,
		role VARCHAR(255) NOT NULL DEFAULT "user", 
		auto_datetime DATETIME NOT NULL
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		logs.Error(err)
	}

	userRepo := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	gin.SetMode("release")
	server := gin.Default()
	// server := gin.New()
	normalRoute := server.Group("/")
	normalRoute.Use(middleware.InitAccessLog())
	{
		normalRoute.POST("/register", userHandler.Register())
		normalRoute.GET("/getAllUser", userHandler.GetAllUser())
		normalRoute.GET("/getUserById/:id/", userHandler.GetUserById())
		normalRoute.POST("/updateUser", userHandler.UpdateUser())
		normalRoute.POST("/deleteUser/:id/", userHandler.DeleteUserById())
	}

	srv := &http.Server{
		Addr:    ":" + "9090",
		Handler: server,
	}

	logs.Info("Start server on " + srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logs.Error(err)
	}

}

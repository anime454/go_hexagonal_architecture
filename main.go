package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/anime454/go_hexagonal_architecture/handler"
	"github.com/anime454/go_hexagonal_architecture/logs"
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
	server.POST("/register", res(), userHandler.Register())
	server.GET("/getAllUser", userHandler.GetAllUser())
	server.GET("/getUserById/:id/", userHandler.GetUserById())
	server.POST("/updateUser", userHandler.UpdateUser())
	server.POST("/deleteUser/:id/", userHandler.DeleteUserById())

	srv := &http.Server{
		Addr:    ":" + "9090",
		Handler: server,
	}

	logs.Info("Start server on " + srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logs.Error(err)
	}

}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	fmt.Println("on write")
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	fmt.Println("on writeString")
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func res() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		fmt.Println(c.Request)
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		responseTime := time.Since(now)
		fmt.Println("Method: " + c.Request.Method + " URI: " + c.Request.RequestURI + " Time: " + responseTime.String() + " Response body: " + blw.body.String())
	}
}

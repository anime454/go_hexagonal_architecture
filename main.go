package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/anime454/go_hexagonal_architecture/handler"
	"github.com/anime454/go_hexagonal_architecture/repository"
	"github.com/anime454/go_hexagonal_architecture/service"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./tmp.db")
	if err != nil {
		log.Fatalln("Can't create Database", err)
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
		log.Fatalln("Can't create Table", err)
	}

	// id := uuid.NewString()
	// mockuser := service.User{
	// 	Id:           id,
	// 	Username:     "admin_service",
	// 	Password:     "1234",
	// 	FullName:     "admin admin",
	// 	Email:        "admin@admin.com",
	// 	Role:         "admin",
	// 	AutoDatetime: time.Now(),
	// }

	userRepo := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	server := gin.Default()
	server.GET("/getAll", userHandler.GetAll())

	srv := &http.Server{
		Addr:    ":" + "9090",
		Handler: server,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}

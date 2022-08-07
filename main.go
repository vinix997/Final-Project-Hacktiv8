package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/vinix997/Final-Project-Hacktiv8/handler"
	"github.com/vinix997/Final-Project-Hacktiv8/helper"
	"github.com/vinix997/Final-Project-Hacktiv8/middleware"
	"github.com/vinix997/Final-Project-Hacktiv8/repository"
	"github.com/vinix997/Final-Project-Hacktiv8/service"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error: ", r)
		}
	}()
	//Connection binding
	conn := Connection()

	//Repository
	userRepo := repository.NewUserRepository(conn)
	photoRepo := repository.NewPhotoRepository(conn)
	commentRepo := repository.NewCommentRepository(conn)
	socialRepo := repository.NewSocialRepository(conn)

	//Service
	userService := service.NewUserSvc(userRepo)
	photoService := service.NewPhotoSvc(photoRepo)
	commentService := service.NewCommentSvc(commentRepo)
	socialService := service.NewSocialSvc(socialRepo)

	//Handler
	userHandler := handler.NewUserHandler(userService, conn)
	photoHandler := handler.NewPhotoHandler(photoService, conn)
	commentHandler := handler.NewCommentHandler(commentService, conn)
	socialHandler := handler.NewSocialHandler(socialService, conn)

	authMiddleware := middleware.NewAuthMiddleware()

	//Endpoints
	r := mux.NewRouter()
	userRoutes := r.PathPrefix("/users").Subrouter()
	userRoutes.HandleFunc("/register", userHandler.UserHandler)
	userRoutes.HandleFunc("/login", userHandler.UserHandler)
	userRoutes.HandleFunc("/{userId}", userHandler.UserHandler)

	photoRoutes := r.PathPrefix("/photos").Subrouter()
	photoRoutes.HandleFunc("", photoHandler.PhotoHandler)
	photoRoutes.HandleFunc("/{photoId}", photoHandler.PhotoHandler)

	commentRoutes := r.PathPrefix("/comments").Subrouter()
	commentRoutes.HandleFunc("/{commentId}", commentHandler.CommentHandler)
	commentRoutes.HandleFunc("", commentHandler.CommentHandler)

	socialRoutes := r.PathPrefix("/socialmedias").Subrouter()
	socialRoutes.HandleFunc("", socialHandler.SocialHandler)
	socialRoutes.HandleFunc("/{socialMediaId}", socialHandler.SocialHandler)

	//Middleware
	userRoutes.Use(authMiddleware.AuthMiddleware)
	photoRoutes.Use(authMiddleware.AuthMiddleware)
	commentRoutes.Use(authMiddleware.AuthMiddleware)
	socialRoutes.Use(authMiddleware.AuthMiddleware)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func Connection() *sql.DB {
	user := helper.GetEnvVar("DB_USERNAME")
	password := helper.GetEnvVar("PASSWORD")
	db := helper.GetEnvVar("DB_NAME")
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", user, password, db))
	if err != nil {
		panic(err)
	}
	return database
}

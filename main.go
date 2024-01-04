package main

import (
	"fchramn/movies-restfulapi/app"
	"fchramn/movies-restfulapi/controller"
	"fchramn/movies-restfulapi/exception"
	"fchramn/movies-restfulapi/helper"
	"fchramn/movies-restfulapi/middleware"
	"fchramn/movies-restfulapi/repository"
	"fchramn/movies-restfulapi/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	movieRepository := repository.NewMovieRepository()
	movieService := service.NewMovieService(movieRepository, db, validate)
	movieController := controller.NewMovieController(movieService)

	router := httprouter.New()

	router.GET("/movies", movieController.FindAll)
	router.GET("/movies/:id", movieController.FindById)
	router.POST("/movies", movieController.Create)
	router.PATCH("/movies/:id", movieController.Update)
	router.DELETE("/movies/:id", movieController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

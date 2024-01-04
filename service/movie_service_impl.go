package service

import (
	"context"
	"database/sql"
	"fchramn/movies-restfulapi/exception"
	"fchramn/movies-restfulapi/helper"
	"fchramn/movies-restfulapi/model/domain"
	"fchramn/movies-restfulapi/model/web"
	"fchramn/movies-restfulapi/repository"

	"github.com/go-playground/validator/v10"
)

type MovieServiceImpl struct {
	MovieRepository repository.MovieRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewMovieService(MovieRepository repository.MovieRepository, DB *sql.DB, validate *validator.Validate) MovieService {
	return &MovieServiceImpl{
		MovieRepository: MovieRepository,
		DB:              DB,
		Validate:        validate,
	} //polymorphism repository -> service -> controller
}

func (service *MovieServiceImpl) FindAll(ctx context.Context) []web.MovieResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movies := service.MovieRepository.FindAll(ctx, tx)

	return helper.ToMovieResponses(movies)
}

func (service *MovieServiceImpl) FindById(ctx context.Context, movieId int) web.MovieResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movie, err := service.MovieRepository.FindById(ctx, tx, movieId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMovieResponse(movie)
}

func (service *MovieServiceImpl) Create(ctx context.Context, request web.MovieCreateRequest) web.MovieResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movie := domain.Movie{
		Title:       request.Title,
		Description: request.Description,
		Rating:      request.Rating,
		Image:       request.Image,
	}

	movie = service.MovieRepository.Save(ctx, tx, movie)

	return helper.ToMovieResponse(movie)
}

func (service *MovieServiceImpl) Update(ctx context.Context, request web.MovieUpdateRequest) web.MovieResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movie, err := service.MovieRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	//set update from request
	movie.Title = request.Title
	movie.Description = request.Description
	movie.Rating = request.Rating
	movie.Image = request.Image

	movie = service.MovieRepository.Update(ctx, tx, movie)

	return helper.ToMovieResponse(movie)
}

func (service *MovieServiceImpl) Delete(ctx context.Context, movieId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movie, err := service.MovieRepository.FindById(ctx, tx, movieId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.MovieRepository.Delete(ctx, tx, movie)
}

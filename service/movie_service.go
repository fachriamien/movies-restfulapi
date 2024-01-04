package service

import (
	"context"
	"fchramn/movies-restfulapi/model/web"
)

type MovieService interface {
	FindAll(ctx context.Context) []web.MovieResponse
	FindById(ctx context.Context, movieId int) web.MovieResponse
	Create(ctx context.Context, request web.MovieCreateRequest) web.MovieResponse
	Update(ctx context.Context, request web.MovieUpdateRequest) web.MovieResponse
	Delete(ctx context.Context, movieId int)
}

package helper

import (
	"fchramn/movies-restfulapi/model/domain"
	"fchramn/movies-restfulapi/model/web"
)

func ToMovieResponse(movie domain.Movie) web.MovieResponse {
	return web.MovieResponse{
		Id:          movie.Id,
		Title:       movie.Title,
		Description: movie.Description,
		Rating:      movie.Rating,
		Image:       movie.Image,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
	}
}

func ToMovieResponses(movies []domain.Movie) []web.MovieResponse {
	var movieResponses []web.MovieResponse
	for _, movie := range movies {
		movieResponses = append(movieResponses, ToMovieResponse(movie))
	}
	return movieResponses
}

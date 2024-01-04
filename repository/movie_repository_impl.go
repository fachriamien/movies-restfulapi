package repository

import (
	"context"
	"database/sql"
	"errors"
	"fchramn/movies-restfulapi/helper"
	"fchramn/movies-restfulapi/model/domain"
)

type MovieRepositoryImpl struct {
}

func NewMovieRepository() MovieRepository {
	return &MovieRepositoryImpl{}
}

func (repository *MovieRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Movie {
	SQL := "SELECT * FROM movies"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var movies []domain.Movie
	for rows.Next() {
		movie := domain.Movie{}
		var updatedAt sql.NullString
		var createdAt sql.NullString

		err := rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.Rating, &movie.Image, &createdAt, &updatedAt)

		movie.UpdatedAt = helper.ParserNull(updatedAt)
		movie.CreatedAt = helper.ParserNull(createdAt)
		movie.UpdatedAt = helper.ParserTime(movie.UpdatedAt)
		movie.CreatedAt = helper.ParserTime(movie.CreatedAt)

		helper.PanicIfError(err)
		movies = append(movies, movie)
	}

	return movies
}

func (repository *MovieRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, movieId int) (domain.Movie, error) {
	SQL := "SELECT * FROM movies WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, movieId)
	helper.PanicIfError(err)
	defer rows.Close()

	movie := domain.Movie{}

	var updatedAt sql.NullString
	var createdAt sql.NullString

	if rows.Next() {
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.Rating, &movie.Image, &createdAt, &updatedAt)
		movie.UpdatedAt = helper.ParserNull(updatedAt)
		movie.CreatedAt = helper.ParserNull(createdAt)
		movie.UpdatedAt = helper.ParserTime(movie.UpdatedAt)
		movie.CreatedAt = helper.ParserTime(movie.CreatedAt)

		helper.PanicIfError(err)
		return movie, nil
	} else {
		return movie, errors.New("movie is not found")
	}
}

func (repository *MovieRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie {
	SQL := "INSERT INTO movies(title, description, rating, image, created_at) VALUES (?, ?, ?, ?, NOW())"
	result, err := tx.ExecContext(ctx, SQL, movie.Title, movie.Description, movie.Rating, movie.Image)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	movie.Id = int(id)
	return movie
}

func (repository *MovieRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie {
	SQL := "UPDATE movies SET title = ?, description = ?, rating = ?, image = ?, updated_at = NOW() WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, movie.Title, movie.Description, movie.Rating, movie.Image, movie.Id)
	helper.PanicIfError(err)

	return movie
}

func (repository *MovieRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, movie domain.Movie) {
	SQL := "DELETE FROM movies WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, movie.Id)
	helper.PanicIfError(err)
}

package repository

import (
	"context"
	"database/sql"
	"fchramn/movies-restfulapi/model/domain"
)

type MovieRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Movie
	FindById(ctx context.Context, tx *sql.Tx, movieId int) (domain.Movie, error)
	Save(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie
	Update(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie
	Delete(ctx context.Context, tx *sql.Tx, movie domain.Movie)
}

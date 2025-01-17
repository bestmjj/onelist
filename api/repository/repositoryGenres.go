package repository

import "github.com/bestmjj/onelist/onelist/api/models"

type GenreRepository interface {
	Save(models.Genre) (models.Genre, error)
	FindAll(page int, size int) ([]models.Genre, int, error)
	FindByID(string) (models.Genre, error)
	UpdateByID(string, models.Genre) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.Genre, int, error)
	FindByIdFilte(string, string, string, string, string, int, int) (models.Genre, int, error)
}

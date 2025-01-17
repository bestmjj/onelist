package repository

import "github.com/bestmjj/onelist/onelist/api/models"

type TheMovieRepository interface {
	Save(models.TheMovie) (models.TheMovie, error)
	FindAll(page int, size int) ([]models.TheMovie, int, error)
	FindByID(string) (models.TheMovie, error)
	UpdateByID(string, models.TheMovie) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.TheMovie, int, error)
	Sort(string, string, string, int, int) ([]models.TheMovie, int, error)
	FindByGalleryId(string, int, int) ([]models.TheMovie, int, error)
}

package repository

import "github.com/bestmjj/onelist/onelist/api/models"

type CastItemRepository interface {
	Save(models.CastItem) (models.CastItem, error)
	FindAll(page int, size int) ([]models.CastItem, int, error)
	FindByID(string) (models.CastItem, error)
	UpdateByID(string, models.CastItem) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.CastItem, int, error)
}

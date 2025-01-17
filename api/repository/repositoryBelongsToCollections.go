package repository

import "github.com/bestmjj/onelist/onelist/api/models"

type BelongsToCollectionRepository interface {
	Save(models.BelongsToCollection) (models.BelongsToCollection, error)
	FindAll(page int, size int) ([]models.BelongsToCollection, int, error)
	FindByID(string) (models.BelongsToCollection, error)
	UpdateByID(string, models.BelongsToCollection) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.BelongsToCollection, int, error)
}

package repository

import "github.com/bestmjj/onelist/onelist/api/models"

type CrewItemRepository interface {
	Save(models.CrewItem) (models.CrewItem, error)
	FindAll(page int, size int) ([]models.CrewItem, int, error)
	FindByID(string) (models.CrewItem, error)
	UpdateByID(string, models.CrewItem) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.CrewItem, int, error)
}

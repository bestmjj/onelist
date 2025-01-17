package repository

import "github.com/bestmjj/onelist/onelist/api/models"

type ProductionCountrieRepository interface {
	Save(models.ProductionCountrie) (models.ProductionCountrie, error)
	FindAll(page int, size int) ([]models.ProductionCountrie, int, error)
	FindByID(string) (models.ProductionCountrie, error)
	UpdateByID(string, models.ProductionCountrie) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.ProductionCountrie, int, error)
}

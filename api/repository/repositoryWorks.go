package repository

import "github.com/bestmjj/onelist/onelist/api/models"

type WorkRepository interface {
	Save(models.Work) (models.Work, error)
	FindAll(page int, size int) ([]models.Work, int, error)
	FindByID(string) (models.Work, error)
	UpdateByID(string, models.Work) (int64, error)
	DeleteByID(string) (int64, error)
	Search(string, int, int) ([]models.Work, int, error)
	GetWorkListByGalleryId(string, int, int) ([]models.Work, int, error)
}

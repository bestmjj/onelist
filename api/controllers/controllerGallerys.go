package controllers

import (
	"strconv"
	"strings"

	"github.com/bestmjj/onelist/onelist/api/database"
	"github.com/bestmjj/onelistelist/onelist/api/models"
	"github.com/bestmjj/onelistelist/onelist/api/repository"
	"github.com/bestmjj/onelistelist/onelist/api/repository/crud"
	"github.com/bestmjj/onelistelist/onelist/plugins/alist"

	"github.com/gin-gonic/gin"
)

func CreateGallery(c *gin.Context) {
	gallery := models.Gallery{}
	err := c.ShouldBind(&gallery)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": gallery})
		return
	}
	if !strings.Contains(gallery.AlistHost, "http") && gallery.IsAlist {
		c.JSON(200, gin.H{"code": 201, "msg": "域名应该含有'http'!", "data": gallery})
		return
	}
	db := database.NewDb()
	gallery.AlistHost = strings.TrimRight(gallery.AlistHost, "/")
	// https://github.com/bestmjj/onelistelist/onelist/issues/97
	if gallery.IsAlist {
		_, err := alist.AlistLogin(gallery)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "alist 验证失败!", "data": err})
			return
		}
	}
	repo := crud.NewRepositoryGallerysCRUD(db)
	func(galleryRepository repository.GalleryRepository) {
		gallery, err := galleryRepository.Save(gallery)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "创建失败!", "data": gallery})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "创建成功!", "data": gallery})
	}(repo)
}

func DeleteGalleryById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryGallerysCRUD(db)
	func(galleryRepository repository.GalleryRepository) {
		gallery, err := galleryRepository.DeleteByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": gallery})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "删除资源成功!", "data": gallery})
	}(repo)
}

func UpdateGalleryById(c *gin.Context) {
	id := c.Query("id")
	gallery := models.Gallery{}
	err := c.ShouldBind(&gallery)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": "创建失败,表单解析出错!", "data": gallery})
		return
	}
	// https://github.com/bestmjj/onelistelist/onelist/issues/97
	if gallery.IsAlist {
		_, err := alist.AlistLogin(gallery)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "alist 验证失败!", "data": err})
			return
		}
	}
	db := database.NewDb()
	repo := crud.NewRepositoryGallerysCRUD(db)
	func(galleryRepository repository.GalleryRepository) {
		gallery, err := galleryRepository.UpdateByID(id, gallery)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": gallery})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "更新资源成功!", "data": gallery})
	}(repo)
}

func GetGalleryById(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryGallerysCRUD(db)
	func(galleryRepository repository.GalleryRepository) {
		gallery, err := galleryRepository.FindByID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": gallery})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": gallery})
	}(repo)
}

func GetGalleryList(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryGallerysCRUD(db)
	func(galleryRepository repository.GalleryRepository) {
		gallerys, num, err := galleryRepository.FindAll(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": gallerys, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": gallerys, "num": num})
	}(repo)
}

func GetGalleryListAdmin(c *gin.Context) {
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryGallerysCRUD(db)
	func(galleryRepository repository.GalleryRepository) {
		gallerys, num, err := galleryRepository.FindAllByAdmin(page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": gallerys, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": gallerys, "num": num})
	}(repo)
}

func GetGalleryHostByUid(c *gin.Context) {
	id := c.Query("id")
	db := database.NewDb()
	repo := crud.NewRepositoryGallerysCRUD(db)
	func(galleryRepository repository.GalleryRepository) {
		gallery, err := galleryRepository.FindByUID(id)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": "", "is_ali_open": false})
			return
		}
		if gallery.IsAlist {
			c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": gallery.AlistHost, "is_ali_open": gallery.IsAliOpen})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": "", "is_ali_open": gallery.IsAliOpen})
	}(repo)
}

func SearchGallery(c *gin.Context) {
	q := c.Query("q")
	if len(q) == 0 {
		c.JSON(200, gin.H{"code": 201, "msg": "参数错误!", "data": ""})
		return
	}
	page, errPage := strconv.Atoi(c.Query("page"))
	size, errSize := strconv.Atoi(c.Query("size"))
	if errPage != nil {
		page = 1
	}
	if errSize != nil {
		size = 8
	}
	db := database.NewDb()
	repo := crud.NewRepositoryGallerysCRUD(db)
	func(galleryRepository repository.GalleryRepository) {
		gallerys, num, err := galleryRepository.Search(q, page, size)
		if err != nil {
			c.JSON(200, gin.H{"code": 201, "msg": "没有查询到资源!", "data": gallerys, "num": num})
			return
		}
		c.JSON(200, gin.H{"code": 200, "msg": "查询资源成功!", "data": gallerys, "num": num})
	}(repo)
}

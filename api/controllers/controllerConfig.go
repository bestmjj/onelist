package controllers

import (
	"github.com/bestmjj/onelist/onelist/api/models"
	"github.com/bestmjj/onelistelist/onelist/config"
	"github.com/gin-gonic/gin"
)

func GetWebConfig(c *gin.Context) {
	configData := config.GetConfig()
	configData.KeyDb = ""
	c.JSON(200, gin.H{"code": 200, "msg": "获取成功!", "data": configData})
}

func GetConfig(c *gin.Context) {
	configData := config.GetConfig()
	c.JSON(200, gin.H{"code": 200, "msg": "获取成功!", "data": configData})
}

func SaveConfig(c *gin.Context) {
	configData := models.Config{}
	err := c.ShouldBind(&configData)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": err.Error(), "data": ""})
		return
	}
	data, err := config.SaveConfig(configData)
	if err != nil {
		c.JSON(200, gin.H{"code": 201, "msg": err.Error(), "data": ""})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "保存成功!", "data": data})
}

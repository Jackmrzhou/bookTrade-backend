package controllers

import (
	error2 "bookTrade-backend/app/error"
	"bookTrade-backend/app/services"
	"bookTrade-backend/dao"
	"bookTrade-backend/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type uploadImageResp struct {
	Key string `json:"key"`
}

func UploadImage(c *gin.Context) {
	if file, err := c.FormFile("image"); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
	} else {
		if key, err := services.StoreImage(file); err != nil {
			c.Set(error2.CodeKey, error2.StoreImageFail)
		} else {
			c.JSON(http.StatusOK, uploadImageResp{
				Key: key,
			})
		}
	}
}

type downloadImagegReq struct {
	Key string `binding:"required" form:"key"`
}

func DownloadImage(c *gin.Context) {

	var query downloadImagegReq
	if err := c.ShouldBindQuery(&query); err != nil {
		log.WithError(err).Error("bad request in download image")
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	fileRef, err := dao.GetFileRefByFileKey(query.Key)
	if err != nil {
		log.WithError(err).Error("get file ref failed")
		c.Set(error2.CodeKey, error2.ImageNotFound)
		return
	}
	fileData, err := services.FetchImage(query.Key)
	if err != nil {
		log.WithError(err).Error("fetch image failed")
		c.Set(error2.CodeKey, error2.FetchImageFail)
		return
	}
	ext := utils.GetExt(fileRef.FileName)
	c.Header("Content-Type", utils.TranslateExt(ext))
	c.Header("Content-Length", strconv.Itoa(len(fileData)))
	if _, err := c.Writer.Write(fileData); err != nil {
		log.WithError(err).Error("Write response failed")
		c.Set(error2.CodeKey, error2.ImageResp)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
}

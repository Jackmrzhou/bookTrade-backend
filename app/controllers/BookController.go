package controllers

import (
	error2 "bookTrade-backend/app/error"
	"bookTrade-backend/dao"
	"bookTrade-backend/models"
	"bookTrade-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type newSellBookReq struct {
	BookName string `json:"book_name" binding:"required"`
	ISBN string `json:"ISBN" binding:"required"`
	Author string `json:"author" binding:"required"`
	Price float32 `json:"price" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	CoverKey string `json:"cover_key" binding:"required"`
	CatalogID int `json:"catalog_id" binding:"required"`
}

type newSellBookResp struct {
	BookID int `json:"book_id"`
}

func NewSellBook(c *gin.Context) {
	var query newSellBookReq
	if err := c.ShouldBindJSON(&query); err != nil {
		logrus.Info(err)
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	var principle utils.Principle
	var e error
	if principle, e = utils.GetPrinciple(c); e != nil {
		logrus.WithError(e).Error("get principle failed")
		c.Set(error2.CodeKey, error2.ServerError)
		return
	}

	query.ISBN = utils.FormatISBN(query.ISBN)

	book := models.Book{
		UserID: principle.UserId,
		Name:query.BookName,
		ISBN:query.ISBN,
		Author:query.Author,
		Price:query.Price,
		Introduction:query.Introduction,
		CoverKey:query.CoverKey,
		CatalogID:query.CatalogID,
		Type:models.SELL,
		OutLink: models.OUTLINK + query.ISBN,
	}

	if err := dao.CreateBook(&book); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
		return
	}else {
		c.JSON(http.StatusOK, newSellBookResp{
			BookID:book.ID,
		})
	}
}

type newRequestBookReq struct {
	BookName string `json:"book_name" binding:"required"`
	Author string `json:"author" binding:"required"`
	Price float32 `json:"price" binding:"required"`
	CoverKey string `json:"cover_key" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	ISBN string `json:"ISBN"`
	CatalogID int `json:"catalog_id" binding:"required"`
}

type newRequestBookResp struct {
	BookID int `json:"book_id"`
}

func NewRequestBook(c *gin.Context) {
	var query newRequestBookReq
	if err := c.ShouldBindJSON(&query); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	var principle utils.Principle
	var e error
	if principle, e = utils.GetPrinciple(c); e != nil {
		logrus.WithError(e).Error("get principle failed")
		c.Set(error2.CodeKey, error2.ServerError)
		return
	}

	query.ISBN = utils.FormatISBN(query.ISBN)
	// ISBN not required
	book := models.Book{
		UserID:principle.UserId,
		Name:query.BookName,
		Author:query.Author,
		Price:query.Price,
		CoverKey:query.CoverKey,
		Introduction: query.Introduction,
		ISBN:query.ISBN,
		CatalogID:query.CatalogID,
		Type:models.REQUEST,
	}
	if query.ISBN != "" {
		book.OutLink = models.OUTLINK + query.ISBN
	}

	if err := dao.CreateBook(&book); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
		return
	}else {
		c.JSON(http.StatusOK, newRequestBookResp{
			BookID:book.ID,
		})
	}
}

type typicalCatalog struct {
	CatalogID int `json:"catalog_id"`
	Name string `json:"name"`
}

type catalogGroup struct {
	RootCatalog typicalCatalog `json:"root_catalog"`
	SubCatalogs []typicalCatalog `json:"sub_catalogs"`
}

type getCatalogsResp struct {
 	Groups []catalogGroup `json:"groups"`
}

func GetCatalogs(c *gin.Context) {
	if catalogs, err := dao.GetAllCatalogs(); err != nil{
		c.Set(error2.CodeKey, error2.ServerError)
		return
	}else {
		var resp getCatalogsResp
		for _, catalog := range catalogs{
			if catalog.ParentID == -1 {
				// get group
				group := catalogGroup{
					RootCatalog:typicalCatalog{
						CatalogID:catalog.ID,
						Name:catalog.Name,
					},
				}
				for _, c := range catalogs {
					// get all sub-catalog
					if c.ParentID == catalog.ID{
						group.SubCatalogs = append(group.SubCatalogs, typicalCatalog{
							CatalogID:c.ID,
							Name:c.Name,
						})
					}
				}
				resp.Groups = append(resp.Groups, group)
			}
		}
		c.JSON(http.StatusOK, resp)
	}
}

type getBookDetailReq struct {
	BookID int `form:"book_id" binding:"required"`
}

type getBookDetailResp struct {
	BookID int `json:"book_id"`
	UserID int `json:"user_id"`
	BookName string `json:"book_name"`
	Author string `json:"author"`
	Price float32 `json:"price"`
	CatalogName string `json:"catalog_name"`
	CoverKey string `json:"cover_key"`
	Introduction string `json:"introduction"`
	Type int `json:"type"`
}

func GetBookDetail(c *gin.Context) {
	var query getBookDetailReq
	if err := c.ShouldBindQuery(&query); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	// if user is registered, record it
	if session, err := c.Cookie("sessionID"); err == nil {
		if userID, err := utils.GetUserIDBySessionID(session); err == nil {
			if err = dao.UpdateOrCreateUserRecord(&models.UserRecord{
				UserID:userID,
				BookID:query.BookID,
				ViewedTime:time.Now(),
			}); err != nil {
				logrus.WithError(err).Info("store user record failed")
			}
		}else {
			logrus.WithError(err).Info("get user id failed")
		}
	}

	if book, err := dao.GetBookByID(query.BookID); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
	} else if catalog, err := dao.GetCatalogByID(book.CatalogID); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
	} else {
		logrus.Info(book)
		c.JSON(http.StatusOK, getBookDetailResp{
			BookID: book.ID,
			UserID:book.UserID,
			BookName:book.Name,
			Author:book.Author,
			Price:book.Price,
			CatalogName:catalog.Name,
			CoverKey:book.CoverKey,
			Introduction:book.Introduction,
			Type: book.Type,
		})
	}

}

func GetBookDetailInOrder(c *gin.Context) {
	var query getBookDetailReq
	if err := c.ShouldBindQuery(&query); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	if book, err := dao.GetBookByIDDeletedOrNot(query.BookID); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
	} else if catalog, err := dao.GetCatalogByID(book.CatalogID); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
	} else {
		logrus.Info(book)
		c.JSON(http.StatusOK, getBookDetailResp{
			BookID: book.ID,
			UserID:book.UserID,
			BookName:book.Name,
			Author:book.Author,
			Price:book.Price,
			CatalogName:catalog.Name,
			CoverKey:book.CoverKey,
			Introduction:book.Introduction,
			Type: book.Type,
		})
	}
}

type listAllReq struct {
	Start int `form:"start"`
	Limit int `form:"limit" binding:"required"`
	CatalogID int `form:"catalog_id" binding:"required"`
	Type int `form:"type" binding:"required"`
}

type bookBrief struct {
	BookID int `json:"book_id"`
	BookName string `json:"book_name"`
	Author string `json:"author"`
	Price float32 `json:"price"`
	CoverKey string `json:"cover_key"`
}

type listAllResp struct {
	Books []bookBrief `json:"books"`
}

func ListAll(c *gin.Context) {
	var query listAllReq
	if err := c.ShouldBindQuery(&query); err != nil {
		logrus.WithError(err)
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	var books []models.Book
	var err error

	if query.CatalogID == -1 {
		if books, err = dao.GetBooks(query.Start, query.Limit, query.Type); err != nil {
			c.Set(error2.CodeKey, error2.ServerError)
			return
		}
	} else {
		if books, err = dao.GetBooksByCatalogID(query.Start, query.Limit, query.CatalogID, query.Type); err != nil {
			c.Set(error2.CodeKey, error2.ServerError)
			return
		}
	}

	var resp listAllResp
	for _, book := range books {
		resp.Books = append(resp.Books, bookBrief{
			BookID:book.ID,
			BookName:book.Name,
			Author:book.Author,
			Price:book.Price,
			CoverKey:book.CoverKey,
		})
	}
	c.JSON(http.StatusOK, resp)
}
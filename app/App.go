package app

import (
	"bookTrade-backend/common"
	"bookTrade-backend/conf"
	"bookTrade-backend/dao"
	"bookTrade-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Application struct {
	StorageManager common.StorageService
	MailManager    common.MailSender
	Router         *gin.Engine
}

func InitApp(config *conf.AppConfig) {
	App = &Application{
		StorageManager: common.NewDefaultStorageService(config),
		Router:         gin.Default(),
		MailManager:    common.NewDefaultMailSender(config),
	}
}

func (app *Application) Run() {
	for k, v := range builtInCatalog{
		c := models.Catalog{Name:k, ParentID:-1}
		if err := dao.CreatCatalogIfNotExists(&c); err != nil{
			logrus.WithError(err).Fatal("start app failed due to check catalog")
		}
		for _, item := range v{
			item.ParentID = c.ID
			if err := dao.CreatCatalogIfNotExists(&item); err != nil{
				logrus.WithError(err).Fatal("start app failed due to check catalog")
			}
		}
	}
	app.Router.Run()
}

var App *Application

var builtInCatalog = map[string][]models.Catalog{
	"小说": {
		models.Catalog{Name:"中国当代小说"},
		models.Catalog{Name:"外国小说"},
		models.Catalog{Name:"魔幻"},
		models.Catalog{Name:"科幻"},
		models.Catalog{Name:"武侠"},
		models.Catalog{Name:"历史"},
	},
	"文学": {
		models.Catalog{Name:"文集"},
		models.Catalog{Name:"中国古诗词"},
		models.Catalog{Name:"外国诗歌"},
		models.Catalog{Name:"戏剧"},
		models.Catalog{Name:"中国现当代随笔"},
	},
	"计算机": {
		models.Catalog{Name:"计算机理论"},
		models.Catalog{Name:"操作系统"},
		models.Catalog{Name:"数据库"},
		models.Catalog{Name:"程序设计"},
		models.Catalog{Name:"网络与数据通讯"},
		models.Catalog{Name:"软件工程/开发项目管理"},
	},
	"其他":{},
}
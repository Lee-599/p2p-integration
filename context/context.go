package context

import (
	_ "embed"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"link-server/app"
	"link-server/module/account"
	"link-server/module/p2p"
	"link-server/module/resource"
	"link-server/utils"
	"os"
	"path/filepath"
)

var (
	gormDB   *gorm.DB
	httpUtil *utils.HttpUtil

	AccountService  account.Service
	P2pService      p2p.Service
	ResourceService resource.Service

	AccountApp  app.Account
	LoginApp    app.Login
	P2pApp      app.P2p
	ResourceApp app.Resource
	SettingApp  app.Setting
)

func init() {
	//初始化数据库
	initDB()
	//初始化http工具累
	initHttp()
	//初始化service
	initService()
	//初始化app
	initApp()
}

func initDB() {
	configPath := initConfigPath()
	db, err := gorm.Open(sqlite.Open(filepath.Join(configPath, "link.db")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		&account.Account{},
		&p2p.P2pConfig{},
		&resource.Resource{},
	)
	if err != nil {
		panic("failed to AutoMigrate Account")
	}
	gormDB = db
}

func initHttp() {
	httpUtil = utils.NewHttp()
}

func initService() {
	accountServiceImpl := account.NewServiceImpl(gormDB, httpUtil)
	AccountService = &accountServiceImpl
	p2pServiceImpl := p2p.NewServiceImpl(gormDB)
	P2pService = &p2pServiceImpl
	resourceServiceImpl := resource.NewServiceImpl(gormDB, httpUtil)
	ResourceService = &resourceServiceImpl
}

func initApp() {
	AccountApp = app.NewAccountApp(AccountService)
	LoginApp = app.NewLoginApp(AccountService, P2pService)
	P2pApp = app.NewP2pApp(P2pService, ResourceService)
	ResourceApp = app.NewResourceApp(ResourceService, AccountService)
	SettingApp = app.NewSettingApp(P2pService, AccountService)
}

func initConfigPath() string {
	// 初始化配置文件
	dir := "~/.link/"
	linkConfig, err := homedir.Expand(dir)
	if err != nil {
		panic("failed to homedir Expand")
	}
	_, err = os.Stat(linkConfig)
	if err == nil {
		return linkConfig
	}
	err = os.MkdirAll(linkConfig, os.ModePerm)
	if err != nil {
		fmt.Printf("failed to config Mkdir err%s\n", err)
		panic("failed to config Mkdir err")
	}
	return linkConfig
}

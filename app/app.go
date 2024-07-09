package app

import (
	"github.com/outagelab/outagelab/config"
	"github.com/outagelab/outagelab/db"
	"github.com/outagelab/outagelab/db/migrator"
	"github.com/outagelab/outagelab/httpserver"
	"github.com/outagelab/outagelab/repositories/accountrepo"
	"github.com/outagelab/outagelab/repositories/apikeyrepo"
	"github.com/outagelab/outagelab/repositories/userrepo"
	"github.com/outagelab/outagelab/services/accountservice"
	"github.com/outagelab/outagelab/services/authservice"
	"github.com/outagelab/outagelab/services/datapageservice"
	"github.com/outagelab/outagelab/services/userservice"
)

type OutageLabApp struct {
	httpServer *httpserver.HttpServer
	migrator   *migrator.Migrator
}

func New() *OutageLabApp {
	config := config.New()

	db := db.New(config.Db)

	accountRepo := accountrepo.New(db)
	apiKeyRepo := apikeyrepo.New(db)
	userRepo := userrepo.New(db)

	accountService := accountservice.New(accountRepo, apiKeyRepo)
	datapageService := datapageservice.New(accountRepo, apiKeyRepo)
	userService := userservice.New(userRepo)
	authService := authservice.New(userService, accountService)

	httpServer := httpserver.New(accountService, datapageService, authService)
	migrator := migrator.New(db.ToSqlDB())

	return &OutageLabApp{
		httpServer: httpServer,
		migrator:   migrator,
	}
}

func (a *OutageLabApp) Start() {
	a.migrator.Migrate()
	a.httpServer.Start()
}

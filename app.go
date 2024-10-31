package outagelab

import (
	"github.com/outagelab/outagelab/internal/config"
	"github.com/outagelab/outagelab/internal/db"
	"github.com/outagelab/outagelab/internal/db/migrator"
	"github.com/outagelab/outagelab/internal/httpserver"
	"github.com/outagelab/outagelab/internal/repositories/accountrepo"
	"github.com/outagelab/outagelab/internal/repositories/apikeyrepo"
	"github.com/outagelab/outagelab/internal/repositories/userrepo"
	"github.com/outagelab/outagelab/internal/services/accountservice"
	"github.com/outagelab/outagelab/internal/services/authservice"
	"github.com/outagelab/outagelab/internal/services/datapageservice"
	"github.com/outagelab/outagelab/internal/services/userservice"
)

type OutageLabApp struct {
	httpServer *httpserver.HttpServer
	migrator   *migrator.Migrator
}

type AppOptions struct {
	CustomIndexHead []byte
}

func NewApp(options *AppOptions) *OutageLabApp {
	if options == nil {
		options = &AppOptions{}
	}

	config := config.New()

	db := db.New(config.Db)

	accountRepo := accountrepo.New(db)
	apiKeyRepo := apikeyrepo.New(db)
	userRepo := userrepo.New(db)

	accountService := accountservice.New(accountRepo, apiKeyRepo)
	datapageService := datapageservice.New(accountService, apiKeyRepo)
	userService := userservice.New(userRepo)
	authService := authservice.New(userService, accountService)

	httpServer := httpserver.New(accountService, datapageService, authService, options.CustomIndexHead)
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
